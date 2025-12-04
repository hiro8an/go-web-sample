package web

import (
	"go-web-sample/web/handler"
	"go-web-sample/web/middleware"
	"net/http"
)

// GetRouteHandler はルーティングを登録したハンドラーを返します。
func GetRouteHandler() http.Handler {
	mux := http.NewServeMux()

	// 認証不要なルート
	mux.HandleFunc("GET /login", handler.ShowLogin)
	mux.HandleFunc("POST /login", handler.Login)

	// 認証が必要なルートをまとめて登録
	protectedMux := http.NewServeMux()
	{
		// ホーム
		protectedMux.HandleFunc("GET /", handler.ShowHome)
		// ログアウト
		protectedMux.HandleFunc("GET /logout", handler.Logout)

		// Tasks CRUD
		// Index - タスク一覧表示
		protectedMux.HandleFunc("GET /tasks", handler.IndexTasks)
		// New - 新規作成フォーム表示
		protectedMux.HandleFunc("GET /tasks/new", handler.NewTask)
		// Create - 新規作成処理
		protectedMux.HandleFunc("POST /tasks", handler.CreateTask)
		// Show - 詳細表示
		protectedMux.HandleFunc("GET /tasks/{id}", handler.ShowTask)
		// Edit - 編集フォーム表示
		protectedMux.HandleFunc("GET /tasks/{id}/edit", handler.EditTask)
		// Update - 更新処理
		protectedMux.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
		// Destroy - 削除処理
		protectedMux.HandleFunc("DELETE /tasks/{id}", handler.DestroyTask)

		// 認証が必要なルート全体にRequireLoginを適用
		mux.Handle("/", middleware.RequireLogin(protectedMux.ServeHTTP))
	}

	// 共通ミドルウェア適用
	handler := middleware.MethodOverride(mux)
	handler = middleware.CSRF(handler)
	return handler
}
