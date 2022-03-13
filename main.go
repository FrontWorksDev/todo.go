package main

import (
	"app/controller"
	"app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// sessions setting
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24, Path: "/"})
	engine.Use(sessions.Sessions("mysession", store))

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
	engine.POST("/login", controller.UserLogin)
	engine.POST("/logout", controller.UserLogout)
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
