package dto

type WebAnalysisRequest struct {
	RequestId string `json:"requestId"`
	WebUrl    string `json:"webUrl"`
}

type WebAnalysisResponse struct {
	WebUrl           string         `json:"webUrl"`
	PageTitle        string         `json:"pageTitle"`
	HeadersCount     map[string]int `json:"headersCount"`
	RequestId        string         `json:"requestId"`
	Hyperlinks       []string       `json:"hyperlinks"`
	ValidationErrors []string       `json:"validationErrors"`
}
