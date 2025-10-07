package routes

import (
	"gin-gorm-mvc/internal/controllers"
	"gin-gorm-mvc/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.Engine, userCtrl *controllers.UserController, articleCtrl *controllers.ArticleController, adminCtrl *controllers.AdminController) {
	// HTMLテンプレートの読み込み
	r.LoadHTMLGlob("internal/views/**/*")

	// ミドルウェアの適用
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// トップページ
	r.GET("/", adminCtrl.Index)
	r.GET("/admin", adminCtrl.Index)

	// 管理画面HTMLページ
	r.GET("/admin/users", adminCtrl.Users)

	// JSON API
	admin := r.Group("/admin/api")
	{
		// ユーザー管理API
		users := admin.Group("/users")
		{
			users.POST("", userCtrl.CreateUser)
			users.GET("", userCtrl.GetAllUsers)
			users.GET("/:id", userCtrl.GetUser)
			users.PUT("/:id", userCtrl.UpdateUser)
			users.DELETE("/:id", userCtrl.DeleteUser)
		}

		// 記事管理API
		articles := admin.Group("/articles")
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
			"module": "admin",
		})
	})
}
