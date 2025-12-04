package handler

import (
	"go-web-sample/web/auth"
	"net/http"
)

// ホーム画面表示
func ShowHome(w http.ResponseWriter, r *http.Request) {
	homeTmpl.Execute(w, map[string]interface{}{
		"Username": auth.Username(r),
	})
}
