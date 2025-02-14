package user

import (
	"testing"
	"time"
)

func TestGetOne(t *testing.T) {
	//arrange
	expect := User{
		ID:       42,
		Username: "mrobot",
	}
	users = []User{expect}

	//act
	got, err := getOne(expect.ID)

	//assert
	if err != nil {
		t.Fatal(err)
	}
	if got != expect {
		t.Errorf("Did not get the expected user. Got %+v, expected %+v", got, expect)
	}
}

func TestSlowOne(t *testing.T) {
	t.Parallel()
	t.Skip("Skipped")
	time.Sleep(1 * time.Second)
}

func TestSlowTwo(t *testing.T) {
	t.Parallel()
	t.Skip("Skipped")
	time.Sleep(1 * time.Second)
}
