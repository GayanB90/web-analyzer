package utils

import (
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
