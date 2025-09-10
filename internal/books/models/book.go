package models

import "time"

type Book struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Quantity  int       `json:"qtd"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
