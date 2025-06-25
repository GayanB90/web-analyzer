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
		WebUrl: requestModel.WebUrl,
	}
}

func ToWebAnalysisResponseDto(resultModel model.WebAnalysisResultModel) WebAnalysisResponse {
	return WebAnalysisResponse{
		WebUrl:           resultModel.WebUrl,
		PageTitle:        resultModel.PageTitle,
		RequestId:        "",
		Hyperlinks:       resultModel.WebLinks,
		ValidationErrors: resultModel.ValidationErrors,
	}
}
