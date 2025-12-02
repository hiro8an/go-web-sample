package handler

import (
	"net/http"

	"go-web-sample/web/auth"
)

// ログイン画面表示
func ShowLogin(w http.ResponseWriter, r *http.Request) {
	loginTmpl.Execute(w, nil)
}

// ログイン処理（セッション生成）
func Login(w http.ResponseWriter, r *http.Request) {
	// フォームパース
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "フォーム解析エラー", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// 簡易的なユーザー認証(実際にはデータベースで確認)
	if username == "" || password == "" {
		http.Error(w, "ユーザー名またはパスワードが空です", http.StatusBadRequest)
		return
	}

	ok, err := auth.AuthenticateUser(username, password)
	if err != nil {
		http.Error(w, "認証エラー", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "ユーザ名またはパスワードが不正です", http.StatusUnauthorized)
		return
	}

	// セッション作成 (cookie は内部で保存される)
	_, err = auth.AuthManager.CreateSession(w, r, username)
	if err != nil {
		http.Error(w, "セッション生成エラー", http.StatusInternalServerError)
		return
	}

	// ホームへリダイレクト
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ログアウト
func Logout(w http.ResponseWriter, r *http.Request) {
	_ = auth.AuthManager.DeleteSession(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
