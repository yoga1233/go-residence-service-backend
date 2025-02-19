package model

import "gorm.io/gorm"

type TenantOrder struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	TenantID int    `json:"service_id" gorm:"not null"`
	UserID   int    `json:"user_id" gorm:"not null"`
	Status   string `json:"status" gorm:"not null;type:ENUM('pending', 'accepted','completed', 'rejected');default:'pending'"`
	gorm.Model

	Tenant Tenant `json:"tenant" gorm:"foreignKey:TenantID;references:ID"`
	User   User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
