package controllers

import (
	"gin-gorm-mvc/internal/models"
	"gin-gorm-mvc/internal/services"
	"gin-gorm-mvc/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

// CreateUser 新しいユーザーを作成
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := ctrl.service.CreateUser(&user); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "User created successfully", user)
}

// GetUser IDでユーザーを取得
func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	user, err := ctrl.service.GetUserByID(uint(id))
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	response.Success(c, user)
}

// GetAllUsers すべてのユーザーを取得
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := ctrl.service.GetAllUsers(page, pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve users")
		return
	}

	response.SuccessPaginated(c, users, total, page, pageSize)
}

// UpdateUser ユーザー情報を更新
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user.ID = uint(id)
	if err := ctrl.service.UpdateUser(&user); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "User updated successfully", user)
}

// DeleteUser ユーザーを削除
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	if err := ctrl.service.DeleteUser(uint(id)); err != nil {
		response.InternalServerError(c, "Failed to delete user")
		return
	}

	response.SuccessWithMessage(c, "User deleted successfully", nil)
}
