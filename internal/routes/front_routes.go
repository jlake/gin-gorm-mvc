package routes

import (
	"gin-gorm-mvc/internal/controllers"
	"gin-gorm-mvc/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupFrontRoutes(r *gin.Engine, userCtrl *controllers.UserController, articleCtrl *controllers.ArticleController) {
	// ミドルウェアの適用
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// フロントエンドルート
	front := r.Group("/front")
	{
		// 記事の閲覧（公開のみ）
		front.GET("/articles", articleCtrl.GetAllArticles)
		front.GET("/articles/:id", articleCtrl.GetArticle)
		front.GET("/articles/author/:author_id", articleCtrl.GetArticlesByAuthor)

		// ユーザー情報の閲覧
		front.GET("/users/:id", userCtrl.GetUser)
	}

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"module": "front",
		})
	})
}
