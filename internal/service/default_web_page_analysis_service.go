package service

import (
	"bytes"
	"io"
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An error occurred while buffering the resp body: %v\n", err)
	}

	htmlVersion := utils.ExtractHtmlVersion(bytes.NewReader(data))
	log.Printf("htmlVersion: %v", htmlVersion)

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		log.Fatalf("An error occurred while parsing the HTML: %v\n", err)
	}
	htmlTitleText := utils.ExtractHtmlTitleText(doc)
	log.Printf("htmlTitleText: %v", htmlTitleText)
	headingCountMap := make(map[string]int)
	utils.ExtractHeadingCount(doc, headingCountMap)
	log.Printf("headingCountMap: %v", headingCountMap)
	hyperlinksList := make([]string, 0)
	utils.ExtractHyperlinks(doc, &hyperlinksList)
	log.Printf("hyperlinks: %v", hyperlinksList)
	brokenLinks := s.findBrokenHyperlinks(hyperlinksList)
	log.Printf("brokenLinks: %v", brokenLinks)
	loginFormAvailable := utils.IsLoginFormAvailable(doc)
	log.Printf("loginFormAvailable: %v", loginFormAvailable)
	return model.WebAnalysisResultModel{
		RequestId:      request.RequestId,
		WebUrl:         urlString,
		PageTitle:      htmlTitleText,
		HeadersCount:   headingCountMap,
		WebLinks:       hyperlinksList,
		BrokenWebLinks: brokenLinks,
		LoginForm:      loginFormAvailable,
	}
}

func (s *DefaultWebPageAnalysisService) findBrokenHyperlinks(links []string) []string {
	var brokenLinks = make([]string, 0)

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
