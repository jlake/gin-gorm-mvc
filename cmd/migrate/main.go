package main

import (
	"gin-gorm-mvc/internal/config"
	"gin-gorm-mvc/internal/database"
	"gin-gorm-mvc/internal/models"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration tool",
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		runMigrations()
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback migrations",
	Run: func(cmd *cobra.Command, args []string) {
		rollbackMigrations()
	},
}

var freshCmd = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run migrations",
	Run: func(cmd *cobra.Command, args []string) {
		dropTables()
		runMigrations()
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(freshCmd)
}

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

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func runMigrations() {
	db := database.GetDB()

	log.Println("Running migrations...")

	// AutoMigrateを使用してテーブルを作成・更新
	err := db.AutoMigrate(
		&models.User{},
		&models.Article{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations completed successfully")
}

func rollbackMigrations() {
	db := database.GetDB()

	log.Println("Rolling back migrations...")

	// テーブルを削除
	err := db.Migrator().DropTable(
		&models.Article{},
		&models.User{},
	)

	if err != nil {
		log.Fatalf("Rollback failed: %v", err)
	}

	log.Println("Rollback completed successfully")
}

func dropTables() {
	db := database.GetDB()

	log.Println("Dropping all tables...")

	// すべてのテーブルを削除
	err := db.Migrator().DropTable(
		&models.Article{},
		&models.User{},
	)

	if err != nil {
		log.Fatalf("Drop tables failed: %v", err)
	}

	log.Println("All tables dropped successfully")
}
