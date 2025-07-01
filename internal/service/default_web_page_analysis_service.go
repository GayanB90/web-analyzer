package service

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/pkg/utils"
	"golang.org/x/net/html"
)

type DefaultWebPageAnalysisService struct {
	UrlValidationServices []UrlValidationService
}

func (s *DefaultWebPageAnalysisService) AnalyzeWebPage(request model.WebAnalysisRequestModel) (model.WebAnalysisResultModel, error) {
	var urlString = request.WebUrl

	err := utils.ValidateURL(urlString)
	if err != nil {
		return model.WebAnalysisResultModel{
			ValidationErrors: []string{err.Error()},
		}, err
	}
	resp, err := http.Get(urlString)
	if err != nil {
		return model.WebAnalysisResultModel{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.WebAnalysisResultModel{}, err
	}

	var waitGroup sync.WaitGroup

	htmlVersionCh := make(chan string, 1)
	pageTitleCh := make(chan string, 1)
	headingCountCh := make(chan map[string]int, 1)
	hyperlinksCh := make(chan []string, 1)
	brokenLinksCh := make(chan []string, 1)
	loginFormCh := make(chan bool, 1)

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		htmlVersion := utils.ExtractHtmlVersion(bytes.NewReader(data))
		log.Printf("htmlVersion: %v", htmlVersion)
		htmlVersionCh <- htmlVersion
	}()

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return model.WebAnalysisResultModel{}, err
	}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		htmlTitleText := utils.ExtractHtmlTitleText(doc)
		log.Printf("htmlTitleText: %v", htmlTitleText)
		pageTitleCh <- htmlTitleText
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		headingCountMap := make(map[string]int)
		utils.ExtractHeadingCount(doc, headingCountMap)
		log.Printf("headingCountMap: %v", headingCountMap)
		headingCountCh <- headingCountMap
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		hyperlinksList := make([]string, 0)
		utils.ExtractHyperlinks(doc, &hyperlinksList)
		log.Printf("hyperlinks: %v", hyperlinksList)
		hyperlinksCh <- hyperlinksList
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		hyperlinksList := <-hyperlinksCh
		brokenLinks := s.findBrokenHyperlinks(hyperlinksList)
		log.Printf("brokenLinks: %v", brokenLinks)
		brokenLinksCh <- brokenLinks
		hyperlinksCh <- hyperlinksList
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		loginFormAvailable := utils.IsLoginFormAvailable(doc)
		log.Printf("loginFormAvailable: %v", loginFormAvailable)
		loginFormCh <- loginFormAvailable
	}()

	waitGroup.Wait()

	htmlVersion := <-htmlVersionCh
	pageTitle := <-pageTitleCh
	headingCount := <-headingCountCh
	hyperlinksList := <-hyperlinksCh // we took it back after brokenLinks
	brokenLinks := <-brokenLinksCh
	loginFormAvailable := <-loginFormCh

	return model.WebAnalysisResultModel{
		RequestId:      request.RequestId,
		WebUrl:         urlString,
		HtmlVersion:    htmlVersion,
		PageTitle:      pageTitle,
		HeadersCount:   headingCount,
		WebLinks:       hyperlinksList,
		BrokenWebLinks: brokenLinks,
		LoginForm:      loginFormAvailable,
	}, nil
}

func (s *DefaultWebPageAnalysisService) findBrokenHyperlinks(links []string) []string {
	var brokenLinks = make([]string, 0)

	for _, link := range links {
		log.Printf("Validating hyperlink: %v", link)
		for _, urlValidationService := range s.UrlValidationServices {
			if urlValidationService.ValidateUrl(link) != nil {
				brokenLinks = append(brokenLinks, link)
				break
			}
		}
	}
	return brokenLinks
}
