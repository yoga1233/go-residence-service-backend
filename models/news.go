package model

import "gorm.io/gorm"

type News struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Title    string `json:"title" gorm:"not null"`
	Content  string `json:"content" gorm:"not null"`
	Category string `json:"category" gorm:"not null;type:ENUM('event', 'announcement', 'maintenance');default:'announcement'"`
	gorm.Model
}
