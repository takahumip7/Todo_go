package handlers

import (
	"TODO_GO/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// /todos 用（一覧と作成）
func TodosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// /todos/{id} 用（取得・更新・削除）
func TodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTodoByID(w, id)
	case http.MethodPut:
		updateTodo(w, r, id)
	case http.MethodPatch:
		patchTodo(w, r, id)
	case http.MethodDelete:
		deleteTodo(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /todos
func getTodos(w http.ResponseWriter) {
	todos, err := models.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, todos)
}

// GET /todos/{id}
func getTodoByID(w http.ResponseWriter, id int) {
	todo, err := models.GetTodoByID(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	writeJSON(w, todo)
}

// POST /todos
func createTodo(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := models.CreateTodo(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, t)
}

// PUT /todos/{id}（全体更新）
func updateTodo(w http.ResponseWriter, r *http.Request, id int) {
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.ID = id
	if err := models.UpdateTodo(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, t)
}

// PATCH /todos/{id}（部分更新）
func patchTodo(w http.ResponseWriter, r *http.Request, id int) {
	var data struct {
		Completed bool `json:"completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := models.PatchTodo(id, data.Completed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]any{"id": id, "completed": data.Completed})
}

// DELETE /todos/{id}
func deleteTodo(w http.ResponseWriter, id int) {
	if err := models.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"message": "Todo deleted"})
}

// 共通：JSONレスポンス
func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
