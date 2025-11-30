package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

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
	if err != nil {
		return err
	}

	// タスクテーブル作成（ユーザごとのタスク管理）
	createTasksSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		description TEXT,
		completed INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = db.Exec(createTasksSQL)
	return err
}

func Seed() error {
	// TODO: 初期データの挿入などをここに実装

	// demo user
	username := "demo"
	password := "Password123"
	user, err := GetUserByUsername(username)
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
