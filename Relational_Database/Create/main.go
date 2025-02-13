package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	//data source name properties
	dsn := mysql.Config{
		User:   "root",
		Passwd: "2101040188",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "sakila",
	}

	//get database handler
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	//body
	actorID, err := addActor("JOE", "BERRY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID of added actor: ", actorID)
}

// Create
func addActor(firstname string, lastname string) (int64, error) {
	result, err := db.Exec("INSERT INTO actor(first_name, last_name) VALUES (?, ?)", firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("addActor: %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor: %s", err)
	}
	return id, nil
}
