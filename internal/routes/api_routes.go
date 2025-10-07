package routes

import (
	"gin-gorm-mvc/internal/controllers"
	"gin-gorm-mvc/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(r *gin.Engine, userCtrl *controllers.UserController, articleCtrl *controllers.ArticleController) {
	// ミドルウェアの適用
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// APIルート
	api := r.Group("/api/v1")
	{
		// ユーザーエンドポイント
		users := api.Group("/users")
		{
			users.POST("", userCtrl.CreateUser)
			users.GET("", userCtrl.GetAllUsers)
			users.GET("/:id", userCtrl.GetUser)
			users.PUT("/:id", userCtrl.UpdateUser)
			users.DELETE("/:id", userCtrl.DeleteUser)
		}

		// 記事エンドポイント
		articles := api.Group("/articles")
		{
			articles.POST("", articleCtrl.CreateArticle)
			articles.GET("", articleCtrl.GetAllArticles)
			articles.GET("/:id", articleCtrl.GetArticle)
			articles.GET("/author/:author_id", articleCtrl.GetArticlesByAuthor)
			articles.PUT("/:id", articleCtrl.UpdateArticle)
			articles.DELETE("/:id", articleCtrl.DeleteArticle)
		}
	}

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"module": "api",
		})
	})
}
