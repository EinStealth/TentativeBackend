package controller

import (
	"github.com/gin-gonic/gin"
)

// @Summary     test api
// @Description ping pong
// @Accept      json
// @Success     200 {string} string "pong"
// @Router      /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}