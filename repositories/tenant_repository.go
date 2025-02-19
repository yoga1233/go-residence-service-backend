package repositories

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/gorm"
)

type TenantRepository interface {
	FindAll() ([]*model.Tenant, error)
	FindByID(id int) (*model.Tenant, error)
	CreateService(service *model.Tenant) error
	UpdateService(service *model.Tenant) error
	DeleteService(service *model.Tenant) error
}

type tenantRepository struct {
	db *gorm.DB
}

// CreateService implements TenantRepository.
func (s *tenantRepository) CreateService(service *model.Tenant) error {
	return s.db.Create(&service).Error
}

// DeleteService implements TenantRepository.
func (s *tenantRepository) DeleteService(service *model.Tenant) error {
	result := s.db.Delete(&service)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// FindAll implements TenantRepository.
func (s *tenantRepository) FindAll() ([]*model.Tenant, error) {
	var services []*model.Tenant
	result := s.db.Find(&services).Order("id ASC")
	if result.Error != nil {
		return nil, result.Error
	}
	return services, nil
}

// FindByID implements TenantRepository.
func (s tenantRepository) FindByID(id int) (*model.Tenant, error) {
	var service model.Tenant

	result := s.db.Where("id = ?", id).First(&service)
	if result.Error != nil {
		return nil, result.Error
	}
	return &service, nil
}

// UpdateService implements TenantRepository.
func (s *tenantRepository) UpdateService(service *model.Tenant) error {
	var r model.Tenant

	result := s.db.Where("id = ?", service.ID).First(&r)
	if result.Error != nil {
		return result.Error
	}

	result = s.db.Model(&r).Updates(service)
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
