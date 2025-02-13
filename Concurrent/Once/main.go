package main

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Run()
	Run()
}

var db *sql.DB
var o sync.Once

func Run() {
	//ensure function only being executed once
	o.Do(func() {
		log.Println("Opening connection to the database")
		var err error
		db, err = sql.Open("sqlite3", "./mydb.db")

		if err != nil {
			log.Fatal(err)
		}
	})
}
