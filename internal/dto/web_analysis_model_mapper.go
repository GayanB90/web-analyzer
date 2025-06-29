package dto

import (
	"github.com/GayanB90/go-web-analyzer/internal/model"
)

func ToWebAnalysisRequestModel(requestDto WebAnalysisRequest) model.WebAnalysisRequestModel {
	return model.WebAnalysisRequestModel{
		WebUrl: requestDto.WebUrl,
	}
}

func ToWebAnalysisRequestDto(requestModel model.WebAnalysisRequestModel) WebAnalysisRequest {
	return WebAnalysisRequest{
		RequestId: requestModel.RequestId,
		WebUrl:    requestModel.WebUrl,
	}
}

func ToWebAnalysisResponseDto(resultModel model.WebAnalysisResultModel) WebAnalysisResponse {
	return WebAnalysisResponse{
		WebUrl:           resultModel.WebUrl,
		HtmlVersion:      resultModel.HtmlVersion,
		PageTitle:        resultModel.PageTitle,
		HeadersCount:     resultModel.HeadersCount,
		RequestId:        resultModel.RequestId,
		Hyperlinks:       resultModel.WebLinks,
		IsLoginPage:      resultModel.LoginForm,
		ValidationErrors: resultModel.ValidationErrors,
	}
}
