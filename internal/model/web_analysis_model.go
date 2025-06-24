package model

type WebAnalysisModel struct {
	WebUrl         string   `json:"webUrl"`
	PageTitle      string   `json:"pageTitle"`
	Headers        []string `json:"headers"`
	LoginForm      bool     `json:"loginForm"`
	WebLinks       []string `json:"webLinks"`
	BrokenWebLinks []string `json:"brokenWebLinks"`
}
