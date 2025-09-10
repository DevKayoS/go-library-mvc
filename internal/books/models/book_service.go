package models

type BookService interface {
	CreateBook(book *Book) error
	GetBook(id int64) (*Book, error)
	GetAllBook() ([]*Book, error)
	UpdateBook(id int64, Book *Book) error
	DeleteBook(id int64) error
}
