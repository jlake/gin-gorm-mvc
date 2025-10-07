package services

import (
	"gin-gorm-mvc/internal/models"
	"gin-gorm-mvc/internal/repositories"
)

type ArticleService interface {
	CreateArticle(article *models.Article) error
	GetArticleByID(id uint) (*models.Article, error)
	GetAllArticles(page, pageSize int) ([]models.Article, int64, error)
	GetArticlesByAuthor(authorID uint, page, pageSize int) ([]models.Article, int64, error)
	UpdateArticle(article *models.Article) error
	DeleteArticle(id uint) error
	IncrementViewCount(id uint) error
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) CreateArticle(article *models.Article) error {
	return s.repo.Create(article)
}

func (s *articleService) GetArticleByID(id uint) (*models.Article, error) {
	return s.repo.FindByID(id)
}

func (s *articleService) GetAllArticles(page, pageSize int) ([]models.Article, int64, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *articleService) GetArticlesByAuthor(authorID uint, page, pageSize int) ([]models.Article, int64, error) {
	return s.repo.FindByAuthorID(authorID, page, pageSize)
}

func (s *articleService) UpdateArticle(article *models.Article) error {
	return s.repo.Update(article)
}

func (s *articleService) DeleteArticle(id uint) error {
	return s.repo.Delete(id)
}

func (s *articleService) IncrementViewCount(id uint) error {
	return s.repo.IncrementViewCount(id)
}
