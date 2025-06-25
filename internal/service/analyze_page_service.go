package service

import (
	"log"
	"net/http"

	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/pkg/utils"
	"golang.org/x/net/html"
)

func AnalyzeWebPage(request model.WebAnalysisRequest) model.WebAnalysisResponse {
	var urlString = request.WebUrl

	err := utils.ValidateURL(urlString)
	if err != nil {
		return model.WebAnalysisResponse{
			RequestId:        request.RequestId,
			ValidationErrors: []string{err.Error()},
		}
	}
	resp, err := http.Get(urlString)
	if err != nil {
		log.Fatalf("An error occurred while fetching the URL: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("An error occurred while parsing the HTML: %v\n", err)
	}
	htmlTitleText := utils.ExtractHtmlTitleText(doc)
	return model.WebAnalysisResponse{
		WebUrl:    "",
		PageTitle: htmlTitleText,
		RequestId: "",
	}
}
