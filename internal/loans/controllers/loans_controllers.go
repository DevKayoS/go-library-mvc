package controllers

import "github.com/gin-gonic/gin"

type LoanController struct {
}

func NewLoanController() *LoanController {
	return &LoanController{}
}

func (l *LoanController) RegisterRoutes(r *gin.Engine) {
	Loans := r.Group("/loans")
	{
		Loans.POST("", l.CreateLoan)
		Loans.GET("/:id", l.GetLoan)
		Loans.GET("", l.GetAllLoan)
		Loans.PUT("/:id", l.UpdateLoan)
		Loans.DELETE("/:id", l.UpdateLoan)
	}
}

func (l *LoanController) CreateLoan(ctx *gin.Context) {
}

func (l *LoanController) GetLoan(ctx *gin.Context) {
}

func (l *LoanController) GetAllLoan(ctx *gin.Context) {
	ctx.String(200, "FUNCIONOU")
}

func (l *LoanController) UpdateLoan(ctx *gin.Context) {

}

func (l *LoanController) DeleteLoan(ctx *gin.Context) {

}
