package service

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
)

type NewsService interface {
	FindAll() ([]*model.News, error)
	FindByLimit(limit int) ([]*model.News, error)
	FindById(id int) (*model.News, error)
	CreateNews(news *model.News) error
	UpdateNews(news *model.News) error
	DeleteNews(id int) error
}

type newsService struct {
	newsRepository repositories.NewsRepository
}

// FindByLimit implements NewsService.
func (n *newsService) FindByLimit(limit int) ([]*model.News, error) {
	result, err := n.newsRepository.FindByLimit(limit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateNews implements NewsService.
func (n *newsService) CreateNews(news *model.News) error {
	err := n.newsRepository.CreateNews(news)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNews implements NewsService.
func (n *newsService) DeleteNews(id int) error {
	err := n.newsRepository.DeleteNews(id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements NewsService.
func (n *newsService) FindAll() ([]*model.News, error) {
	result, err := n.newsRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindById implements NewsService.
func (n *newsService) FindById(id int) (*model.News, error) {
	result, err := n.newsRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateNews implements NewsService.
func (n *newsService) UpdateNews(news *model.News) error {
	err := n.newsRepository.UpdateNews(news)
	if err != nil {
		return err
	}
	return nil
}

func NewNewsService(newsRepo repositories.NewsRepository) NewsService {
	return &newsService{
		newsRepository: newsRepo,
	}
}
