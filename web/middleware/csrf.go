package middleware

import (
	"net/http"
	"os"
)

// CSRF は CSRF 保護を適用するミドルウェアです
func CSRF(next http.Handler) http.Handler {
	// 同一オリジン以外のリクエストをブロックする Cross-Origin Protection
	cop := http.NewCrossOriginProtection()

	// 信頼オリジンの追加
	origin := os.Getenv("TRUSTED_ORIGIN")
	if origin != "" {
		cop.AddTrustedOrigin(origin)
	}

	return cop.Handler(next)
}
