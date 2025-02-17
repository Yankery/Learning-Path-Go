package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"rerdinglist.io/interal/data"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		books := []data.Book{
			{
				ID:        1,
				CreatedAt: time.Now(),
				Title:     "Sample 1",
				Published: 2019,
				Pages:     300,
				Genre:     []string{"Fiction"},
				Rating:    4.5,
				Version:   1,
			}, {
				ID:        2,
				CreatedAt: time.Now(),
				Title:     "Sample 2",
				Published: 2019,
				Pages:     300,
				Genre:     []string{"Thriller"},
				Rating:    4.5,
				Version:   1,
			},
		}
		js, err := json.Marshal(books)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		js = append(js, '\n')
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Added new book to the reading list.")
	}
}

func (app *application) getUpdateDeleteBooksHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getBook(w, r)

	case http.MethodPut:
		app.updateBook(w, r)

	case http.MethodDelete:
		app.deleteBook(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	book := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "Sample",
		Published: 2019,
		Pages:     300,
		Genre:     []string{"Fiction", "Thriller"},
		Rating:    4.5,
		Version:   1,
	}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Update details of book with ID: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Delete details of book with ID: %d", idInt)
}
