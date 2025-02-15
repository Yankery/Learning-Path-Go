package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

func TestHandler_httptest(t *testing.T) {
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

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	Handler(rec, req)
	res := rec.Result()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expect, got) {
		t.Fail()
	}
}

func TestHandler_server(t *testing.T) {
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

	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	s := httptest.NewServer(http.HandlerFunc(Handler))
	c := s.Client()

	res, err := c.Get(s.URL + "/users")

	if err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expect, got) {
		t.Fail()
	}
}

func BenchmarkHandler(b *testing.B) {

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
		b.Fatal(err)
	}

	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.Log(b.N)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rw := mockResponseWriter{}
			if err != nil {
				b.Fatal(err)
			}
			Handler(&rw, req)
		}
	})

}
