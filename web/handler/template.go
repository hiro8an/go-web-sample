package handler

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed templates/*.html
var templateFS embed.FS

var (
	loginTmpl    *template.Template
	homeTmpl     *template.Template
	tasksTmpl    *template.Template
	taskFormTmpl *template.Template
	taskShowTmpl *template.Template
)

func init() {
	var err error
	loginTmpl, err = template.ParseFS(templateFS, "templates/login.html")
	if err != nil {
		fmt.Println("ログインテンプレート読み込みエラー:", err)
	}

	homeTmpl, err = template.ParseFS(templateFS, "templates/home.html")
	if err != nil {
		fmt.Println("ホームテンプレート読み込みエラー:", err)
	}

	tasksTmpl, err = template.ParseFS(templateFS, "templates/tasks.html")
	if err != nil {
		fmt.Println("タスク一覧テンプレート読み込みエラー:", err)
	}

	taskFormTmpl, err = template.ParseFS(templateFS, "templates/task_form.html")
	if err != nil {
		fmt.Println("タスクフォームテンプレート読み込みエラー:", err)
	}

	taskShowTmpl, err = template.ParseFS(templateFS, "templates/task_show.html")
	if err != nil {
		fmt.Println("タスク詳細テンプレート読み込みエラー:", err)
	}
}
