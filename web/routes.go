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

	// Tasks CRUD
	// Index - タスク一覧表示
	mux.HandleFunc("GET /tasks", middleware.RequireLogin(handler.IndexTasks))
	// New - 新規作成フォーム表示
	mux.HandleFunc("GET /tasks/new", middleware.RequireLogin(handler.NewTask))
	// Create - 新規作成処理
	mux.HandleFunc("POST /tasks", middleware.RequireLogin(handler.CreateTask))
	// Show - 詳細表示
	mux.HandleFunc("GET /tasks/{id}", middleware.RequireLogin(handler.ShowTask))
	// Edit - 編集フォーム表示
	mux.HandleFunc("GET /tasks/{id}/edit", middleware.RequireLogin(handler.EditTask))
	// Update - 更新処理
	mux.HandleFunc("PUT /tasks/{id}", middleware.RequireLogin(handler.UpdateTask))
	// Destroy - 削除処理
	mux.HandleFunc("DELETE /tasks/{id}", middleware.RequireLogin(handler.DestroyTask))

	handler := middleware.MethodOverride(mux)
	handler = middleware.CSRF(handler)
	return handler
}
