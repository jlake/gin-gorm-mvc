package controllers

import (
	"gin-gorm-mvc/internal/services"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	userService    services.UserService
	articleService services.ArticleService
}

func NewAdminController(userService services.UserService, articleService services.ArticleService) *AdminController {
	return &AdminController{
		userService:    userService,
		articleService: articleService,
	}
}

// Index 管理画面トップページを表示
func (ctrl *AdminController) Index(c *gin.Context) {
	c.HTML(200, "admin/index.html", gin.H{
		"title": "管理画面",
	})
}

// Users ユーザー一覧ページを表示
func (ctrl *AdminController) Users(c *gin.Context) {
	users, _, err := ctrl.userService.GetAllUsers(1, 100)
	if err != nil {
		c.HTML(500, "admin/users.html", gin.H{
			"error": "ユーザーの取得に失敗しました",
			"users": nil,
		})
		return
	}

	c.HTML(200, "admin/users.html", gin.H{
		"users": users,
	})
}
