package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", c.CreateUser)
		users.GET("/:id", c.GetUser)
		users.GET("", c.GetAllUser)
		users.PUT("/:id", c.UpdateUser)
		users.DELETE("/:id", c.UpdateUser)
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
}

func (c *UserController) GetUser(ctx *gin.Context) {
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	ctx.String(200, "FUNCIONOU")
}

func (c *UserController) UpdateUser(ctx *gin.Context) {

}

func (c *UserController) DeleteUser(ctx *gin.Context) {

}
