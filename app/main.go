package main

import (
	"app/controller"
	"app/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// cors setting
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
	}
	engine.Use(cors.New(config))

	// middleware
	engine.Use(middleware.RecordUpAndTime)

	// CRUD task
	taskEngine := engine.Group("/task")
	{
		v1 := taskEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
		}
	}

	engine.Run()
}
