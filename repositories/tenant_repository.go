package repositories

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/gorm"
)

type TenantRepository interface {
	FindAll() ([]*model.Tenant, error)
	FindByID(id int) (*model.Tenant, error)
	CreateTenant(tenant *model.Tenant) error
	UpdateTenant(tenant *model.Tenant) error
	DeleteTenant(tenant *model.Tenant) error
}

type tenantRepository struct {
	db *gorm.DB
}

// CreateService implements TenantRepository.
func (s *tenantRepository) CreateTenant(tenant *model.Tenant) error {
	return s.db.Create(&tenant).Error
}

// DeleteService implements TenantRepository.
func (s *tenantRepository) DeleteTenant(tenant *model.Tenant) error {
	result := s.db.Delete(&tenant)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// FindAll implements TenantRepository.
func (s *tenantRepository) FindAll() ([]*model.Tenant, error) {
	var t []*model.Tenant
	result := s.db.Find(&t).Order("id ASC")
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

	result := s.db.Where("id = ?", tenant.ID).First(&r)
	if result.Error != nil {
		return result.Error
	}

	result = s.db.Model(&r).Updates(tenant)
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
