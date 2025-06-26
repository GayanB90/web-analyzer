package service

import "github.com/GayanB90/go-web-analyzer/internal/model"

type WebPageAnalysisService interface {
	AnalyzeWebPage(webPage model.WebAnalysisRequestModel) model.WebAnalysisResultModel
}
