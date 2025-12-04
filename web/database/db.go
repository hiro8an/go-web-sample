package database

import (
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
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
	db, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// GORM の *sql.DB を取得
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// マイグレーションドライバーを作成
	driver, err := sqlite3.WithInstance(sqlDB, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("マイグレーションドライバーの作成に失敗しました: %w", err)
	}

	// マイグレーションインスタンスを作成
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations", // マイグレーションファイルのパス
		"sqlite3",            // データベース名
		driver,
	)
	if err != nil {
		return fmt.Errorf("マイグレーションインスタンスの作成に失敗しました: %w", err)
	}

	// マイグレーションを実行
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("マイグレーションの実行に失敗しました: %w", err)
	}

	return nil
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
