package service

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
)

type TenantOrderService interface {
	FindAll() ([]*model.TenantOrder, error)
	FindById(id int) (*model.TenantOrder, error)
	FindByUserID(userId int) (*model.TenantOrder, error)
	CreateTenantOrder(tenantOrder *model.TenantOrder) error
	UpdateTenantOrder(tenantOrder *model.TenantOrder) error
	DeleteTenantOrder(id int) error
}

type tenantOrderService struct {
	tenantOrderRepository repositories.TenantOrderRepository
}

// FindByUserID implements TenantOrderService.
func (s *tenantOrderService) FindByUserID(userId int) (*model.TenantOrder, error) {
	result, err := s.tenantOrderRepository.FindByUserID(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateTenantOrder implements TenantOrderService.
func (s *tenantOrderService) CreateTenantOrder(tenantOrder *model.TenantOrder) error {
	result := s.tenantOrderRepository.CreateTenantOrder(tenantOrder)
	if result != nil {
		return result
	}
	return nil
}

// DeleteTenantOrder implements TenantOrderService.
func (s *tenantOrderService) DeleteTenantOrder(id int) error {
	result := s.tenantOrderRepository.DeleteTenantOrder(id)
	if result != nil {
		return result
	}
	return nil
}

// FindAll implements TenantOrderService.
func (s *tenantOrderService) FindAll() ([]*model.TenantOrder, error) {
	result, err := s.tenantOrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindById implements TenantOrderService.
func (s *tenantOrderService) FindById(id int) (*model.TenantOrder, error) {
	result, err := s.tenantOrderRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTenantOrder implements TenantOrderService.
func (s *tenantOrderService) UpdateTenantOrder(tenantOrder *model.TenantOrder) error {
	err := s.tenantOrderRepository.UpdateTenantOrder(tenantOrder)
	if err != nil {
		return err
	}
	return nil
}

func NewTenantOrderService(tenantOrderRepository repositories.TenantOrderRepository) TenantOrderService {
	return &tenantOrderService{
		tenantOrderRepository: tenantOrderRepository,
	}
}
