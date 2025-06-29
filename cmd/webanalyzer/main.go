package main

import (
	"log"

	"github.com/GayanB90/go-web-analyzer/internal/handler"
	"github.com/GayanB90/go-web-analyzer/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	analysisService := &service.DefaultWebPageAnalysisService{}

	router.Static("/static", "../../static")

	router.POST("/analyze", handler.GetAnalyzePageHandler(analysisService))
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server exited with error: ", err)
	}
}
