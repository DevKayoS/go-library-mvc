package repository

import (
	"errors"
	"sync"

	"github.com/DevKayoS/go-library-mvc/internal/users/models"
)

type UserRepository struct {
	users  map[int64]*models.User
	mu     sync.RWMutex
	nextID int64
}

func NewUserRepository() models.UserRepository {
	return &UserRepository{
		users:  make(map[int64]*models.User),
		nextID: 1,
	}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, value := range u.users {
		if value.Email == user.Email {
			return errors.New("User with this email already exists")
		}
	}

	user.Id = u.nextID
	u.nextID++

	u.users[user.Id] = user
	return nil
}

func (u *UserRepository) DeleteUser(id int64) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exits := u.users[id]
	if !exits {
		return errors.New("user not found")
	}

	delete(u.users, user.Id)

	return nil
}

func (u *UserRepository) GetAllUser() ([]*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	users := make([]*models.User, 0, len(u.users))
	for _, user := range u.users {
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) GetUser(id int64) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.users[id]

	if !exists {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(id int64, user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exits := u.users[id]
	if !exits {
		return errors.New("user not found")
	}

	u.users[user.Id] = user

	return nil
}
