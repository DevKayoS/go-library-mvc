package controllers

import (
	"net/http"
	"strconv"

	"github.com/DevKayoS/go-library-mvc/internal/books/models"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService models.BookService
}

func NewBookController(bookService models.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (b *BookController) RegisterRoutes(r *gin.Engine) {
	Books := r.Group("/books")
	{
		Books.POST("", b.CreateBook)
		Books.GET("/:id", b.GetBook)
		Books.GET("", b.GetAllBook)
		Books.PUT("/:id", b.UpdateBook)
		Books.DELETE("/:id", b.UpdateBook)
	}
}

func (b *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Request Body",
		})
		return
	}

	err := b.bookService.CreateBook(&book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"code":    http.StatusCreated,
		"message": "Created successfully",
		"data":    book,
	})
}

func (b *BookController) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Book Id",
		})
		return
	}

	book, err := b.bookService.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"code":   http.StatusOK,
		"data":   book,
	})
}

func (b *BookController) GetAllBook(ctx *gin.Context) {
	book, err := b.bookService.GetAllBook()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"code":   http.StatusOK,
		"data":   book,
	})
}

func (b *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Book Id",
		})
		return
	}

	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Request Body",
		})
		return
	}

	err = b.bookService.UpdateBook(id, &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"code":    http.StatusOK,
		"message": "Updated successfully",
	})
}

func (b *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Book Id",
		})
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"code":    http.StatusOK,
		"message": "Book deleted",
	})
}
