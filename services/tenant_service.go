package service

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
)

type TenantService interface {
	FindAll() ([]*model.Tenant, error)
	FindByID(id int) (*model.Tenant, error)
	CreateTenant(tenant *model.Tenant) error
	UpdateTenant(tenant *model.Tenant) error
	DeleteTenant(tenant *model.Tenant) error
}

type tenantService struct {
	tenantRepository repositories.TenantRepository
}

// CreateTenant implements TenantService.
func (*tenantService) CreateTenant(tenant *model.Tenant) error {
	panic("unimplemented")
}

// DeleteTenant implements TenantService.
func (*tenantService) DeleteTenant(tenant *model.Tenant) error {
	panic("unimplemented")
}

// FindAll implements TenantService.
func (*tenantService) FindAll() ([]*model.Tenant, error) {
	panic("unimplemented")
}

// FindByID implements TenantService.
func (*tenantService) FindByID(id int) (*model.Tenant, error) {
	panic("unimplemented")
}

// UpdateTenant implements TenantService.
func (*tenantService) UpdateTenant(tenant *model.Tenant) error {
	panic("unimplemented")
}

func NewTenantRepository(tenantRepo repositories.TenantRepository) TenantService {
	return &tenantService{
		tenantRepository: tenantRepo,
	}
}
