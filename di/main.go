package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	usecase := NewUseCase(db)

	product, err := usecase.GetProduct("5")
	if err != nil {
		panic(err)
	}

	fmt.Println(product)

}
