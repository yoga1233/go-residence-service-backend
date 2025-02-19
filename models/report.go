package model

import "gorm.io/gorm"

type Report struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"type:varchar(150);not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:enum('pending', 'in_progress', 'resolved');default:'pending'"`
	gorm.Model

	User User `gorm:"foreignKey:UserID"`
}
