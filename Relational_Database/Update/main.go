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
	rowsUpdated, err := updateActor("James", 200)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total actor affected: ", rowsUpdated)
}

// update
func updateActor(firstname string, actorId int64) (int64, error) {
	result, err := db.Exec("UPDATE actor SET first_name = ? WHERE actor_id = ?", firstname, actorId)
	if err != nil {
		return 0, fmt.Errorf("updateActor: %s", err)
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateActor: %s", err)
	}
	return rowsUpdated, nil
}
