package main

import (
	"fmt"
	"go-web-sample/web"
	"go-web-sample/web/database"
	"log"
	"net/http"
)

func main() {
	// データベース初期化
	if err := database.InitDB(); err != nil {
		log.Fatal("データベース初期化エラー:", err)
	}
	defer database.CloseDB()
	if err := database.Seed(); err != nil {
		log.Fatal("初期データ挿入エラー:", err)
	}

	fmt.Println("サーバーは http://localhost:8080 で起動しています")
	fmt.Println("テストアカウント: ユーザー名=demo, パスワード=Password123")
	http.ListenAndServe(":8080", web.GetRouteHandler())
}
