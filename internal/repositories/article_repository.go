package repositories

import (
	"gin-gorm-mvc/internal/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *models.Article) error
	FindByID(id uint) (*models.Article, error)
	FindAll(page, pageSize int) ([]models.Article, int64, error)
	FindByAuthorID(authorID uint, page, pageSize int) ([]models.Article, int64, error)
	Update(article *models.Article) error
	Delete(id uint) error
	IncrementViewCount(id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(article *models.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) FindByID(id uint) (*models.Article, error) {
	var article models.Article
	err := r.db.Preload("Author").First(&article, id).Error
	return &article, err
}

func (r *articleRepository) FindAll(page, pageSize int) ([]models.Article, int64, error) {
	var articles []models.Article
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&models.Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Author").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&articles).Error
	return articles, total, err
}

func (r *articleRepository) FindByAuthorID(authorID uint, page, pageSize int) ([]models.Article, int64, error) {
	var articles []models.Article
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&models.Article{}).Where("author_id = ?", authorID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Author").Where("author_id = ?", authorID).Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&articles).Error
	return articles, total, err
}

func (r *articleRepository) Update(article *models.Article) error {
	return r.db.Save(article).Error
}

func (r *articleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Article{}, id).Error
}

func (r *articleRepository) IncrementViewCount(id uint) error {
	return r.db.Model(&models.Article{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}
