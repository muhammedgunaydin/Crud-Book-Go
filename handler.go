package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/muhammedgunaydin/book-crud/internal"
)

type Application struct {
	database *internal.CrudDB
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	var book internal.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	book.ID = uuid.New()
	app.database.Upsert(&book)
	w.WriteHeader(http.StatusCreated)
}

func (app *Application) ReadAll(w http.ResponseWriter, r *http.Request) {
	books, err := app.database.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(books)
	booksJSON, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(booksJSON)
	w.WriteHeader(http.StatusOK)
}

func (app *Application) Read(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	key, err := uuid.Parse(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	book, err := app.database.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	bookJSON, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(bookJSON)
	w.WriteHeader(http.StatusOK)
}

func (app *Application) Update(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	key, err := uuid.Parse(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	var book internal.Book

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	book.ID = key
	app.database.Upsert(&book)
	w.WriteHeader(http.StatusCreated)
}

func (app *Application) Delete(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	key, err := uuid.Parse(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = app.database.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
