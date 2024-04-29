package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:example@tcp(172.21.0.2:3306)/dev?charset=utf8")
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected")

	return db
}
