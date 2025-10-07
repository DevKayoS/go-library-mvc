package controller

import (
	"html/template"
	"net/http"
	"strconv"

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
	router.GET("/users", wc.ServeUser)
	router.POST("/users", wc.CreateUser)
	router.POST("/users/:id/delete", wc.DeleteUser)
	router.POST("/users/:id/edit", wc.UpdateUser)
	router.GET("/users/:id/edit", wc.EditUser)

	router.GET("/books", wc.ServeBook)
	router.POST("/books", wc.CreateBook)
	router.POST("/books/:id/delete", wc.DeleteBook)
	router.POST("/books/:id/edit", wc.UpdateBook)
	router.GET("/books/:id/edit", wc.EditBook)

	router.GET("/loans", wc.ServeLoan)
	router.POST("/loans", wc.CreateLoan)
	router.POST("/loans/:id/delete", wc.DeleteLoan)
	router.POST("/loans/:id/edit", wc.UpdateLoan)
	router.GET("/loans/:id/edit", wc.EditLoan)
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

func (wc *WebController) ServeUser(ctx *gin.Context) {
	users, _ := wc.userService.GetAllUser()

	flashMessage, flashMessageType := wc.getFlashMessage(ctx)

	data := map[string]interface{}{
		"Title":         "Gerenciamento de Usuarios!",
		"Users":         users,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashMessageType,
	}

	err := wc.templates.ExecuteTemplate(ctx.Writer, "layout", data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Algo deu errado, por favor tente novamente mais tarde %v", err)
		return
	}
}

func (wc *WebController) CreateUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")

	user := userService.User{
		Name:  name,
		Email: email,
	}

	err := wc.userService.CreateUser(&user)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao criar usuario: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(ctx, "Sucesso ao criar usuario", "success")
	}

	ctx.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) DeleteUser(ctx *gin.Context) {
	userIDString := ctx.Param("id")

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Id do usuario invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/users")
		return
	}

	err = wc.userService.DeleteUser(userID)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao deletar usuario: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(ctx, "Sucesso ao deletar usuario", "success")
	}

	ctx.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) UpdateUser(ctx *gin.Context) {
	userIDString := ctx.Param("id")

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Id do usuario invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/users")
		return
	}

	user, err := wc.userService.GetUser(userID)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao tentar pegar o usuario: "+err.Error(), "error")
		ctx.Redirect(http.StatusSeeOther, "/users")
		return
	}

	name := ctx.PostForm("name")
	email := ctx.PostForm("email")

	user.Name = name
	user.Email = email

	err = wc.userService.UpdateUser(userID, user)

	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao atualizar usuario: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(ctx, "Sucesso ao atualizar usuario", "success")
	}

	ctx.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) EditUser(ctx *gin.Context) {
	userIDString := ctx.Param("id")

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Id do usuario invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/users")
		return
	}

	user, err := wc.userService.GetUser(userID)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao tentar pegar o usuario: "+err.Error(), "error")
		ctx.Redirect(http.StatusSeeOther, "/users")
		return
	}

	flashMessage, flashMessageType := wc.getFlashMessage(ctx)

	data := map[string]interface{}{
		"Title":         "Editar Usuarios!",
		"User":          user,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashMessageType,
		"IsEdit":        true,
	}

	err = wc.templates.ExecuteTemplate(ctx.Writer, "layout", data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Algo deu errado, por favor tente novamente mais tarde %v", err)
		return
	}
}

func (wc *WebController) ServeBook(ctx *gin.Context) {
	books, _ := wc.bookService.GetAllBook()

	flashMessage, flashMessageType := wc.getFlashMessage(ctx)

	data := map[string]interface{}{
		"Title":         "Gerenciamento de Livros!",
		"Books":         books,
		"ActiveSection": "books",
		"FlashMessage":  flashMessage,
		"FlashType":     flashMessageType,
	}

	err := wc.templates.ExecuteTemplate(ctx.Writer, "layout", data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Algo deu errado, por favor tente novamente mais tarde %v", err)
		return
	}
}

