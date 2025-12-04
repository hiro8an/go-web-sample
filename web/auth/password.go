package auth

import (
	"go-web-sample/web/database"

	"golang.org/x/crypto/bcrypt"
)

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
