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

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		return
	}

	userService := service.UserService{}
	loginUser := userService.GetUser(&user)[0].ID
	if loginUser == 0 {
		err = userService.SetUser(&user)
		if err != nil {
			c.String(http.StatusInternalServerError, "Server Error")

			return
		}
		newUser := userService.GetUser(&user)[0].ID
		session := sessions.Default(c)
		session.Set("loginUser", newUser)
		session.Save()
		c.Abort()
	} else {
		session := sessions.Default(c)
		session.Set("loginUser", loginUser)
		session.Save()
		c.Next()
	}
	taskService := service.TaskService{}
	session := sessions.Default(c)
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
