package model

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
	Status      string `json:"status" gorm:"not null;type:varchar(20);default:'available'"`

	// Foreign key ke TenantCategory
	CategoryID     uint           `json:"category_id" gorm:"not null"`
	TenantCategory TenantCategory `json:"tenant_category" gorm:"foreignKey:CategoryID;references:ID"`

	// Relasi ke TenantOrder
	TenantOrders []TenantOrder `json:"-" gorm:"foreignKey:TenantID"`
}
