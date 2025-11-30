package middleware

import (
	"net/http"
	"strings"
)

// MethodOverride は _method フォームフィールドを使用してHTTPメソッドをオーバーライドするミドルウェアです。
// Rails 風で、フォームから PATCH や DELETE を送信する場合に使用します。
func MethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			method := strings.ToUpper(r.PostFormValue("_method"))
			if method == http.MethodPatch || method == http.MethodDelete || method == http.MethodPut {
				r.Method = method
			}
		}
		next.ServeHTTP(w, r)
	})
}
