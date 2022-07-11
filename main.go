package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/muhammedgunaydin/book-crud/internal"
)

func main() {
	db := &internal.CrudDB{
		Store: make(map[uuid.UUID]*internal.Book),
	}
	app := &Application{database: db}

	mux := mux.NewRouter()

	mux.HandleFunc("/v1/books", app.Create).Methods(http.MethodPost)
	mux.HandleFunc("/v1/books", app.ReadAll).Methods(http.MethodGet)
	mux.HandleFunc("/v1/books/{id}", app.Read).Methods(http.MethodGet)
	mux.HandleFunc("/v1/books/{id}", app.Update).Methods(http.MethodPut)
	mux.HandleFunc("/v1/books/{id}", app.Delete).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", mux)
}
