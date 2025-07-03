package main

import (
	"net/http"
	"os"

	"github.com/GayanB90/go-web-analyzer/internal/handler"
	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type appHandler func(c *gin.Context) error

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)

	router := gin.Default()
	router.Use(gin.Recovery())

	lexicalUrlValidationService := &service.LexicalUrlValidationService{}
	httpUrlValidationService := &service.HttpUrlValidationService{}
	analysisService := &service.DefaultWebPageAnalysisService{}
	analysisService.UrlValidationServices = []service.UrlValidationService{lexicalUrlValidationService, httpUrlValidationService}

	router.Static("/static", "./static")

	router.POST("/analyze", withErrorHandler(handler.GetAnalyzePageHandler(analysisService)))
	err := router.Run(":8080")
	if err != nil {
		logrus.Error("Server exited with error: ", err)
	}
}

func withErrorHandler(h appHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			logrus.WithFields(logrus.Fields{
				"method": c.Request.Method,
				"path":   c.Request.URL.Path,
				"error":  err.Error(),
			}).Error("Error occurred in handler")

			if httpErr, ok := err.(*model.HttpError); ok {
				c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}
	}
}
