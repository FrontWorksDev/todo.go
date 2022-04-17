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
		"https://todo.frontworks.dev",
	}

	config.AllowCredentials = true
	engine.Use(cors.New(config))

	// middleware
	engine.Use(middleware.RecordUpAndTime)

	// CRUD task
	taskEngine := engine.Group("/task")
	{
		v1 := taskEngine.Group("/v1")
		{
			v1.POST("/add", controller.TaskAdd)
			v1.POST("/list/:userId", controller.TaskList)
			v1.PUT("/update/:id", controller.TaskUpdate)
			v1.DELETE("/delete/:id", controller.TaskDelete)
		}
	}
	engine.Run()
}
