package models

import "time"

type Loan struct {
	ID         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	BookId     int64     `json:"book_id"`
	BorrowedAt time.Time `json:"borrowedAt"`
	ReturnedAt time.Time `json:"returnedAt"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
