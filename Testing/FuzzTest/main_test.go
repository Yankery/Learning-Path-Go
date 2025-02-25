package main

import (
	"strconv"
	"strings"
	"testing"
)

func toUpper(s string) string {
	var b strings.Builder
	for i := range s {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			b.WriteByte(byte(s[i]))
			continue
		}
		b.WriteByte(byte(s[i]) - 32)
	}
	return b.String()
}

func TestUpper(t *testing.T) {
	have := "hello"
	expect := "HELLO"
	got := toUpper(have)
	t.Log(got, expect)
	if expect != got {
		t.Fail()
	}
}

func FuzzToUpper(f *testing.F) {
	f.Add("hello")
	f.Fuzz(func(t *testing.T, s string) {
		out := toUpper(s)
		if out != strings.ToUpper(s) {
			t.Fail()
		}
	})
}
