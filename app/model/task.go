package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `gorm:"varchar(40)" json:"title" form:"title"`
	Status int    `gorm:"integer(1)" json:"status" form:"status"`
	Slug   string `gorm:"varchar(255)" json:"slug" form:"slug"`
}
