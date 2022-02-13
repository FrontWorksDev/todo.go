package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `gorm:"varchar(40)" json:"title" form:"title"`
	//Status    string `gorm:"varchar(40)" json:"status" form:"status"`
	Slug      string `gorm:"varchar(255);unique" json:"slug" form:"slug"`
	Completed bool   `gorm:"default:false;not null" json:"completed" form:"completed"`
}
