package services

import (
	"errors"
	"time"

	"github.com/DevKayoS/go-library-mvc/internal/books/models"
)

type BookService struct {
	bookRepo models.BookRepository
}

func NewBookService(bookRepo models.BookRepository) models.BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (b *BookService) CreateBook(book *models.Book) error {
	if book.Title == "" {
		return errors.New("Title is required")
	}
	if book.Author == "" {
		return errors.New("Author is required")
	}
	if book.Quantity <= 0 {
		return errors.New("Quantity cannot  be negative")
	}

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	return b.bookRepo.CreateBook(book)
}

func (b *BookService) DeleteBook(id int64) error {
	return b.bookRepo.DeleteBook(id)
}

func (b *BookService) GetAllBook() ([]*models.Book, error) {
	return b.bookRepo.GetAllBook()
}

func (b *BookService) GetBook(id int64) (*models.Book, error) {
	return b.bookRepo.GetBook(id)
}

func (b *BookService) UpdateBook(id int64, book *models.Book) error {
	book.UpdatedAt = time.Now()
	return b.bookRepo.UpdateBook(id, book)
}
