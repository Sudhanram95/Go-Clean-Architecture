package domain

import (
	"go-clean-architecture/data"
)

type BookRepo struct{}

func (repo BookRepo) GetAllBooks(bookSource interface{ BookSource }) ([]data.Book, data.Error) {
	return bookSource.GetBooks()
}

func (repo BookRepo) GetBookById(id string, bookSource BookSource) (data.Book, data.Error) {
	return bookSource.GetBookById(id)
}

func (repo BookRepo) InsertBook(newBook data.Book, bookSource BookSource) data.Error {
	return bookSource.AddBook(newBook)
}

func (repo BookRepo) UpdateBook(book data.Book, bookSource BookSource) (int64, data.Error) {
	return bookSource.UpdateBook(book)
}

func (repo BookRepo) DeleteBook(id string, bookSource BookSource) (int64, data.Error) {
	return bookSource.RemoveBook(id)
}
