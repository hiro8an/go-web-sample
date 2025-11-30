package database

import (
	"database/sql"
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var db *sql.DB

// データベース初期化
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		return err
	}

	// ユーザテーブル作成
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	return err
}

func Seed() error {
	// TODO: 初期データの挿入などをここに実装

	// demo user
	username := "demo"
	password := "Password123"
	user, err := FindUserByUsername(username)
	if err != nil {
		log.Fatal(err)
	}
	if user == nil {
		err = RegisterUser(username, password)
		if err != nil {
			return err
		}
	}

	return nil
}

func CloseDB() error {
	return db.Close()
}
