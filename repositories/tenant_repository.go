package repositories

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/gorm"
)

type TenantRepository interface {
	FindAll() ([]*model.Tenant, error)
	FindByID(id int) (*model.Tenant, error)
	FindTenantByQuery(query string) ([]*model.Tenant, error)
	CreateTenant(tenant *model.Tenant) error
	UpdateTenant(tenant *model.Tenant) error
	DeleteTenant(id int) error
}

type tenantRepository struct {
	db *gorm.DB
}

// FindTenantByQuery implements TenantRepository.
func (s *tenantRepository) FindTenantByQuery(query string) ([]*model.Tenant, error) {
	var t []*model.Tenant
	err := s.db.Where("name LIKE ? OR description LIKE ?", query, query).Find(&t)
	if err != nil {
		return nil, err.Error
	}
	return t, nil
}

// CreateService implements TenantRepository.
func (s *tenantRepository) CreateTenant(tenant *model.Tenant) error {
	return s.db.Create(&tenant).Error
}

// DeleteService implements TenantRepository.
func (s *tenantRepository) DeleteTenant(id int) error {
	result := s.db.Where("id = ?", id).Delete(&model.Tenant{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// FindAll implements TenantRepository.
func (s *tenantRepository) FindAll() ([]*model.Tenant, error) {
	var t []*model.Tenant
	result := s.db.Preload("TenantCategory").Where("status = ?", "available").Find(&t).Order("id ASC")
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

// FindByID implements TenantRepository.
func (s tenantRepository) FindByID(id int) (*model.Tenant, error) {
	var t model.Tenant

	result := s.db.Where("id = ?", id).First(&t)
	if result.Error != nil {
		return nil, result.Error
	}
	return &t, nil
}

// UpdateService implements TenantRepository.
func (s *tenantRepository) UpdateTenant(tenant *model.Tenant) error {
	var r model.Tenant

	result := s.db.Model(&r).Where("id = ?", tenant.ID).Updates(tenant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{
		db: db,
	}
}
