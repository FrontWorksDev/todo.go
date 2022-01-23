package controller

import (
	"app/model"
	"app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TaskAdd(c *gin.Context) {
	task := model.Task{}
	err := c.Bind(&task)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		return
	}

	taskService := service.TaskService{}
	err = taskService.SetTask(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func TaskList(c *gin.Context) {
	taskService := service.TaskService{}
	TaskLists := taskService.GetTaskList()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"items":   TaskLists,
	})
}

func TaskUpdate(c *gin.Context) {
	id := c.Param("id")
	intId, errId := strconv.ParseInt(id, 10, 0)

	task := model.Task{}
	err := c.Bind(&task)
	if err != nil || errId != nil {
		c.String(http.StatusBadRequest, "Bad Request")

		return
	}

	taskService := service.TaskService{}
	err = taskService.UpdateTask(&task, int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func TaskDelete(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		return
	}

	taskService := service.TaskService{}
	err = taskService.DeleteBook(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
