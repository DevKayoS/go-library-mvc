package services

import (
	"time"

	"github.com/DevKayoS/go-library-mvc/internal/users/models"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(repository models.UserRepository) models.UserService {
	return &UserService{
		userRepo: repository,
	}
}

func (u *UserService) CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return u.userRepo.CreateUser(user)
}

func (u *UserService) DeleteUser(id int64) error {
	return u.userRepo.DeleteUser(id)
}

func (u *UserService) GetAllUser() ([]*models.User, error) {
	return u.userRepo.GetAllUser()
}

func (u *UserService) GetUser(id int64) (*models.User, error) {
	return u.userRepo.GetUser(id)
}

func (u *UserService) UpdateUser(id int64, user *models.User) error {
	user.UpdatedAt = time.Now()
	return u.userRepo.UpdateUser(id, user)
}
