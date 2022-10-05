package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Questions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "questions",
	})
}
