package handler

import (
	"go-web-sample/web/auth"
	"net/http"
)

// ホーム画面表示
func ShowHome(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthManager.GetUser(r)
	if err != nil {
		http.Error(w, "ユーザー取得エラー", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "ユーザーが見つかりません", http.StatusUnauthorized)
		return
	}
	homeTmpl.Execute(w, map[string]interface{}{
		"Username": user.Username,
	})
}
