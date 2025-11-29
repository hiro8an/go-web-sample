package middleware

import (
	"net/http"

	"go-web-sample/web/auth"
)

// RequireLogin はセッションを検証し、未ログインの場合はログインページにリダイレクトするミドルウェアです。
// 使用例: mux.HandleFunc("/profile", middleware.RequireLogin(ProfileHandler))
func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        _, ok := auth.AuthManager.GetUsername(r)
        if !ok {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next(w, r)
    }
}
