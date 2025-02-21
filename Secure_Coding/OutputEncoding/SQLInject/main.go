package main

import (
	"fmt"
	"log"
	"net/http"

	"sqlinject/patient"
)

func main() {
	http.HandleFunc("/search", patient.HandleSearch)
	http.HandleFunc("/searchsafe", patient.HandleSearchSafe)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
