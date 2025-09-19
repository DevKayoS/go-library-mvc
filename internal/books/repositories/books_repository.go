package repositories

import (
	"errors"
	"sync"

	"github.com/DevKayoS/go-library-mvc/internal/books/models"
)

type BookRepository struct {
	books  map[int64]*models.Book
	mu     sync.RWMutex
	nextID int64
}

func NewBookRepository() models.BookRepository {
	return &BookRepository{
		books:  make(map[int64]*models.Book),
		nextID: 1,
	}
}

func (b *BookRepository) CreateBook(book *models.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	book.Id = b.nextID
	b.nextID++

	b.books[book.Id] = book

	return nil
}

func (b *BookRepository) DeleteBook(id int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	book, exits := b.books[id]

	if !exits {
		return errors.New("book not found")
	}

	delete(b.books, book.Id)
	return nil
}

func (b *BookRepository) GetAllBook() ([]*models.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	books := make([]*models.Book, 0, len(b.books))
	for _, book := range b.books {
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepository) GetBook(id int64) (*models.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	book, exists := b.books[id]

	if !exists {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func (b *BookRepository) UpdateBook(id int64, book *models.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	book, exists := b.books[id]

	if !exists {
		return errors.New("book not found")
	}

	b.books[book.Id] = book
	return nil
}
