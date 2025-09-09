package controllers

import "github.com/gin-gonic/gin"

type LoanController struct {
}

func NewLoanController() *LoanController {
	return &LoanController{}
}

func (c *LoanController) RegisterRoutes(r *gin.Engine) {
	Loans := r.Group("/loans")
	{
		Loans.POST("", c.CreateLoan)
		Loans.GET("/:id", c.GetLoan)
		Loans.GET("", c.GetAllLoan)
		Loans.PUT("/:id", c.UpdateLoan)
		Loans.DELETE("/:id", c.UpdateLoan)
	}
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {
}

func (c *LoanController) GetLoan(ctx *gin.Context) {
}

func (c *LoanController) GetAllLoan(ctx *gin.Context) {
	ctx.String(200, "FUNCIONOU")
}

func (c *LoanController) UpdateLoan(ctx *gin.Context) {

}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {

}
