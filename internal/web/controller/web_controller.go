package controller

import (
	"html/template"

	bookService "github.com/DevKayoS/go-library-mvc/internal/books/services"
	loanService "github.com/DevKayoS/go-library-mvc/internal/loans/services"
	userService "github.com/DevKayoS/go-library-mvc/internal/users/services"
)

type WebController struct {
	templates   *template.Template
	bookService bookService.BookService
	userService userService.UserService
	loanService loanService.LoanService
}

func NewWebController(
	bookService bookService.BookService,
	userService userService.UserService,
	loanService loanService.LoanService,
) *WebController {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	return &WebController{
		templates:   tmpl,
		bookService: bookService,
		userService: userService,
		loanService: loanService,
	}
}
