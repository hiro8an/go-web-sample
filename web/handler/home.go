package handler

import (
	"go-web-sample/web/auth"
	"net/http"
)

// ホーム画面表示
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTmpl.Execute(w, map[string]interface{}{
		"Username": username,
	})
}
