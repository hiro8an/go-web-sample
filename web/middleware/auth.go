package middleware

import (
	"context"
	"net/http"

	"go-web-sample/web/auth"
)

// RequireLogin はセッションを検証し、未ログインの場合はログインページにリダイレクトするミドルウェアです。
// ログイン済みの場合、ユーザーIDとユーザー名をコンテキストに含めて後続のハンドラに渡します。
// 使用例: mux.HandleFunc("/profile", middleware.RequireLogin(ProfileHandler))
func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.GetUserFromSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// ユーザーIDとユーザー名をコンテキストに設定
		ctx := context.WithValue(r.Context(), auth.UserIDContextKey, user.ID)
		ctx = context.WithValue(ctx, auth.UsernameContextKey, user.Username)
		next(w, r.WithContext(ctx))
	}
}
