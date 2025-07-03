package service

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"github.com/GayanB90/go-web-analyzer/internal/model"
	"github.com/GayanB90/go-web-analyzer/pkg/utils"
	"github.com/sirupsen/logrus"
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

	data, err := fetchPage(urlString, request.RequestId)
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
		logrus.WithFields(logrus.Fields{
			"requestId":   request.RequestId,
			"htmlVersion": htmlVersion,
		}).Info("Identified html version")
		htmlVersionCh <- htmlVersion
	}()

	doc, err := s.parseHTML(data)
	if err != nil {
		return model.WebAnalysisResultModel{}, err
	}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		htmlTitleText := utils.ExtractHtmlTitleText(doc)
		logrus.WithFields(logrus.Fields{
			"requestId": request.RequestId,
			"htmlTitle": htmlTitleText,
		}).Info("Identified html title")
		pageTitleCh <- htmlTitleText
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		headingCountMap := make(map[string]int)
		utils.ExtractHeadingCount(doc, headingCountMap)
		logrus.WithFields(logrus.Fields{
			"requestId": request.RequestId,
			"headings":  headingCountMap,
		}).Info("Identified headings counts")
		headingCountCh <- headingCountMap
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		hyperlinksList := make([]string, 0)
		utils.ExtractHyperlinks(doc, &hyperlinksList)
		logrus.WithFields(logrus.Fields{
			"requestId":  request.RequestId,
			"hyperlinks": hyperlinksList,
		}).Info("Identified hyperlinks")
		hyperlinksCh <- hyperlinksList
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		hyperlinksList := <-hyperlinksCh
		brokenLinks := s.findBrokenHyperlinks(hyperlinksList)
		logrus.WithFields(logrus.Fields{
			"requestId":      request.RequestId,
			"brokenLinks":    brokenLinks,
			"hyperlinksList": hyperlinksList,
		}).Info("Identified broken hyperlinks")
		brokenLinksCh <- brokenLinks
		hyperlinksCh <- hyperlinksList
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		loginFormAvailable := utils.IsLoginFormAvailable(doc)
		logrus.WithFields(logrus.Fields{
			"requestId":          request.RequestId,
			"loginFormAvailable": loginFormAvailable,
		}).Info("Identified login form availability")
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

func fetchPage(urlString string, requestId string) ([]byte, error) {
	logrus.Infof("Fetching page %s for request id %s", urlString, requestId)
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (s *DefaultWebPageAnalysisService) parseHTML(data []byte) (*html.Node, error) {
	return html.Parse(bytes.NewReader(data))
}

func runConcurrently(tasks ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(tasks))
	for _, task := range tasks {
		go func(t func()) {
			defer waitGroup.Done()
			t()
		}(task)
	}
	waitGroup.Wait()
}

func (s *DefaultWebPageAnalysisService) findBrokenHyperlinks(links []string) []string {
	var brokenLinks = make([]string, 0)

	for _, link := range links {
		logrus.WithFields(logrus.Fields{
			"link": link,
		}).Info("Validating link")
		for _, urlValidationService := range s.UrlValidationServices {
			if urlValidationService.ValidateUrl(link) != nil {
				brokenLinks = append(brokenLinks, link)
				break
			}
		}
	}
	return brokenLinks
}
