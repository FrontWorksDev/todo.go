package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"not null"`
	Email    string `form:"email" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}
