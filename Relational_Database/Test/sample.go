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

	if err := modifyActor(db, 201, "James"); err != nil {
		log.Fatal(err)
	}
}

// update
func modifyActor(db *sql.DB, actorId int64, firstname string) error {
	if _, err := db.Exec("UPDATE actor SET first_name = ? WHERE actor_id = ?", firstname, actorId); err != nil {
		return err
	}
	return nil
}
