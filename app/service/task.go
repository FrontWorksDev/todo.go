package service

import (
	"app/model"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type TaskService struct{}

func (TaskService) SetTask(task *model.Task) error {
	DsName := "root:root@(db:3306)/todo_app?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	result := DbEngine.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	defer func() {
		if db, err := DbEngine.DB(); err == nil {
			db.Close()
		}
	}()

	return nil
}
