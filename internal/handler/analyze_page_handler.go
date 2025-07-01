package handler

import (
	"log"
	"net/http"

	"github.com/GayanB90/go-web-analyzer/internal/dto"
	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/internal/service"
	"github.com/gin-gonic/gin"
)

func GetAnalyzePageHandler(pageAnalysisService service.WebPageAnalysisService) func(c *gin.Context) error {
	return func(c *gin.Context) error {
		var webAnalysisRequest dto.WebAnalysisRequest

		if err := c.BindJSON(&webAnalysisRequest); err != nil {
			return &model.HttpError{StatusCode: http.StatusInternalServerError, Message: "An error occurred while parsing the request data"}
		}
		log.Printf("Binding web analysisRequest successful: %v", webAnalysisRequest)

		webAnalysisRequestModel := dto.ToWebAnalysisRequestModel(webAnalysisRequest)
		log.Printf("Initiating web page analysis for the request model %v", webAnalysisRequestModel)
		webAnalysisResultModel, err := pageAnalysisService.AnalyzeWebPage(webAnalysisRequestModel)
		if err != nil {
			return err
		}
		log.Printf("Successfully analyzed the page for the request model %v, result: %v", webAnalysisRequestModel, webAnalysisResultModel)
		webAnalysisResponse := dto.ToWebAnalysisResponseDto(webAnalysisResultModel)
		log.Printf("Web Analysis Response: %v", webAnalysisResponse)
		c.IndentedJSON(http.StatusOK, webAnalysisResponse)
		return nil
	}
}
