package service

import (
	"log"
	"net/http"

	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/pkg/utils"
	"golang.org/x/net/html"
)

type DefaultWebPageAnalysisService struct {
	UrlValidationServices []UrlValidationService
}

func (s *DefaultWebPageAnalysisService) AnalyzeWebPage(request model.WebAnalysisRequestModel) model.WebAnalysisResultModel {
	var urlString = request.WebUrl

	err := utils.ValidateURL(urlString)
	if err != nil {
		return model.WebAnalysisResultModel{
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
	log.Printf("htmlTitleText: %v", htmlTitleText)
	headingCountMap := make(map[string]int)
	utils.ExtractHeadingCount(doc, headingCountMap)
	log.Printf("headingCountMap: %v", headingCountMap)
	hyperlinks := utils.ExtractHyperlinks(doc)
	log.Printf("hyperlinks: %v", hyperlinks)
	brokenLinks := s.findBrokenHyperlinks(hyperlinks)
	log.Printf("brokenLinks: %v", brokenLinks)
	loginFormAvailable := utils.IsLoginFormAvailable(doc)
	log.Printf("loginFormAvailable: %v", loginFormAvailable)
	return model.WebAnalysisResultModel{
		WebUrl:         urlString,
		PageTitle:      htmlTitleText,
		HeadersCount:   headingCountMap,
		WebLinks:       hyperlinks,
		BrokenWebLinks: brokenLinks,
		LoginForm:      loginFormAvailable,
	}
}

func (s *DefaultWebPageAnalysisService) findBrokenHyperlinks(links []string) []string {
	var brokenLinks []string

	for _, link := range links {
		log.Printf("Validating hyperlink: %v", link)
		for _, urlValidationService := range s.UrlValidationServices {
			if urlValidationService.ValidateUrl(link) != nil {
				brokenLinks = append(brokenLinks, link)
			}
		}
	}
	return brokenLinks
}
