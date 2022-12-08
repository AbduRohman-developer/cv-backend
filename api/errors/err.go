package err

import "github.com/gin-gonic/gin"

func New(code int, msg string, c *gin.Context) {
	c.JSON(code, map[string]string{
		"error": msg,
	})
}
