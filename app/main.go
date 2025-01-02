package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootpass@tcp(6:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success")
}
