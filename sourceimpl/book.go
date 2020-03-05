package sourceimpl

import (
	"database/sql"
	"go-clean-architecture/data"
)

type BookSourceImpl struct {
	Db   *sql.DB
	Book data.Book
}

func (impl BookSourceImpl) GetBooks() ([]data.Book, data.Error) {
	var error data.Error
	var books []data.Book

	rows, err := impl.Db.Query("select * from novel")

	if err != nil {
		error.Message = "Server Error"
		return []data.Book{}, error
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&impl.Book.ID, &impl.Book.Title, &impl.Book.Author, &impl.Book.Year)
		if err != nil {
			error.Message = "Something went wrong"
			return []data.Book{}, error
		}
		books = append(books, impl.Book)
	}

	return books, error
}

func (impl BookSourceImpl) GetBookById(id string) (data.Book, data.Error) {
	var error data.Error

	row := impl.Db.QueryRow("select * from novel where id=$1", id)
	err := row.Scan(&impl.Book.ID, &impl.Book.Author, &impl.Book.Title, &impl.Book.Year)

	if err != nil {
		error.Message = "Server Error"
		return data.Book{}, error
	}

	return impl.Book, error
}

func (impl BookSourceImpl) AddBook(newBook data.Book) data.Error {
	var error data.Error

	row := impl.Db.QueryRow("insert into novel (title, author, year) values($1, $2, $3) RETURNING id",
		impl.Book.Title, impl.Book.Author, impl.Book.Year)

	err := row.Scan(&impl.Book.ID)

	if err != nil {
		error.Message = "Server Error"
	} else {
		error.Message = "Inserted Successfully"
	}

	return error
}

func (impl BookSourceImpl) UpdateBook(book data.Book) (int64, data.Error) {
	var error data.Error

	result, err := impl.Db.Exec("update novel set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		impl.Book.Title, impl.Book.Author, impl.Book.Year, impl.Book.ID)

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		error.Message = "Server Error"
		return 0, error
	}

	return rowsUpdated, error
}

func (impl BookSourceImpl) RemoveBook(id string) (int64, data.Error) {
	var error data.Error

	result, err := impl.Db.Exec("delete from novel where id=$1", id)
	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		error.Message = "Server Error"
		return 0, error
	}

	return rowsDeleted, error
}
