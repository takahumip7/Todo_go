package main

import (
	"TODO_GO/db"
	"TODO_GO/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// DB接続
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// ルーティング設定
	http.HandleFunc("/todos", handlers.TodosHandler)     // GET, POST
	http.HandleFunc("/todos/", handlers.TodoByIDHandler) // GET(id), PUT, PATCH, DELETE

	fmt.Println("🌐 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
