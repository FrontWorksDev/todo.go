package main

import (
	"app/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// middleware
	engine.Use(middleware.RecordUpAndTime)

	engine.Run()
}
