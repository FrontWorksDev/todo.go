package controller

import (
	"app/model"
	"app/service"
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
		session.Set("loginUser", loginUser)
	} else {
		newUser := userService.GetUser(&user)[0].ID
		session.Set("loginUser", newUser)
	}

	session.Save()
	taskService := service.TaskService{}
	UserId := session.Get("loginUser")
	taskData := taskService.GetTaskList(UserId)

	c.JSON(http.StatusCreated, gin.H{
		"status": session.Get("loginUser"),
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
