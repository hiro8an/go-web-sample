package handler

import (
	"net/http"

	"go-web-sample/web/auth"
)

// ログイン画面表示
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        loginTmpl.Execute(w, nil)
        return
    }
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// ログイン処理（セッション生成）
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

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

    // 簡易的な認証(実装例: demo/demo)
    if username != "demo" || password != "demo" {
        http.Error(w, "ユーザー名またはパスワードが不正です", http.StatusUnauthorized)
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
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    _ = auth.AuthManager.DeleteSession(w, r)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}
