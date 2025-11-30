package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go-web-sample/web/auth"
	"go-web-sample/web/database"
)

// Index - タスク一覧表示
func IndexTasks(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	user, _ := database.GetUserByUsername(username)
	if user == nil {
		http.Error(w, "ユーザが見つかりません", http.StatusUnauthorized)
		return
	}

	tasks, err := database.GetTasksByUser(user.ID)
	if err != nil {
		http.Error(w, "タスク取得エラー", http.StatusInternalServerError)
		return
	}

	tasksTmpl.Execute(w, map[string]interface{}{
		"Username": username,
		"Tasks":    tasks,
	})
}

// New - 新規作成フォーム表示
func NewTask(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	user, _ := database.GetUserByUsername(username)
	if user == nil {
		http.Error(w, "ユーザが見つかりません", http.StatusUnauthorized)
		return
	}
	taskFormTmpl.Execute(w, map[string]interface{}{
		"Username": username,
		"Mode":     "create",
	})
}

// Create - 新規作成処理
func CreateTask(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	user, _ := database.GetUserByUsername(username)
	if user == nil {
		http.Error(w, "ユーザが見つかりません", http.StatusUnauthorized)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "フォーム解析エラー", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" {
		http.Error(w, "タイトルは必須です", http.StatusBadRequest)
		return
	}

	_, err := database.CreateTask(user.ID, title, description)
	if err != nil {
		http.Error(w, "タスク作成エラー", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// Show - 詳細表示
func ShowTask(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	user, _ := database.GetUserByUsername(username)
	if user == nil {
		http.Error(w, "ユーザが見つかりません", http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		http.Error(w, "無効なID", http.StatusBadRequest)
		return
	}

	task, err := database.GetTaskByID(id)
	if err != nil || task == nil || task.UserID != user.ID {
		http.Error(w, "タスクが見つかりません", http.StatusNotFound)
		return
	}

	taskShowTmpl.Execute(w, map[string]interface{}{
		"Username": username,
		"Task":     task,
	})
}

// Edit - 編集フォーム表示
func EditTask(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	user, _ := database.GetUserByUsername(username)
	if user == nil {
		http.Error(w, "ユーザが見つかりません", http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		http.Error(w, "無効なID", http.StatusBadRequest)
		return
	}

	task, err := database.GetTaskByID(id)
	if err != nil || task == nil || task.UserID != user.ID {
		http.Error(w, "タスクが見つかりません", http.StatusNotFound)
		return
	}

	taskFormTmpl.Execute(w, map[string]interface{}{
		"Username": username,
		"Mode":     "edit",
		"Task":     task,
	})
}

// Update - 更新処理
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	user, _ := database.GetUserByUsername(username)
	if user == nil {
		http.Error(w, "ユーザが見つかりません", http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		http.Error(w, "無効なID", http.StatusBadRequest)
		return
	}

	task, err := database.GetTaskByID(id)
	if err != nil || task == nil || task.UserID != user.ID {
		http.Error(w, "タスクが見つかりません", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "フォーム解析エラー", http.StatusBadRequest)
		return
	}
	task.Title = r.FormValue("title")
	task.Description = r.FormValue("description")
	task.Completed = r.FormValue("completed") == "on"

	if err := database.UpdateTask(task); err != nil {
		http.Error(w, "タスク更新エラー", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/tasks/%d", task.ID), http.StatusSeeOther)
}

// Destroy - 削除処理
func DestroyTask(w http.ResponseWriter, r *http.Request) {
	username, _ := auth.AuthManager.GetUsername(r)
	user, _ := database.GetUserByUsername(username)
	if user == nil {
		http.Error(w, "ユーザが見つかりません", http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		http.Error(w, "無効なID", http.StatusBadRequest)
		return
	}

	// 所有者チェック
	task, err := database.GetTaskByID(id)
	if err != nil || task == nil || task.UserID != user.ID {
		http.Error(w, "タスクが見つかりません", http.StatusNotFound)
		return
	}

	if err := database.DeleteTask(id); err != nil {
		http.Error(w, "タスク削除エラー", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
