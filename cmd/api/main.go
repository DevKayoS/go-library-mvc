package main

import (
	"log"

	"github.com/DevKayoS/go-library-mvc/internal/users/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userController := controllers.NewUserController()
	userController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
