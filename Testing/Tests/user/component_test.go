package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type mockResponseWriter struct {
	bytes.Buffer
}

func (mRW mockResponseWriter) Header() http.Header {
	return http.Header{}
}
func (mRW mockResponseWriter) WriteHeader(status int) {}

func (mRW mockResponseWriter) getData() []byte {
	return mRW.Bytes()
}

func TestHandler(t *testing.T) {
	users = []User{
		User{
			ID:       1,
			Username: "adent",
		},
		User{
			ID:       2,
			Username: "tmacmillan",
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := mockResponseWriter{}
	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	Handler(&rw, req)

	got := rw.getData()

	if !bytes.Equal(expect, got) {
		t.Fail()
	}

}
