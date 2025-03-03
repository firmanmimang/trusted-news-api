package services

import (
	"errors"

	"github.com/firmanmimang/go-api-trusted-news/models"
	"gorm.io/gorm"
)

type NewsService struct {
	DB *gorm.DB
}

func NewNewsService(db *gorm.DB) *NewsService {
	return &NewsService{
		DB: db,
	}
}

func (s *NewsService) GetNews(limit, offset int, categorySlug *string) ([]models.News, error) {
	var news []models.News
	query := s.DB.Session(&gorm.Session{NewDB: true}).Preload("Category").Order("created_at DESC")

	if categorySlug != nil {
		var category models.Category
		result := s.DB.Session(&gorm.Session{NewDB: true}).Where("slug = ?", *categorySlug).First(&category)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}

		query = query.Where("category_id = ?", category.ID)
	}

	err := query.Limit(limit).Offset(offset).Find(&news).Error

	return news, err
}
