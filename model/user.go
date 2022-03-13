package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"varchar(255)" json:"name" form:"name"`
	Email string `gorm:"varchar(255);unique" json:"email" form:"email"`
	Image string `gorm:"varchar(255)" json:"image" form:"image"`
}
