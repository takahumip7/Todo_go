package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	// DSN: ユーザー名:パスワード@tcp(ホスト:ポート)/DB名
	dsn := "root:Takahumi7@tcp(127.0.0.1:3306)/todo_go?parseTime=true&loc=Local"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}

	log.Println("MySQLに接続成功！")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
