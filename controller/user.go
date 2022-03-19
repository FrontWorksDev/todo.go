package controller

import (
	"app/model"
	"app/service"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var result *model.User

func UserLogin(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	session := sessions.Default(c)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		return
	}

	userService := service.UserService{}
	err = userService.SetUser(&user)
	if err != nil {
		loginUser := userService.GetUser(&user)[0].ID
		session.Set("UserId", loginUser)
	} else {
		newUser := userService.GetUser(&user)[0].ID
		session.Set("UserId", newUser)
	}

	session.Save()
	taskService := service.TaskService{}
	userId := session.Get("UserId")
	taskData := taskService.GetTaskList(userId)
	fmt.Println("controller", userId)
	c.JSON(http.StatusCreated, gin.H{
		"status": session.Get("UserId"),
		"items":  taskData,
	})
}

func UserLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusCreated, gin.H{
		"status": "logout success",
	})
}
