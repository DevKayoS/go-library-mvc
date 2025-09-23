package repositories

import (
	"errors"
	"sync"

	"github.com/DevKayoS/go-library-mvc/internal/loans/models"
)

type LoanRepository struct {
	loans  map[int64]*models.Loan
	mu     sync.RWMutex
	nextID int64
}

func NewLoanRepository() models.LoanRepository {
	return &LoanRepository{
		loans:  make(map[int64]*models.Loan),
		nextID: 1,
	}
}

func (l *LoanRepository) CreateLoan(loan models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan.ID = l.nextID
	l.nextID++

	l.loans[loan.ID] = &loan
	return nil
}

func (l *LoanRepository) GetActiveUserLoan(userId int64) ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var activeLoan []*models.Loan

	for _, loan := range l.loans {
		if loan.UserId == userId && loan.Status == models.Active {
			activeLoan = append(activeLoan, loan)
		}
	}

	return activeLoan, nil
}

func (l *LoanRepository) GetAllLoan() ([]*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loans := make([]*models.Loan, 0, len(l.loans))
	for _, loan := range l.loans {
		loans = append(loans, loan)
	}

	return loans, nil
}

func (l *LoanRepository) GetLoan(id int64) (*models.Loan, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loans[id]

	if !exists {
		return nil, errors.New("loan not found")
	}

	return loan, nil
}

func (l *LoanRepository) ReturnBook(loanId int64) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loans[loanId]

	if !exists {
		return errors.New("loan not found")
	}

	loan.Status = models.Returned

	l.loans[loanId] = loan

	return nil
}

func (l *LoanRepository) UpdateLoan(loan *models.Loan) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	loan, exists := l.loans[loan.ID]
	if !exists {
		return errors.New("loan not found")
	}

	l.loans[loan.ID] = loan

	return nil
}
