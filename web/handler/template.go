package handler

import (
	"fmt"
	"html/template"
)

var (
	loginTmpl    *template.Template
	homeTmpl     *template.Template
	tasksTmpl    *template.Template
	taskFormTmpl *template.Template
	taskShowTmpl *template.Template
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

	tasksTmpl, err = template.ParseFiles("web/templates/tasks.html")
	if err != nil {
		fmt.Println("タスク一覧テンプレート読み込みエラー:", err)
	}

	taskFormTmpl, err = template.ParseFiles("web/templates/task_form.html")
	if err != nil {
		fmt.Println("タスクフォームテンプレート読み込みエラー:", err)
	}

	taskShowTmpl, err = template.ParseFiles("web/templates/task_show.html")
	if err != nil {
		fmt.Println("タスク詳細テンプレート読み込みエラー:", err)
	}
}
