package controllers

import (
	"net/http"
	"strconv"

	"github.com/DevKayoS/go-library-mvc/internal/loans/models"
	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanService models.LoanService
}

func NewLoanController(loanService models.LoanService) *LoanController {
	return &LoanController{
		loanService: loanService,
	}
}

func (l *LoanController) RegisterRoutes(r *gin.RouterGroup) {
	loans := r.Group("/loans")
	{
		loans.POST("", l.CreateLoan)
		loans.GET("/:id", l.GetLoan)
		loans.GET("", l.GetAllLoan)
		loans.PUT("/:id/return", l.ReturnBook)
	}
	user := r.Group("/loans/users")
	{
		user.GET("/:userId", l.GetUserLoan)
	}
}

func (l *LoanController) CreateLoan(ctx *gin.Context) {
	var request struct {
		BookId int64 `json:"book_id"`
		UserId int64 `json:"user_id"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Request Body",
		})
		return
	}

	loan, err := l.loanService.CreateLoan(request.BookId, request.UserId)
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
		"data":    loan,
	})
}

func (l *LoanController) GetLoan(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Loan Id",
		})
		return
	}

	loan, err := l.loanService.GetLoan(id)
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
		"data":   loan,
	})
}

func (l *LoanController) GetAllLoan(ctx *gin.Context) {
	loan, err := l.loanService.GetAllLoan()
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
		"data":   loan,
	})
}

func (l *LoanController) ReturnBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Loan Id",
		})
		return
	}

	err = l.loanService.ReturnBook(id)
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
		"message": "Book has returned successfully",
	})
}

func (l *LoanController) GetUserLoan(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid User Id",
		})
		return
	}

	loan, err := l.loanService.GetUserLoan(userId)
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
		"data":   loan,
	})
}
