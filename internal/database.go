package internal

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type CrudDB struct {
	sync.Mutex
	Store map[uuid.UUID]*Book
}

func (db *CrudDB) Get(ID uuid.UUID) (*Book, error) {
	db.Lock()
	defer db.Unlock()
	_, exists := db.Store[ID]
	if !exists {
		return nil, errors.New("ID not exists")
	}
	return db.Store[ID], nil

}

func (db *CrudDB) GetAll() ([]Book, error) {
	db.Lock()
	defer db.Unlock()

	if len(db.Store) == 0 {
		return nil, errors.New("empty database")
	}

	books := make([]Book, 0)
	for _, book := range db.Store {
		books = append(books, *book)
	}
	return books, nil
}

func (db *CrudDB) Upsert(b *Book) {
	db.Lock()
	defer db.Unlock()

	db.Store[b.ID] = b
}

func (db *CrudDB) Delete(ID uuid.UUID) error {
	db.Lock()
	defer db.Unlock()

	_, exists := db.Store[ID]
	if !exists {
		return errors.New("ID not exists")
	}

	delete(db.Store, ID)
	return nil
}
