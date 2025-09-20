package models

import (
	"TODO_GO/db"
	"time"
)

// Todo構造体
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// GET /todos
func GetTodos() ([]Todo, error) {
	rows, err := db.DB.Query("SELECT id, title, completed, created_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

// /todos/{id}
func GetTodoByID(id int) (*Todo, error) {
	row := db.DB.QueryRow("SELECT id, title, completed, created_at FROM todos WHERE id = ?", id)
	var t Todo
	err := row.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// POST /todos
func CreateTodo(t *Todo) error {
	result, err := db.DB.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", t.Title, t.Completed)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	t.ID = int(id)
	return nil
}

// UPDATE /todos
func UpdateTodo(t *Todo) error {
	_, err := db.DB.Exec("UPDATE todos SET title = ?, completed = ? WHERE id =?", t.Title, t.Completed, t.ID)
	return err
}

// Patch /todos
func PatchTodo(id int, completed bool) error {
	_, err := db.DB.Exec("UPDATE todos SET completed = ? WHERE id =?", completed, id)
	return err
}

// DELETE /todos
func DeleteTodo(id int) error {
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
