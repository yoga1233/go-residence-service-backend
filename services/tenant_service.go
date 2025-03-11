package service

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
)

type TenantService interface {
	FindAll() ([]*model.Tenant, error)
	FindByID(id int) (*model.Tenant, error)
	FindTenantByQuery(query string) ([]*model.Tenant, error)
	CreateTenant(tenant *model.Tenant) error
	UpdateTenant(tenant *model.Tenant) error
	DeleteTenant(id int) error
}

type tenantService struct {
	tenantRepository repositories.TenantRepository
}

// FindTenantByQuery implements TenantService.
func (s *tenantService) FindTenantByQuery(query string) ([]*model.Tenant, error) {
	result, err := s.tenantRepository.FindTenantByQuery(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTenant implements TenantService.
func (s *tenantService) CreateTenant(tenant *model.Tenant) error {
	result := s.tenantRepository.CreateTenant(tenant)
	if result != nil {
		return result
	}
	return nil
}

// DeleteTenant implements TenantService.
func (s *tenantService) DeleteTenant(id int) error {
	result := s.tenantRepository.DeleteTenant(id)
	if result != nil {
		return result
	}
	return nil
}

// FindAll implements TenantService.
func (s *tenantService) FindAll() ([]*model.Tenant, error) {
	result, err := s.tenantRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindByID implements TenantService.
func (s *tenantService) FindByID(id int) (*model.Tenant, error) {
	result, err := s.tenantRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTenant implements TenantService.
func (s *tenantService) UpdateTenant(tenant *model.Tenant) error {
	err := s.tenantRepository.UpdateTenant(tenant)
	if err != nil {
		return err
	}
	return nil

}

func NewTenantService(tenantRepo repositories.TenantRepository) TenantService {
	return &tenantService{
		tenantRepository: tenantRepo,
	}
}
