package service

import (
	"app/model"
	"errors"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserService struct{}

func (UserService) SetUser(user *model.User) error {
	loadEnv()
	DsName := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&charset=utf8mb4&loc=Local"
	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	result := DbEngine.Create(&user)

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

func (UserService) GetUser(user *model.User) []model.User {
	loadEnv()
	DsName := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&charset=utf8mb4&loc=Local"

	err := errors.New("")
	DbEngine, err = gorm.Open(mysql.Open(DsName), &gorm.Config{})
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	tests := make([]model.User, 0)
	result := DbEngine.Where("Email", user.Email).Find(&tests)
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
