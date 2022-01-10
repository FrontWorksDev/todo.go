package service

import (
	"app/model"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DbEngine *gorm.DB

func init() {
	DsName := "root:root@(db:3306)/todo_app?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	DbEngine.AutoMigrate(new(model.Task))
	defer func() {
		if db, err := DbEngine.DB(); err == nil {
			db.Close()
		}
	}()
}
