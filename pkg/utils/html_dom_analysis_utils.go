package utils

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func ExtractHtmlVersion(htmlStringReader io.Reader) string {
	tokenizer := html.NewTokenizer(htmlStringReader)
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return "Error"
		case html.DoctypeToken:
			token := tokenizer.Token()
			if token.Data == "html" {
				return "HTML5"
			} else {
				return token.Data
			}
		}
	}
}

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

func ExtractHyperlinks(node *html.Node, hyperlinks *[]string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				*hyperlinks = append(*hyperlinks, attr.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ExtractHyperlinks(c, hyperlinks)
	}
}

func ExtractHeadingCount(node *html.Node, headingCountMap map[string]int) {
	switch node.Data {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		headingCountMap[node.Data]++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ExtractHeadingCount(c, headingCountMap)
	}
}

func IsLoginFormAvailable(node *html.Node) bool {
	passwordCount, signupHint := analyzeHtmlInputTags(node)
	return passwordCount == 1 && !signupHint
}

func analyzeHtmlInputTags(node *html.Node) (int, bool) {
	passwordCount := 0
	signupHint := false
	if node.Type == html.ElementNode && node.Data == "input" {
		typ := getHtmlTagAttribute(node, "type")
		name := getHtmlTagAttribute(node, "name")
		id := getHtmlTagAttribute(node, "id")
		placeholder := getHtmlTagAttribute(node, "placeholder")

		t := strings.ToLower(typ)
		if t == "password" {
			passwordCount++
		}

		if containsSignupHint(name, id, placeholder) {
			signupHint = true
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		passwordCount2, signupHint2 := analyzeHtmlInputTags(c)
		passwordCount = passwordCount + passwordCount2
		signupHint = signupHint && signupHint2
	}
	return passwordCount, signupHint
}

func getHtmlTagAttribute(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if strings.ToLower(attr.Key) == key {
			return attr.Val
		}
	}
	return ""
}

func containsSignupHint(values ...string) bool {
	for _, v := range values {
		l := strings.ToLower(v)
		if strings.Contains(l, "signup") || strings.Contains(l, "register") || strings.Contains(l, "create") {
			return true
		}
	}
	return false
}
