package model

type WebAnalysisRequestModel struct {
	WebUrl string `json:"webUrl"`
}
type WebAnalysisResultModel struct {
	WebUrl           string   `json:"webUrl"`
	PageTitle        string   `json:"pageTitle"`
	Headers          []string `json:"headers"`
	LoginForm        bool     `json:"loginForm"`
	WebLinks         []string `json:"webLinks"`
	BrokenWebLinks   []string `json:"brokenWebLinks"`
	ValidationErrors []string `json:"validationErrors"`
}
