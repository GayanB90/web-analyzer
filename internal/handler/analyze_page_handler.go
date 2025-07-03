package handler

import (
	"net/http"

	"github.com/GayanB90/go-web-analyzer/internal/dto"
	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetAnalyzePageHandler(pageAnalysisService service.WebPageAnalysisService) func(c *gin.Context) error {
	return func(c *gin.Context) error {
		var webAnalysisRequest dto.WebAnalysisRequest

		if err := c.BindJSON(&webAnalysisRequest); err != nil {
			return &model.HttpError{StatusCode: http.StatusInternalServerError, Message: "An error occurred while parsing the request data"}
		}
		logrus.WithFields(logrus.Fields{
			"Method": c.Request.Method,
			"URL":    c.Request.URL,
		}).Infof("Binding web analysisRequest successful: %v", webAnalysisRequest)

		webAnalysisRequestModel := dto.ToWebAnalysisRequestModel(webAnalysisRequest)
		logrus.WithFields(logrus.Fields{
			"Method": c.Request.Method,
			"URL":    c.Request.URL,
		}).Infof("Initiating web page analysis for the request model %s", webAnalysisRequestModel)
		webAnalysisResultModel, err := pageAnalysisService.AnalyzeWebPage(webAnalysisRequestModel)
		if err != nil {
			return err
		}
		logrus.WithFields(logrus.Fields{
			"Method": c.Request.Method,
			"URL":    c.Request.URL,
		}).Infof("Successfully analyzed the page for the request model %s, result: %s", webAnalysisRequestModel, webAnalysisResultModel)
		webAnalysisResponse := dto.ToWebAnalysisResponseDto(webAnalysisResultModel)
		logrus.WithFields(logrus.Fields{
			"Method": c.Request.Method,
			"URL":    c.Request.URL,
		}).Infof("Web Analysis Response: %v", webAnalysisResponse)
		c.IndentedJSON(http.StatusOK, webAnalysisResponse)
		return nil
	}
}
