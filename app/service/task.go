package service

import (
	"app/model"
	"errors"
	"fmt"
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

func (TaskService) GetTaskList() []model.Task {
	DsName := "root:root@(db:3306)/todo_app?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	tests := make([]model.Task, 0)
	fmt.Println(tests)
	result := DbEngine.Find(&tests)
	if result.Error != nil {
		panic(result.Error)
	}

	defer func() {
		if db, err := DbEngine.DB(); err == nil {
			db.Close()
		}
	}()

	return tests
}

func (TaskService) UpdateTask(newTask *model.Task) error {
	DsName := "root:root@(db:3306)/todo_app?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	result := DbEngine.Model(&newTask).Where("ID", newTask.ID).Updates(model.Task{Title: newTask.Title, Status: newTask.Status})
	if result.Error != nil {
		panic(result.Error)
	}

	defer func() {
		if db, err := DbEngine.DB(); err == nil {
			db.Close()
		}
	}()

	return nil
}

func (TaskService) DeleteBook(id int) error {
	DsName := "root:root@(db:3306)/todo_app?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	task := new(model.Task)
	result := DbEngine.Where("ID", id).Delete(task)
	if result.Error != nil {
		panic(result.Error)
	}

	defer func() {
		if db, err := DbEngine.DB(); err == nil {
			db.Close()
		}
	}()

	return nil
}
