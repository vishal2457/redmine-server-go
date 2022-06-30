package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishal2457/redmine-server/models"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project

	err := models.GetProjects(&projects)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, projects)
	}
}
