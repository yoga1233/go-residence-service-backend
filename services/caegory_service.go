package service

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
)

type CategoryService interface {
	searchCategory(name string) ([]*model.TenantCategory, error)
	FindAll() ([]*model.TenantCategory, error)
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

// FindAll implements CategoryService.
func (c *categoryService) FindAll() ([]*model.TenantCategory, error) {
	result, err := c.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// searchCategory implements CategoryService.
func (c *categoryService) searchCategory(name string) ([]*model.TenantCategory, error) {
	result, err := c.categoryRepo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}
