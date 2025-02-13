package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_modifyActor_PASS(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE actor").WillReturnResult(sqlmock.NewResult(1, 1))

	if err = modifyActor(db, 50, "JOE"); err != nil {
		t.Errorf("Error was not expected while modifying actors: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectation: %s", err)
	}
}

func Test_modifyActor_FAIL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE actor").WillReturnError(fmt.Errorf("Some error"))

	if err = modifyActor(db, 50, "JOE"); err == nil {
		t.Errorf("Error was expected while modifying actors but no error occured")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectation: %s", err)
	}
}
