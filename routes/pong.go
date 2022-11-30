package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Engine) {
	c.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
