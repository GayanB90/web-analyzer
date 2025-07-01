package main

import (
	"log"
	"net/http"

	"github.com/GayanB90/go-web-analyzer/internal/handler"
	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/internal/service"
	"github.com/gin-gonic/gin"
)

type appHandler func(c *gin.Context) error

func main() {
	router := gin.Default()

	lexicalUrlValidationService := &service.LexicalUrlValidationService{}
	httpUrlValidationService := &service.HttpUrlValidationService{}
	analysisService := &service.DefaultWebPageAnalysisService{}
	analysisService.UrlValidationServices = []service.UrlValidationService{lexicalUrlValidationService, httpUrlValidationService}

	router.Static("/static", "./static")

	router.POST("/analyze", withErrorHandler(handler.GetAnalyzePageHandler(analysisService)))
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server exited with error: ", err)
	}
}

func withErrorHandler(h appHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			log.Printf("error: %v", err)
			if httpErr, ok := err.(*model.HttpError); ok {
				c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}
	}
}
