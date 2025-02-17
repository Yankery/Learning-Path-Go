package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	//create handle function to handle route
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
	mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHander)

	return mux
}
