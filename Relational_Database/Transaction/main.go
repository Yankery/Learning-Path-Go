package main

import (
	"context"
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
	ctx := context.Background()
	actorId, err := txActor(ctx, "JOEN", "BERRY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Actor ID: ", actorId)
}

// transaction
func txActor(ctx context.Context, firstname string, lastname string) (int64, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("Adding Actor Failed: %s", err)
	}
	//defer rollback in case of failure
	defer tx.Rollback()
	var actId int64
	if err = tx.QueryRowContext(ctx, "SELECT actor_id from actor where first_name = ? and last_name = ?", firstname, lastname).Scan(&actId); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Actor doesn't exist")
		} else {
			return 0, fmt.Errorf("txActor: %s", err)
		}
	}
	//rollback if actor exist
	if actId > 0 {
		if err = tx.Rollback(); err != nil {
			return 0, fmt.Errorf("txActor: %s", err)
		}
		fmt.Println("Actor already exist: ", actId)
		fmt.Println("---Transaction Rolling Back---")
		return actId, nil
	}

	//create new row
	result, err := tx.ExecContext(ctx, "INSERT INTO actor (first_name, last_name) VALUES (?,?)", firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("txActor: %s", err)
	}

	//get id of inserted actor
	newActorId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("txActor: %s", err)
	}

	//commit transaction
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("txActor: %s", err)
	} else {
		fmt.Println("New actor created: ", newActorId)
		fmt.Println("---Transaction Commited---")
	}
	return newActorId, nil
}
