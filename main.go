package main

import (
	"app/controller"
	"app/middleware"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// sessions setting
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24, Path: "/", SameSite: http.SameSiteNoneMode, Secure: true})
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

func sessionCheck(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		UserId := session.Get("content")
		fmt.Println(UserId)

		if UserId == nil {
			fmt.Println("not login")
			engine.POST("/login", controller.UserLogin)
			c.Abort()
		} else {
			c.Set("UserId", UserId)
			c.Next()
		}

		fmt.Println("ended login")
	}
}
