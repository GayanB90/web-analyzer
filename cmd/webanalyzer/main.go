package main

import (
	"log"

	"github.com/GayanB90/go-web-analyzer/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/analyze", handler.AnalyzePageHandler)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server exited with error: ", err)
	}
}
