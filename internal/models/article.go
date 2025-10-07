package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title"`
	Content     string         `gorm:"type:text" json:"content"`
	AuthorID    uint           `gorm:"not null" json:"author_id"`
	Author      User           `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Status      int            `gorm:"default:1" json:"status"` // 1: published, 0: draft
	ViewCount   int            `gorm:"default:0" json:"view_count"`
	PublishedAt *time.Time     `json:"published_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Article) TableName() string {
	return "articles"
}
