package database

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

// パスワードハッシュ化
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// パスワード検証
func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ユーザ登録
func RegisterUser(username, password string) error {
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hash)
	return err
}

// ユーザ名で検索
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil // ユーザが見つからない場合は nil を返す
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ユーザ認証
func AuthenticateUser(username, password string) (bool, error) {
	user, err := GetUserByUsername(username)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return checkPassword(password, user.Password), nil
}
