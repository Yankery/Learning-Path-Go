package user_test

import (
	"testing"
	"unites/user"
)

func TestGetOne(t *testing.T) {
	u, err := user.GetOne(999)
	if u != (user.User{}) {
		t.Error()
	}
	if err == nil {
		t.Error()
	}
}
