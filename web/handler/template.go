package handler

import (
	"fmt"
	"html/template"
)

var (
	loginTmpl	*template.Template
	homeTmpl    *template.Template
)

func init() {
	var err error
	loginTmpl, err = template.ParseFiles("web/templates/login.html")
	if err != nil {
		fmt.Println("ログインテンプレート読み込みエラー:", err)
	}

	homeTmpl, err = template.ParseFiles("web/templates/home.html")
	if err != nil {
		fmt.Println("ホームテンプレート読み込みエラー:", err)
	}
}


