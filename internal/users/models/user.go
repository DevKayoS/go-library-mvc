package models

import "time"

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" binding:"required,min=3,max=255"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
