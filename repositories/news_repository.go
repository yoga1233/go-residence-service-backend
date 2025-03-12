package repositories

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/gorm"
)

type NewsRepository interface {
	FindAll() ([]*model.News, error)
	FindById(id int) (*model.News, error)
	FindByLimit(limit int) ([]*model.News, error)
	CreateNews(news *model.News) error
	UpdateNews(news *model.News) error
	DeleteNews(id int) error
}

type newsRepository struct {
	db *gorm.DB
}

// FindByLimit implements NewsRepository.
func (n *newsRepository) FindByLimit(limit int) ([]*model.News, error) {
	var news []*model.News
	result := n.db.Limit(limit).Order("updated_at DESC").Find(&news)

	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

// CreateNews implements NewsRepository.
func (n *newsRepository) CreateNews(news *model.News) error {
	return n.db.Create(news).Error
}

// DeleteNews implements NewsRepository.
func (n *newsRepository) DeleteNews(id int) error {
	news := new(model.News)
	result := n.db.Where("id = ?", id).Delete(&news)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindAll implements NewsRepository.
func (n *newsRepository) FindAll() ([]*model.News, error) {
	var news []*model.News
	result := n.db.Find(&news).Order("id ASC")

	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

// FindById implements NewsRepository.
func (n *newsRepository) FindById(id int) (*model.News, error) {
	news := new(model.News)
	result := n.db.Where("id = ?", id).First(&news)

	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

// UpdateNews implements NewsRepository.
func (n *newsRepository) UpdateNews(news *model.News) error {
	return n.db.Save(news).Error
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{
		db: db,
	}
}
