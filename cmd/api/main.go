package main

import (
	"log"

	bookService "github.com/DevKayoS/go-library-mvc/internal/books/services"
	loanService "github.com/DevKayoS/go-library-mvc/internal/loans/services"
	userService "github.com/DevKayoS/go-library-mvc/internal/users/services"
	"github.com/gin-contrib/cors"

	bookRepository "github.com/DevKayoS/go-library-mvc/internal/books/repositories"
	loanRepository "github.com/DevKayoS/go-library-mvc/internal/loans/repositories"
	userRepository "github.com/DevKayoS/go-library-mvc/internal/users/repositories"

	bookController "github.com/DevKayoS/go-library-mvc/internal/books/controllers"
	loanController "github.com/DevKayoS/go-library-mvc/internal/loans/controllers"
	userController "github.com/DevKayoS/go-library-mvc/internal/users/controllers"
	webController "github.com/DevKayoS/go-library-mvc/internal/web/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// users
	userRepository := userRepository.NewUserRepository()
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)
	userController.RegisterRoutes(router)

	// Book
	bookRepository := bookRepository.NewBookRepository()
	bookService := bookService.NewBookService(bookRepository)
	bookController := bookController.NewBookController(bookService)
	bookController.RegisterRoutes(router)

	// Loan
	loanRepository := loanRepository.NewLoanRepository()
	loanService := loanService.NewLoanService(loanRepository, bookService, userService)
	loanController := loanController.NewLoanController(loanService)
	loanController.RegisterRoutes(router)

	// Web
	webController := webController.NewWebController(bookService, userService, loanService)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
