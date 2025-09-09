package controllers

import "github.com/gin-gonic/gin"

type BookController struct {
}

func NewBookController() *BookController {
	return &BookController{}
}

func (c *BookController) RegisterRoutes(r *gin.Engine) {
	Books := r.Group("/books")
	{
		Books.POST("", c.CreateBook)
		Books.GET("/:id", c.GetBook)
		Books.GET("", c.GetAllBook)
		Books.PUT("/:id", c.UpdateBook)
		Books.DELETE("/:id", c.UpdateBook)
	}
}

func (c *BookController) CreateBook(ctx *gin.Context) {
}

func (c *BookController) GetBook(ctx *gin.Context) {
}

func (c *BookController) GetAllBook(ctx *gin.Context) {
	ctx.String(200, "FUNCIONOU")
}

func (c *BookController) UpdateBook(ctx *gin.Context) {

}

func (c *BookController) DeleteBook(ctx *gin.Context) {

}
