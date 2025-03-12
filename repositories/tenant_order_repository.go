package repositories

import (
	model "github.com/yoga1233/go-residence-service-backend/models"
	"gorm.io/gorm"
)

type TenantOrderRepository interface {
	FindAll() ([]*model.TenantOrder, error)
	FindById(id int) (*model.TenantOrder, error)
	FindByUserID(userId int) ([]*model.TenantOrder, error)
	CreateTenantOrder(tenantOrder *model.TenantOrder) error
	UpdateTenantOrder(tenantOrder *model.TenantOrder) error
	DeleteTenantOrder(id int) error
}

type tenantOrderRepository struct {
	db *gorm.DB
}

// FindByUserID implements TenantOrderRepository.
func (t *tenantOrderRepository) FindByUserID(userId int) ([]*model.TenantOrder, error) {
	var r []*model.TenantOrder
	result := t.db.Preload("Tenant").Where("user_id = ?", userId).Find(&r)
	if result.Error != nil {
		return nil, result.Error
	}
	return r, nil
}

// CreateTenantOrder implements TenantOrderRepository.
func (t *tenantOrderRepository) CreateTenantOrder(tenantOrder *model.TenantOrder) error {
	return t.db.Create(tenantOrder).Error

}

// DeleteTenantOrder implements TenantOrderRepository.
func (t *tenantOrderRepository) DeleteTenantOrder(id int) error {
	return t.db.Where("id = ?", id).Delete(&model.TenantOrder{}).Error
}

// FindAll implements TenantOrderRepository.
func (t *tenantOrderRepository) FindAll() ([]*model.TenantOrder, error) {
	var tenantOrder []*model.TenantOrder
	err := t.db.Find(&tenantOrder)
	if err != nil {
		return nil, err.Error
	}
	return tenantOrder, nil
}

// FindById implements TenantOrderRepository.
func (t *tenantOrderRepository) FindById(id int) (*model.TenantOrder, error) {
	var tenantOrder model.TenantOrder
	err := t.db.Where("id = ?", id).First(&tenantOrder)
	if err != nil {
		return nil, err.Error
	}
	return &tenantOrder, nil
}

// UpdateTenantOrder implements TenantOrderRepository.
func (t *tenantOrderRepository) UpdateTenantOrder(tenantOrder *model.TenantOrder) error {
	var r model.TenantOrder
	err := t.db.Model(&r).Where("id = ?", tenantOrder.ID).Updates(tenantOrder)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewTenantOrderRepository(db *gorm.DB) TenantOrderRepository {
	return &tenantOrderRepository{
		db: db,
	}
}
