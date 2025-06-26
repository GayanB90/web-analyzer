package handler

import (
	"log"
	"net/http"

	"github.com/GayanB90/go-web-analyzer/internal/dto"
	"github.com/GayanB90/go-web-analyzer/internal/service"
	"github.com/gin-gonic/gin"
)

func GetAnalyzePageHandler(pageAnalysisService service.WebPageAnalysisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var webAnalysisRequest dto.WebAnalysisRequest

		if err := c.BindJSON(&webAnalysisRequest); err != nil {
			return
		}
		log.Printf("Binding web analysisRequest successful: %v", webAnalysisRequest)

		webAnalysisRequestModel := dto.ToWebAnalysisRequestModel(webAnalysisRequest)
		log.Printf("Initiating web page analysis for the request model %v", webAnalysisRequestModel)
		webAnalysisResultModel := pageAnalysisService.AnalyzeWebPage(webAnalysisRequestModel)
		log.Printf("Successfully analyzed the page for the request model %v, result: %v", webAnalysisRequestModel, webAnalysisResultModel)
		webAnalysisResponse := dto.ToWebAnalysisResponseDto(webAnalysisResultModel)
		log.Printf("Web Analysis Response: %v", webAnalysisResponse)
		c.IndentedJSON(http.StatusOK, webAnalysisResponse)
	}
}
