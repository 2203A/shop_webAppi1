package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string `gorm:"type:char(11);unique;not null"`
	Password string `gorm:"type:char(32);not null"`
}
