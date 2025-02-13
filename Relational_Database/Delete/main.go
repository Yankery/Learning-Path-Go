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
	rowsDeleted, err := deleteActor(201)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total actor affected: ", rowsDeleted)
}

// update
func deleteActor(actorId int64) (int64, error) {
	result, err := db.Exec("DELETE FROM actor WHERE actor_id = ?", actorId)
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %s", err)
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %s", err)
	}
	return rowsDeleted, nil
}
