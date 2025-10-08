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

	// CORS対応ラッパー
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ✅ 必須ヘッダーを毎回セット
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// ✅ Preflight対応
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// ルーティング
		http.DefaultServeMux.ServeHTTP(w, r)
	})

	// ハンドラ登録
	http.HandleFunc("/todos", handlers.TodosHandler)     // GET, POST
	http.HandleFunc("/todos/", handlers.TodoByIDHandler) // GET(id), PUT, PATCH, DELETE

	fmt.Println("🌐 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
