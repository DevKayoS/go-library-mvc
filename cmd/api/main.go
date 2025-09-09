package main

import (
	"log"

	bookController "github.com/DevKayoS/go-library-mvc/internal/books/controllers"
	userController "github.com/DevKayoS/go-library-mvc/internal/users/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// User
	userController := userController.NewUserController()
	userController.RegisterRoutes(router)

	// Book
	bookController := bookController.NewBookController()
	bookController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
