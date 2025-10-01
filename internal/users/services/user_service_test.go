package services

import (
	"testing"

	"github.com/DevKayoS/go-library-mvc/internal/users/models"
	repository "github.com/DevKayoS/go-library-mvc/internal/users/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	repository := repository.NewUserRepository()
	service := NewUserService(repository)

	user := &models.User{Name: "teste", Email: "teste@teste.com"}
	err := service.CreateUser(user)

	assert.Equal(t, err, nil)

	userSaved, err := service.GetUser(1)

	assert.Equal(t, "teste", userSaved.Name)
}
