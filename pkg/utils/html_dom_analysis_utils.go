package utils

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func ExtractHtmlTitleText(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
		return strings.TrimSpace(node.FirstChild.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if title := ExtractHtmlTitleText(c); title != "" {
			return title
		}
	}
	return ""
}

func ExtractHyperlinks(node *html.Node) []string {
	var hyperlinks []string = make([]string, 0)
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				hyperlinks = append(hyperlinks, attr.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ExtractHyperlinks(c)
	}
	log.Printf("Successfully extracted hyperlinks: %v", hyperlinks)
	return hyperlinks
}
