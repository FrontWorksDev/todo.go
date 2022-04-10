package controller

import (
	"app/model"
	"app/service"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var result *model.User

type UserInfo struct {
	username        string
	email           string
	passwordEncrypt byte
}

func UserCreate(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		return
	}

	passwordEncrypt, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(passwordEncrypt)
	userService := service.UserService{}
	userService.SetUser(&user)
}

func UserLogin(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	session := sessions.Default(c)

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		return
	}

	userService := service.UserService{}
	userInfo := userService.GetUser(&user)[0]
	dbPassword := userInfo.Password

	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(user.Password)); err != nil {
		log.Println("Login Failed")
		c.Abort()
	} else {
		log.Println("Login Success")
		session.Set("username", userInfo.Username)
	}
	session.Save()
	c.JSON(http.StatusCreated, gin.H{
		"session": session.Get("username"),
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
