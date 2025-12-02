package database

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int `gorm:"primaryKey"`
	UserID      int `gorm:"index"`
	Title       string
	Description string
	Completed   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

// TableName は Task 構造体が対応するテーブル名を指定します
func (Task) TableName() string {
	return "tasks"
}

// CreatedAtJST 作成日時を日本時間でフォーマットして返す
func (t *Task) CreatedAtJST() string {
	return t.CreatedAt.In(jst).Format("2006-01-02 15:04:05")
}

// CreateTask 新しいタスクを作成する
func CreateTask(userID int, title, description string) (int, error) {
	task := Task{
		UserID:      userID,
		Title:       title,
		Description: description,
	}
	result := db.Create(&task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.ID, nil
}

// GetTasksByUser ユーザに紐づくタスク一覧を取得
func GetTasksByUser(userID int) ([]Task, error) {
	var tasks []Task
	result := db.Where(&Task{UserID: userID}).Order("created_at DESC").Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// GetTaskByID タスクをIDで取得
func GetTaskByID(id int) (*Task, error) {
	var t Task
	result := db.First(&t, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &t, nil
}

// UpdateTask タスクを更新する
func UpdateTask(t *Task) error {
	return db.Save(t).Error
}

// DeleteTask タスクを削除する
func DeleteTask(id int) error {
	return db.Delete(&Task{}, id).Error
}
