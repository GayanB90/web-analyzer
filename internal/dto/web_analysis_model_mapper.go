package dto

import (
	"github.com/GayanB90/go-web-analyzer/internal/model"
)

func ToWebAnalysisRequestModel(requestDto WebAnalysisRequest) model.WebAnalysisRequestModel {
	return model.WebAnalysisRequestModel{
		RequestId: requestDto.RequestId,
		WebUrl:    requestDto.WebUrl,
	}
}

func ToWebAnalysisResponseDto(resultModel model.WebAnalysisResultModel) WebAnalysisResponse {
	return WebAnalysisResponse{
		RequestId:        resultModel.RequestId,
		WebUrl:           resultModel.WebUrl,
		HtmlVersion:      resultModel.HtmlVersion,
		PageTitle:        resultModel.PageTitle,
		HeadersCount:     resultModel.HeadersCount,
		Hyperlinks:       resultModel.WebLinks,
		IsLoginPage:      resultModel.LoginForm,
		BrokenLinks:      resultModel.BrokenWebLinks,
		ValidationErrors: resultModel.ValidationErrors,
	}
}
