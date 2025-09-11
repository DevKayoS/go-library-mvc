package models

type UserRepository interface {
	CreateUser(user *User) error
	GetUser(id int64) (*User, error)
	GetAllUser() ([]*User, error)
	UpdateUser(id int64, user *User) error
	DeleteUser(id int64) error
}
