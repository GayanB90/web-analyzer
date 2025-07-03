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

	var (
		htmlVersion        string
		pageTitle          string
		headingCounts      map[string]int
		hyperlinks         []string
		brokenLinks        []string
		loginFormAvailable bool
	)

	doc, err := s.parseHTML(data)

	runConcurrently(
		func() { htmlVersion = s.analyzeHtmlVersion(data) },
		func() { pageTitle = s.analyzeTitle(doc) },
		func() { headingCounts = s.analyzeHeadings(doc) },
		func() { hyperlinks = s.analyzeHyperlinks(doc) },
		func() { loginFormAvailable = s.analyzeLoginForm(doc) },
	)

	brokenLinks = s.findBrokenHyperlinks(hyperlinks)

	return model.WebAnalysisResultModel{
		RequestId:      request.RequestId,
		WebUrl:         urlString,
		HtmlVersion:    htmlVersion,
		PageTitle:      pageTitle,
		HeadersCount:   headingCounts,
		WebLinks:       hyperlinks,
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

func (s *DefaultWebPageAnalysisService) analyzeHtmlVersion(data []byte) string {
	return utils.ExtractHtmlVersion(bytes.NewReader(data))
}

func (s *DefaultWebPageAnalysisService) analyzeTitle(doc *html.Node) string {
	return utils.ExtractHtmlTitleText(doc)
}

func (s *DefaultWebPageAnalysisService) analyzeHeadings(doc *html.Node) map[string]int {
	headingCountMap := make(map[string]int)
	utils.ExtractHeadingCount(doc, headingCountMap)
	return headingCountMap
}

func (s *DefaultWebPageAnalysisService) analyzeHyperlinks(doc *html.Node) []string {
	var hyperLinks []string
	utils.ExtractHyperlinks(doc, &hyperLinks)
	return hyperLinks
}

func (s *DefaultWebPageAnalysisService) analyzeLoginForm(doc *html.Node) bool {
	return utils.IsLoginFormAvailable(doc)
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
