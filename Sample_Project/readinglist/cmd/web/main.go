package main

import (
	"flag"
	"log"
	"net/http"

	"rerdinglist.io/interal/models"
)

type application struct {
	readinglist *models.ReadinglistModel
}

func main() {
	endpoint := flag.String("endpoint", "http://localhost:4000/v1/books", "Endpoint for readinglist web service")
	app := &application{
		readinglist: &models.ReadinglistModel{Endpoint: *endpoint},
	}

	addr := flag.String("addr", ":8080", "HTTP network address")
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Println("Starting the server on ", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
