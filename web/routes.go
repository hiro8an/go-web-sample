package web

import (
	"go-web-sample/web/handler"
	"go-web-sample/web/middleware"
	"net/http"
)

// GetRouteHandler はルーティングを登録したハンドラーを返します。
func GetRouteHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", middleware.RequireLogin(handler.ShowHome))
	mux.HandleFunc("GET /login", handler.ShowLogin)
	mux.HandleFunc("POST /login", handler.Login)
	mux.HandleFunc("GET /logout", middleware.RequireLogin(handler.Logout))
	handler := middleware.CSRF(mux)
	return handler
}
