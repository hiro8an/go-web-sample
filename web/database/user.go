package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

// TableName は User 構造体が対応するテーブル名を指定します
func (User) TableName() string {
	return "users"
}

// CreateUser 新しいユーザを作成する
func CreateUser(username, passwordHash string) (int, error) {
	user := User{
		Username: username,
		Password: passwordHash,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

// GetUserByID ユーザをIDで取得
func GetUserByID(id int) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByUsername ユーザ名で検索
func GetUserByUsername(username string) (*User, error) {
	var user User
	result := db.Where(&User{Username: username}).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser ユーザを更新する
func UpdateUser(user *User) error {
	return db.Save(user).Error
}

// DeleteUser ユーザを削除する
func DeleteUser(id int) error {
	return db.Delete(&User{}, id).Error
}
