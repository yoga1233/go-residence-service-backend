package repositories

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]*model.TenantCategory, error)
	FindByName(name string) ([]*model.TenantCategory, error)
	CreateCategory(category *model.TenantCategory) error
	UpdateCategory(category *model.TenantCategory) error
	DeleteCategory(id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

// createCategory implements CategoryRepository.
func (c *categoryRepository) CreateCategory(category *model.TenantCategory) error {
	err := c.db.Create(category).Error
	return err
}

// deleteCategory implements CategoryRepository.
func (c *categoryRepository) DeleteCategory(id int) error {
	err := c.db.Delete(&model.TenantCategory{}, id).Error

	return err
}

// findAll implements CategoryRepository.
func (c *categoryRepository) FindAll() ([]*model.TenantCategory, error) {
	var categories []*model.TenantCategory
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil

}

// findByName implements CategoryRepository.
func (c *categoryRepository) FindByName(name string) ([]*model.TenantCategory, error) {
	var categories []*model.TenantCategory
	err := c.db.Where("name = ?", name).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// updateCategory implements CategoryRepository.
func (c *categoryRepository) UpdateCategory(category *model.TenantCategory) error {
	err := c.db.Save(category).Error
	return err
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}
