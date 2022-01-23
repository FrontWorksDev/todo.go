package service

import (
	"app/model"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DbEngine *gorm.DB

func init() {
	loadEnv()
	DsName := os.Getenv("CLEARDB_DATABASE_URL")
	fmt.Println(DsName)
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

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Faild to load: %v", err)
	}

	return
}
