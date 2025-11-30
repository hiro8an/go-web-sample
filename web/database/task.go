package database

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// CreatedAtJST 作成日時を日本時間でフォーマットして返す
func (t *Task) CreatedAtJST() string {
	return t.CreatedAt.In(jst).Format("2006-01-02 15:04:05")
}

// CreateTask 新しいタスクを作成する
func CreateTask(userID int, title, description string) (int64, error) {
	res, err := db.Exec("INSERT INTO tasks (user_id, title, description) VALUES (?, ?, ?)", userID, title, description)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetTasksByUser ユーザに紐づくタスク一覧を取得
func GetTasksByUser(userID int) ([]Task, error) {
	rows, err := db.Query("SELECT id, user_id, title, description, completed, created_at FROM tasks WHERE user_id = ? ORDER BY created_at DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		var completed int
		var created string
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &completed, &created); err != nil {
			return nil, err
		}
		t.Completed = completed != 0
		t.CreatedAt, _ = time.Parse(time.RFC3339, created)
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID タスクをIDで取得
func GetTaskByID(id int) (*Task, error) {
	var t Task
	var completed int
	var created string
	err := db.QueryRow("SELECT id, user_id, title, description, completed, created_at FROM tasks WHERE id = ?", id).
		Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &completed, &created)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t.Completed = completed != 0
	t.CreatedAt, _ = time.Parse(time.RFC3339, created)
	return &t, nil
}

// UpdateTask タスクを更新する
func UpdateTask(t *Task) error {
	completed := 0
	if t.Completed {
		completed = 1
	}
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?", t.Title, t.Description, completed, t.ID)
	return err
}

// DeleteTask タスクを削除する
func DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
