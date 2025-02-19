package model

import "gorm.io/gorm"

type User struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	gorm.Model

	TenantOrders []TenantOrder `json:"tenant_order" gorm:"foreignKey:UserID"`
	Reports      []Report      `json:"reports" gorm:"foreignKey:UserID"`
}
