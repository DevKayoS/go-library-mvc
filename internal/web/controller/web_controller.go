package controller

import (
	"html/template"
	"net/http"

	bookService "github.com/DevKayoS/go-library-mvc/internal/books/models"
	loanModels "github.com/DevKayoS/go-library-mvc/internal/loans/models"
	userService "github.com/DevKayoS/go-library-mvc/internal/users/models"
	"github.com/gin-gonic/gin"
)

type WebController struct {
	templates   *template.Template
	bookService bookService.BookService
	userService userService.UserService
	loanService loanModels.LoanService
}

func NewWebController(
	bookService bookService.BookService,
	userService userService.UserService,
	loanService loanModels.LoanService,
) *WebController {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	return &WebController{
		templates:   tmpl,
		bookService: bookService,
		userService: userService,
		loanService: loanService,
	}
}

func (wc *WebController) RegisterRoutes(router *gin.Engine) {
	router.GET("/", wc.ServeHome)
}

func (wc *WebController) ServeHome(ctx *gin.Context) {
	books, _ := wc.bookService.GetAllBook()
	users, _ := wc.userService.GetAllUser()
	loans, _ := wc.loanService.GetAllLoan()

	activeLoans := 0
	for _, loan := range loans {
		if loan.Status == loanModels.Active {
			activeLoans++
		}
	}

	availableBooks := 0
	for _, book := range books {
		if book.Quantity > 0 {
			availableBooks++
		}
	}

	flashMessage, flashMessageType := wc.getFlashMessage(ctx)
	data := map[string]interface{}{
		"Title":         "Sistema de Biblioteca",
		"Books":         books,
		"Users":         users,
		"Loans":         loans,
		"ActiveSection": "dashboard",
		"FlashMessage":  flashMessage,
		"FlashType":     flashMessageType,
		"Stats": map[string]interface{}{
			"TotalBooks":     len(books),
			"TotalUsers":     len(users),
			"TotalLoans":     len(loans),
			"ActiveLoans":    activeLoans,
			"AvailableBooks": availableBooks,
		},
	}
	err := wc.templates.ExecuteTemplate(ctx.Writer, "layout", data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Algo deu errado, por favor tente novamente mais tarde %v", err)
		return
	}
}

func (wc *WebController) addFlashMessage(ctx *gin.Context, message, messageType string) {
	ctx.SetCookie("flash_message", message, 1, "/", "", false, true)
	ctx.SetCookie("flash_type", messageType, 1, "/", "", false, true)
}

func (wc *WebController) getFlashMessage(ctx *gin.Context) (string, string) {
	message, _ := ctx.Cookie("flash_message")
	messageType, _ := ctx.Cookie("flash_type")

	ctx.SetCookie("flash_message", "", 1, "/", "", false, true)
	ctx.SetCookie("flash_type", "", 1, "/", "", false, true)

	return message, messageType
}
