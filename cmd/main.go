package main

import (
	"github.com/AbduRohman-developer/cv-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Pong(router)

	router.Run(":80")
}
