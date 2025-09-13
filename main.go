package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Todo構造体
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// DSN: ユーザー名:パスワード@tcp(ホスト:ポート)/DB名
	dsn := "root:Takahumi7@tcp(127.0.0.1:3306)/todo_go?parseTime=true&loc=Local"

	// DB接続
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("MySQLに接続成功！")

	// ルーティング設定
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			createTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🌐 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// GET /todos
func getTodos(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT id, title, completed, created_at FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// POST /todos
func createTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /todos が呼ばれました")
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("受け取ったデータ: %+v\n", t)

	result, err := db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", t.Title, t.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	t.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}
