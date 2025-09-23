package services

import (
	"errors"
	"time"

	bookModels "github.com/DevKayoS/go-library-mvc/internal/books/models"
	"github.com/DevKayoS/go-library-mvc/internal/loans/models"
	loanModels "github.com/DevKayoS/go-library-mvc/internal/loans/models"
	userModels "github.com/DevKayoS/go-library-mvc/internal/users/models"
)

type LoanService struct {
	loanRepo    loanModels.LoanRepository
	bookService bookModels.BookService
	userService userModels.UserService
}

func NewLoanService(
	loanRepo loanModels.LoanRepository,
	bookService bookModels.BookService,
	userService userModels.UserService,
) loanModels.LoanService {
	return &LoanService{
		loanRepo:    loanRepo,
		bookService: bookService,
		userService: userService,
	}
}

func (l *LoanService) CreateLoan(bookId int64, userId int64) (*loanModels.Loan, error) {
	book, err := l.bookService.GetBook(bookId)
	if err != nil {
		return nil, err
	}

	if book.Quantity <= 0 {
		return nil, errors.New("This book is out of stock right now.")
	}

	_, err = l.userService.GetUser(userId)
	if err != nil {
		return nil, err
	}

	activeLoans, err := l.GetUserLoan(userId)
	if err != nil {
		return nil, err
	}

	if len(activeLoans) > 0 {
		return nil, errors.New("User has active loans")
	}

	book.Quantity = book.Quantity - 1
	err = l.bookService.UpdateBook(book.Id, book)
	if err != nil {
		return nil, err
	}

	loan := &loanModels.Loan{
		UserId:     userId,
		BookId:     bookId,
		BorrowedAt: time.Now(),
		Status:     models.Active,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = l.loanRepo.CreateLoan(*loan)
	if err != nil {
		return nil, err
	}
	return loan, err
}

func (l *LoanService) ReturnBook(loanId int64) error {
	loan, err := l.GetLoan(loanId)
	if err != nil {
		return err
	}

	book, err := l.bookService.GetBook(loan.BookId)
	if err != nil {
		return err
	}

	book.Quantity = book.Quantity + 1
	if err = l.bookService.UpdateBook(book.Id, book); err != nil {
		return err
	}

	loan.Status = models.Returned
	loan.ReturnedAt = time.Now()
	loan.UpdatedAt = time.Now()

	if err = l.loanRepo.UpdateLoan(loan); err != nil {
		return err
	}

	return l.loanRepo.ReturnBook(loanId)
}

func (l *LoanService) GetAllLoan() ([]*loanModels.Loan, error) {
	return l.loanRepo.GetAllLoan()
}

func (l *LoanService) GetLoan(id int64) (*loanModels.Loan, error) {
	return l.loanRepo.GetLoan(id)
}

func (l *LoanService) GetUserLoan(userId int64) ([]*loanModels.Loan, error) {
	return l.loanRepo.GetActiveUserLoan(userId)
}
