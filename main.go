package main

import (
	"fmt"
	"go-web-sample/web"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	web.RegisterRoutes(mux)

	fmt.Println("サーバーは http://localhost:8080 で起動しています")
	fmt.Println("テストアカウント: ユーザー名=demo, パスワード=demo")
	http.ListenAndServe(":8080", mux)
}
