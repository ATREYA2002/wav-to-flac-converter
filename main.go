package main

import (
	"wavtoflacconverter/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080") // Start server on port 8080
}
