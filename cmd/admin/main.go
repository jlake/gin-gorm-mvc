package main

import (
	"fmt"
	"gin-gorm-mvc/internal/config"
	"gin-gorm-mvc/internal/controllers"
	"gin-gorm-mvc/internal/database"
	"gin-gorm-mvc/internal/redis"
	"gin-gorm-mvc/internal/repositories"
	"gin-gorm-mvc/internal/routes"
	"gin-gorm-mvc/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設定の読み込み
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// データベース接続
	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Redis接続
	if err := redis.Initialize(); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	defer redis.Close()

	// 依存性注入
	db := database.GetDB()

	// リポジトリの初期化
	userRepo := repositories.NewUserRepository(db)
	articleRepo := repositories.NewArticleRepository(db)

	// サービスの初期化
	userService := services.NewUserService(userRepo)
	articleService := services.NewArticleService(articleRepo)

	// コントローラの初期化
	userCtrl := controllers.NewUserController(userService)
	articleCtrl := controllers.NewArticleController(articleService)
	adminCtrl := controllers.NewAdminController(userService, articleService)

	// Ginエンジンの初期化
	r := gin.Default()

	// ルートの設定
	routes.SetupAdminRoutes(r, userCtrl, articleCtrl, adminCtrl)

	// サーバー起動
	port := fmt.Sprintf(":%s", config.AppConfig.Server.AdminPort)
	log.Printf("Admin Server starting on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
