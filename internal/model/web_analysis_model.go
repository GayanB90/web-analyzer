package model

type WebAnalysisRequestModel struct {
	RequestId string `json:"requestId"`
	WebUrl    string `json:"webUrl"`
}
type WebAnalysisResultModel struct {
	RequestId        string         `json:"requestId"`
	WebUrl           string         `json:"webUrl"`
	HtmlVersion      string         `json:"htmlVersion"`
	PageTitle        string         `json:"pageTitle"`
	HeadersCount     map[string]int `json:"headersCountCount"`
	LoginForm        bool           `json:"loginForm"`
	WebLinks         []string       `json:"webLinks"`
	BrokenWebLinks   []string       `json:"brokenWebLinks"`
	ValidationErrors []string       `json:"validationErrors"`
}
