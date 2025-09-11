package models

import "time"

type Book struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title" binding:"required, min=3"`
	Author    string    `json:"author" binding:"required,min=3"`
	Quantity  int       `json:"qtd" binding:"required,min=1"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
