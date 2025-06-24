package model

type WebAnalysisRequest struct {
	RequestId string `json:"requestId"`
	WebUrl    string `json:"webUrl"`
}

type WebAnalysisResponse struct {
	WebUrl    string `json:"webUrl"`
	PageTitle string `json:"pageTitle"`
	RequestId string `json:"requestId"`
}