func (wc *WebController) CreateBook(ctx *gin.Context) {
	author := ctx.PostForm("author")
	title := ctx.PostForm("title")
	quantityString := ctx.PostForm("quantity")

	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		wc.addFlashMessage(ctx, "A quantidade e um valor invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book := bookService.Book{
		Title:    title,
		Author:   author,
		Quantity: quantity,
	}

	err = wc.bookService.CreateBook(&book)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao criar usuario: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(ctx, "Sucesso ao criar usuario", "success")
	}

	ctx.Redirect(http.StatusSeeOther, "/books")
}

func (wc *WebController) DeleteBook(ctx *gin.Context) {
	bookIDString := ctx.Param("id")

	bookID, err := strconv.ParseInt(bookIDString, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Id do livro invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/books")
		return
	}

	err = wc.bookService.DeleteBook(bookID)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao deletar livro: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(ctx, "Sucesso ao deletar livro", "success")
	}

	ctx.Redirect(http.StatusSeeOther, "/books")
}

func (wc *WebController) UpdateBook(ctx *gin.Context) {
	bookIDString := ctx.Param("id")

	bookID, err := strconv.ParseInt(bookIDString, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Id do livro invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book, err := wc.bookService.GetBook(bookID)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao tentar pegar o livro: "+err.Error(), "error")
		ctx.Redirect(http.StatusSeeOther, "/books")
		return
	}

	author := ctx.PostForm("author")
	title := ctx.PostForm("title")
	quantityString := ctx.PostForm("quantity")

	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		wc.addFlashMessage(ctx, "A quantidade e um valor invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book.Author = author
	book.Quantity = quantity
	book.Title = title

	err = wc.bookService.UpdateBook(bookID, book)

	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao atualizar o livro: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(ctx, "Sucesso ao atualizar o livro", "success")
	}

	ctx.Redirect(http.StatusSeeOther, "/books")
}

func (wc *WebController) EditBook(ctx *gin.Context) {
	bookIDString := ctx.Param("id")

	bookID, err := strconv.ParseInt(bookIDString, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Id do livro invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book, err := wc.bookService.GetBook(bookID)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao tentar pegar o livro: "+err.Error(), "error")
		ctx.Redirect(http.StatusSeeOther, "/books")
		return
	}

	flashMessage, flashMessageType := wc.getFlashMessage(ctx)

	data := map[string]interface{}{
		"Title":         "Editar Livro!",
		"Book":          book,
		"ActiveSection": "books",
		"FlashMessage":  flashMessage,
		"FlashType":     flashMessageType,
		"IsEdit":        true,
	}

	err = wc.templates.ExecuteTemplate(ctx.Writer, "layout", data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Algo deu errado, por favor tente novamente mais tarde %v", err)
		return
	}
}

func (wc *WebController) ServeLoan(ctx *gin.Context) {
	loans, _ := wc.loanService.GetAllLoan()
	books, _ := wc.bookService.GetAllBook()
	users, _ := wc.userService.GetAllUser()

	flashMessage, flashMessageType := wc.getFlashMessage(ctx)

	data := map[string]interface{}{
		"Title":         "Gerenciamento de Emprestimos!",
		"Loans":         loans,
		"Books":         books,
		"Users":         users,
		"ActiveSection": "loans",
		"FlashMessage":  flashMessage,
		"FlashType":     flashMessageType,
	}

	err := wc.templates.ExecuteTemplate(ctx.Writer, "layout", data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Algo deu errado, por favor tente novamente mais tarde %v", err)
		return
	}
}

func (wc *WebController) CreateLoan(ctx *gin.Context) {
	userIDStr := ctx.PostForm("user_id")
	bookIDStr := ctx.PostForm("book_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Usuario invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(ctx, "Livro invalido", "error")
		ctx.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	_, err = wc.loanService.CreateLoan(bookID, userID)
	if err != nil {
		wc.addFlashMessage(ctx, "Erro ao criar emprestimo: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(ctx, "Sucesso ao criar emprestimo", "success")
	}

	ctx.Redirect(http.StatusSeeOther, "/loans")
}

func (wc *WebController) UpdateLoan(ctx *gin.Context) {
}

func (wc *WebController) EditLoan(ctx *gin.Context) {
}

func (wc *WebController) DeleteLoan(ctx *gin.Context) {
}
