package model

import "gorm.io/gorm"

type Report struct {
	// ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"type:varchar(150);not null" validate:"required"`
	Description string `gorm:"type:text" validate:"required"`
	Status      string `gorm:"type:enum('pending', 'in_progress', 'resolved');default:'pending'"`
	ImageUrl    string `gorm:"type:varchar(250)"`
	gorm.Model

	User User `json:"-" gorm:"foreignKey:UserID"`
}
