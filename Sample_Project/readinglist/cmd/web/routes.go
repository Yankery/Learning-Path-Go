package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("D:/Coding/Learn_Go/Sample_Project/readinglist/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/book/view", app.bookView)
	mux.HandleFunc("/book/create", app.bookCreate)

	return mux
}
