package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishal2457/redmine-server/models"
)

// CreateUser creates a new user.
func CreateTask(c *gin.Context) {
	var task models.Task

	c.BindJSON(&task)

	err := models.CreateTask(&task)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusCreated, task)
	}
}

// GetUsers gets all existing users.
func GetTasks(c *gin.Context) {
	var tasks []models.Task

	err := models.GetTasks(&tasks)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}
