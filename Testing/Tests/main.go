	package main

import (
	"log"
	"net/http"
	"unites/user"
)

func main() {
	const address = ":3000"

	http.HandleFunc("/users", user.Handler)

	log.Fatal(http.ListenAndServe(address, nil))
}
