package main

import (
	"fmt"
	"log"
	"net/http"

	"go-web-sample/web"
	"go-web-sample/web/auth"
	"go-web-sample/web/database"
)

func main() {
	// データベース初期化
	if err := database.InitDB(); err != nil {
		log.Fatal("データベース初期化エラー:", err)
	}
	defer database.CloseDB()

	// 初期データ挿入
	if err := seed(); err != nil {
		log.Fatal("初期データ挿入エラー:", err)
	}

	fmt.Println("サーバーは http://localhost:8080 で起動しています")
	fmt.Println("テストアカウント: ユーザー名=demo, パスワード=Password123")
	http.ListenAndServe(":8080", web.GetRouteHandler())
}

func seed() error {
	// TODO: 初期データの挿入などをここに実装

	// デモユーザー登録
	username := "demo"
	password := "Password123"
	user, err := database.GetUserByUsername(username)
	if err != nil {
		log.Fatal("ユーザー検索エラー:", err)
	}
	if user == nil {
		err = auth.RegisterUser(username, password)
		if err != nil {
			log.Fatal("デモユーザー登録エラー:", err)
		}
	}
	return nil
}
