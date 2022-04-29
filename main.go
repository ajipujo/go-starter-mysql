package main

import (
	"gostart/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Router(router)
	router.Run(":3000")
}
