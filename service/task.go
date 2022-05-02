package service

import (
	"app/model"
	"errors"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TaskService struct{}

func (TaskService) SetTask(task *model.Task) (*model.Task, error) {
	loadEnv()
	DsName := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&charset=utf8mb4&loc=Local"
	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	result := DbEngine.Create(&task)
	if result.Error != nil {
		return task, result.Error
	}

	defer func() {
		if db, err := DbEngine.DB(); err == nil {
			db.Close()
		}
	}()

	return task, result.Error
}

func (TaskService) GetTaskList(userId string) []model.Task {
	loadEnv()
	DsName := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&charset=utf8mb4&loc=Local"

	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	tests := make([]model.Task, 0)
	result := DbEngine.Where(map[string]interface{}{"user_id": userId, "completed": false}).Order("updated_at asc").Find(&tests)
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

func (TaskService) GetCompletedList(userId string) []model.Task {
	loadEnv()
	DsName := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&charset=utf8mb4&loc=Local"

	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	tests := make([]model.Task, 0)
	result := DbEngine.Where(&model.Task{UserID: userId, Completed: true}).Order("updated_at asc").Find(&tests)
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

func (TaskService) UpdateTask(newTask *model.Task, id int) error {
	loadEnv()
	DsName := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&charset=utf8mb4&loc=Local"

	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	result := DbEngine.Model(&newTask).Where("ID", id).Updates(map[string]interface{}{"Title": newTask.Title, "Completed": newTask.Completed})
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
	loadEnv()
	DsName := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&charset=utf8mb4&loc=Local"

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
