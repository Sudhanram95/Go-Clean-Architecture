package domain

import (
	"go-clean-architecture/data"
)

type BookSource interface {
	GetBooks() ([]data.Book, data.Error)
	GetBookById(id string) (data.Book, data.Error)
	AddBook(newBook data.Book) data.Error
	UpdateBook(book data.Book) (int64, data.Error)
	RemoveBook(id string) (int64, data.Error)
}
