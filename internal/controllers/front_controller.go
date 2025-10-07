package controllers

import (
	"gin-gorm-mvc/internal/services"

	"github.com/gin-gonic/gin"
)

type FrontController struct {
	userService    services.UserService
	articleService services.ArticleService
}

func NewFrontController(userService services.UserService, articleService services.ArticleService) *FrontController {
	return &FrontController{
		userService:    userService,
		articleService: articleService,
	}
}

// Index トップページを表示
func (ctrl *FrontController) Index(c *gin.Context) {
	c.HTML(200, "front/index.html", gin.H{
		"title": "ホーム",
	})
}

// Articles 記事一覧ページを表示
func (ctrl *FrontController) Articles(c *gin.Context) {
	articles, _, err := ctrl.articleService.GetAllArticles(1, 100)
	if err != nil {
		c.HTML(500, "front/articles.html", gin.H{
			"error":    "記事の取得に失敗しました",
			"articles": nil,
		})
		return
	}

	c.HTML(200, "front/articles.html", gin.H{
		"articles": articles,
	})
}
