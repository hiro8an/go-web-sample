package web

import (
	"go-web-sample/web/handler"
	"go-web-sample/web/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", middleware.RequireLogin(handler.HomeHandler))
    mux.HandleFunc("/login", handler.LoginPageHandler)
    mux.HandleFunc("/api/login", handler.LoginHandler)
	mux.HandleFunc("/logout", middleware.RequireLogin(handler.LogoutHandler))
}
