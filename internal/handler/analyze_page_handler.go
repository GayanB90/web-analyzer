package handler

import (
	"log"
	"net/http"

	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/internal/service"
	"github.com/gin-gonic/gin"
)

func AnalyzePageHandler(c *gin.Context) {
	var webAnalysisRequest model.WebAnalysisRequest

	if err := c.BindJSON(&webAnalysisRequest); err != nil {
		return
	}
	log.Printf("Binding web analysisRequest successful: %v", webAnalysisRequest)

	webAnalysisResponse := service.AnalyzeWebPage(webAnalysisRequest)
	log.Printf("Web Analysis Response: %v", webAnalysisResponse)
	c.IndentedJSON(http.StatusOK, webAnalysisResponse)
}
