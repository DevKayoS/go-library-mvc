package models

type LoanRepository interface {
	CreateLoan(loan Loan) error
	ReturnBook(loanId int64) error
	GetLoan(id int64) (*Loan, error)
	GetActiveUserLoan(userId int64) ([]*Loan, error)
	GetAllLoan() ([]*Loan, error)
}
