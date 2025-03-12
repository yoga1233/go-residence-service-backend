package model

import "gorm.io/gorm"

type TenantCategory struct {
	gorm.Model
	Name        string   `json:"name" gorm:"type:varchar(100);not null"`
	Description string   `json:"description" gorm:"type:text"`
	Tenants     []Tenant `json:"tenants" gorm:"foreignKey:CategoryID;constraint:onDelete:CASCADE"`
}
