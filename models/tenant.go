package model

import "gorm.io/gorm"

type Tenant struct {
	// Id          int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
	Status      string `json:"status" gorm:"not null;type:ENUM('available', 'unavailable');default:'available'"`
	gorm.Model

	TenantOrders []TenantOrder `json:"tenant_orders" gorm:"foreignKey:TenantID"`
}
