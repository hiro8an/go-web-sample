package auth

import "net/http"

type contextKey string

const (
	// UserIDContextKey はコンテキストに保存されるユーザーIDのキーです
	UserIDContextKey = contextKey("userID")
	// UsernameContextKey はコンテキストに保存されるユーザー名のキーです
	UsernameContextKey = contextKey("username")
)

// Username コンテキストからユーザー名を取得する
func Username(r *http.Request) string {
	username := r.Context().Value(UsernameContextKey).(string)
	return username
}

// UserID コンテキストからユーザーIDを取得する
func UserID(r *http.Request) int {
	userID := r.Context().Value(UserIDContextKey).(int)
	return userID
}
