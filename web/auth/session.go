package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"

	"go-web-sample/web/database"
)

const (
	SessionCookieName = "session_id"
	SessionTTL        = 24 * time.Hour
)

var store *sessions.CookieStore

func init() {
	envKey := os.Getenv("SESSION_KEY")
	var key []byte
	if envKey != "" {
		key = []byte(envKey)
	} else {
		fmt.Println("警告: SESSION_KEY 環境変数が設定されていません。デフォルトの固定キーを使用します。")
		key = []byte("a-very-secret-key-that-is-32-bytes!")
	}

	store = sessions.NewCookieStore(key)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(SessionTTL.Seconds()),
		HttpOnly: true,
	}
}

func randID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// CreateSession セッションを作成する
func CreateSession(w http.ResponseWriter, r *http.Request, username string) (string, error) {
	sess, err := store.Get(r, SessionCookieName)
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
		MaxAge:   int(SessionTTL.Seconds()),
		HttpOnly: true,
	}
	if err := sess.Save(r, w); err != nil {
		return "", err
	}
	return id, nil
}

// GetUsernameFromSession セッションからユーザー名を取得する
func GetUsernameFromSession(r *http.Request) (string, error) {
	sess, err := store.Get(r, SessionCookieName)
	if err != nil {
		return "", err
	}
	username, ok := sess.Values["username"].(string)
	if !ok || username == "" {
		return "", fmt.Errorf("username not found in session")
	}
	return username, nil
}

// GetUserFromSession セッションからユーザー情報を取得する
func GetUserFromSession(r *http.Request) (*database.User, error) {
	username, err := GetUsernameFromSession(r)
	if err != nil {
		return nil, err
	}
	return database.GetUserByUsername(username)
}

// DeleteSession セッションを削除する
func DeleteSession(w http.ResponseWriter, r *http.Request) error {
	sess, err := store.Get(r, SessionCookieName)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	return sess.Save(r, w)
}
