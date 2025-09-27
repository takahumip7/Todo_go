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

	// CORS対応ラッパー
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		http.DefaultServeMux.ServeHTTP(w, r)
	})

	fmt.Println("🌐 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
