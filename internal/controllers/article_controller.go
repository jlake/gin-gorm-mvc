package controllers

import (
	"gin-gorm-mvc/internal/models"
	"gin-gorm-mvc/internal/services"
	"gin-gorm-mvc/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	service services.ArticleService
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{service: service}
}

// CreateArticle 新しい記事を作成
func (ctrl *ArticleController) CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := ctrl.service.CreateArticle(&article); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "Article created successfully", article)
}

// GetArticle IDで記事を取得
func (ctrl *ArticleController) GetArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid article ID")
		return
	}

	article, err := ctrl.service.GetArticleByID(uint(id))
	if err != nil {
		response.NotFound(c, "Article not found")
		return
	}

	// 閲覧数をインクリメント
	_ = ctrl.service.IncrementViewCount(uint(id))

	response.Success(c, article)
}

// GetAllArticles すべての記事を取得
func (ctrl *ArticleController) GetAllArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	articles, total, err := ctrl.service.GetAllArticles(page, pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve articles")
		return
	}

	response.SuccessPaginated(c, articles, total, page, pageSize)
}

// GetArticlesByAuthor 著者IDで記事を取得
func (ctrl *ArticleController) GetArticlesByAuthor(c *gin.Context) {
	authorID, err := strconv.ParseUint(c.Param("author_id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid author ID")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	articles, total, err := ctrl.service.GetArticlesByAuthor(uint(authorID), page, pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve articles")
		return
	}

	response.SuccessPaginated(c, articles, total, page, pageSize)
}

// UpdateArticle 記事を更新
func (ctrl *ArticleController) UpdateArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid article ID")
		return
	}

	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	article.ID = uint(id)
	if err := ctrl.service.UpdateArticle(&article); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.SuccessWithMessage(c, "Article updated successfully", article)
}

// DeleteArticle 記事を削除
func (ctrl *ArticleController) DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid article ID")
		return
	}

	if err := ctrl.service.DeleteArticle(uint(id)); err != nil {
		response.InternalServerError(c, "Failed to delete article")
		return
	}

	response.SuccessWithMessage(c, "Article deleted successfully", nil)
}
