package handler

import (
	"go-web-sample/web/auth"
	"net/http"
)

// ホーム画面表示
func ShowHome(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	homeTmpl.Execute(w, map[string]interface{}{
		"Username": username,
	})
}
