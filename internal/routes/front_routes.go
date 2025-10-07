package routes

import (
	"gin-gorm-mvc/internal/controllers"
	"gin-gorm-mvc/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupFrontRoutes(r *gin.Engine, userCtrl *controllers.UserController, articleCtrl *controllers.ArticleController, frontCtrl *controllers.FrontController) {
	// HTMLテンプレートの読み込み
	r.LoadHTMLGlob("internal/views/**/*")

	// ミドルウェアの適用
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// トップページ
	r.GET("/", frontCtrl.Index)

	// フロントエンドルート
	front := r.Group("/front")
	{
		// HTMLページ
		front.GET("/articles", frontCtrl.Articles)

		// JSON API（既存）
		front.GET("/api/articles", articleCtrl.GetAllArticles)
		front.GET("/api/articles/:id", articleCtrl.GetArticle)
		front.GET("/api/articles/author/:author_id", articleCtrl.GetArticlesByAuthor)
		front.GET("/api/users/:id", userCtrl.GetUser)
	}

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"module": "front",
		})
	})
}
