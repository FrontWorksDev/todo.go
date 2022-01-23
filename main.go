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
	engine.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"https://localhost:3000",
			"https://todo.frontworks.dev",
		},
	}))

	// middleware
	engine.Use(middleware.RecordUpAndTime)

	// CRUD task
	taskEngine := engine.Group("/task")
	{
		v1 := taskEngine.Group("/v1")
		{
			v1.POST("/add", controller.TaskAdd)
			v1.GET("/list", controller.TaskList)
			v1.PUT("/update/:id", controller.TaskUpdate)
			v1.DELETE("/delete/:id", controller.TaskDelete)
		}
	}

	engine.Run()
}
