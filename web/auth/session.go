package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"go-web-sample/web/database"
)

var (
	sessionCookieName = "session_id"
	sessionTTL        = 24 * time.Hour
)

// AuthManager は handler や middleware からアクセス可能なグローバルな認証用マネージャです。
var AuthManager *Manager

func init() {
	AuthManager = NewManager()
}

type Manager struct {
	store             *sessions.CookieStore
	sessionCookieName string
	sessionTTL        time.Duration
}

func NewManager() *Manager {
	envKey := os.Getenv("SESSION_KEY")
	var key []byte
	if envKey != "" {
		key = []byte(envKey)
	} else {
		fmt.Println("警告: SESSION_KEY 環境変数が設定されていません。デフォルトの固定キーを使用します。")
		key = []byte("a-very-secret-key-that-is-32-bytes!")
	}

	store := sessions.NewCookieStore(key)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(sessionTTL.Seconds()),
		HttpOnly: true,
	}
	return &Manager{store: store, sessionCookieName: sessionCookieName, sessionTTL: sessionTTL}
}

func randID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (m *Manager) CreateSession(w http.ResponseWriter, r *http.Request, username string) (string, error) {
	sess, err := m.store.Get(r, m.sessionCookieName)
	if err != nil {
		return "", err
	}
	id, err := randID()
	if err != nil {
		return "", err
	}
	sess.Values["id"] = id
	sess.Values["username"] = username
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(m.sessionTTL.Seconds()),
		HttpOnly: true,
	}
	if err := sess.Save(r, w); err != nil {
		return "", err
	}
	return id, nil
}

func (m *Manager) GetUsername(r *http.Request) (string, bool) {
	sess, err := m.store.Get(r, m.sessionCookieName)
	if err != nil {
		return "", false
	}
	username, ok := sess.Values["username"].(string)
	if !ok || username == "" {
		return "", false
	}
	return username, true
}

// GetUser セッションからユーザー情報を取得する
func (m *Manager) GetUser(r *http.Request) (*database.User, error) {
	username, ok := m.GetUsername(r)
	if !ok {
		return nil, nil
	}
	return database.GetUserByUsername(username)
}

func (m *Manager) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	sess, err := m.store.Get(r, m.sessionCookieName)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	return sess.Save(r, w)
}

// hashPassword パスワードをハッシュ化する
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPassword パスワードを検証する
func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RegisterUser ユーザを登録する
func RegisterUser(username, password string) error {
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = database.CreateUser(username, hash)
	return err
}

// AuthenticateUser ユーザ認証を行う
func AuthenticateUser(username, password string) (bool, error) {
	user, err := database.GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return checkPassword(password, user.Password), nil
}
