package controller

import (
	"database/sql"
	"encoding/json"
	"go-clean-architecture/data"
	"go-clean-architecture/domain"
	"go-clean-architecture/sourceimpl"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct{}

func (controller Controller) GetAllBooksFromDB(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book data.Book

		bookRepo := domain.BookRepo{}
		books, error := bookRepo.GetAllBooks(sourceimpl.BookSourceImpl{db, book})

		if error.Message != "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}

func (controller Controller) GetBookById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book data.Book
		params := mux.Vars(r)

		bookRepo := domain.BookRepo{}
		book, error := bookRepo.GetBookById(params["id"], sourceimpl.BookSourceImpl{db, book})

		if error.Message != "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}
}

func (controller Controller) AddNewBookToDB(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook data.Book

		json.NewDecoder(r.Body).Decode(&newBook)

		bookRepo := domain.BookRepo{}
		response := bookRepo.InsertBook(newBook, sourceimpl.BookSourceImpl{db, newBook})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (controller Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book data.Book

		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := domain.BookRepo{}
		rowsUpdated, error := bookRepo.UpdateBook(book, sourceimpl.BookSourceImpl{db, book})

		if error.Message != "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&rowsUpdated)
	}
}

func (controller Controller) DeleteBookFromDB(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book data.Book
		params := mux.Vars(r)

		bookRepo := domain.BookRepo{}
		rowsDeleted, error := bookRepo.DeleteBook(params["id"], sourceimpl.BookSourceImpl{db, book})

		if error.Message != "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&rowsDeleted)
	}
}
