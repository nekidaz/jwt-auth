package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WelcomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Error : it works",
	})
}
