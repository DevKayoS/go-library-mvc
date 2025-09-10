package controllers

import (
	"net/http"
	"strconv"

	"github.com/DevKayoS/go-library-mvc/internal/users/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService models.UserService
}

func NewUserController(userService models.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
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
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Request Body",
		})
		return
	}

	err := c.userService.CreateUser(&user)

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
		"data":    user,
	})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid User Id",
		})
		return
	}

	user, err := c.userService.GetUser(id)
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
		"data":   user,
	})
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	user, err := c.userService.GetAllUser()

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
		"data":   user,
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid User Id",
		})
		return
	}

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid Request Body",
		})
		return
	}

	err = c.userService.UpdateUser(id, &user)
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

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"code":   http.StatusBadRequest,
			"error":  "Invalid User Id",
		})
		return
	}

	err = c.userService.DeleteUser(id)
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
		"message": "User deleted",
	})
}
