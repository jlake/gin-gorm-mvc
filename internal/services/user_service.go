package services

import (
	"errors"
	"gin-gorm-mvc/internal/models"
	"gin-gorm-mvc/internal/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers(page, pageSize int) ([]models.User, int64, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	VerifyPassword(user *models.User, password string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// ユーザー名の重複チェック
	_, err = s.repo.FindByUsername(user.Username)
	if err == nil {
		return errors.New("username already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// メールアドレスの重複チェック
	_, err = s.repo.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) GetAllUsers(page, pageSize int) ([]models.User, int64, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *userService) UpdateUser(user *models.User) error {
	// パスワードが変更された場合のみハッシュ化
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) VerifyPassword(user *models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
