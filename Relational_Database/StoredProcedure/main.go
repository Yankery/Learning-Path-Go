package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
}

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

	actors, err := getActorSP("Joe", "Simple")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Actor found: ", actors)
}

// retrieve
func getActorSP(firstname string, lastname string) ([]Actor, error) {
	var actors []Actor

	result, err := db.Query("CALL addActor(?,?)", firstname, lastname)
	if err != nil {
		return nil, fmt.Errorf("getActorSP: %s", err)
	}
	defer result.Close()

	//loop through rows
	for result.Next() {
		var act Actor
		//scan copy the value in current row into value pointed by destination
		if err := result.Scan(&act.actor_id, &act.first_name, &act.last_name); err != nil {
			return nil, fmt.Errorf("getActorSP : %s", err)
		}
		actors = append(actors, act)

		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("getActor: %s", err)
		}
	}
	return actors, nil
}
