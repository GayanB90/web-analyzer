package utils

import (
	"fmt"
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
	var hyperlinks = make([]string, 0)
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
	return hyperlinks
}

func ExtractHeadingCount(node *html.Node, headingCountMap map[string]int) {
	switch node.Data {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		level := node.Data
		levelCount, exists := headingCountMap[level]
		if exists {
			levelCount++
		} else {
			headingCountMap[level] = 1
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			ExtractHeadingCount(c, headingCountMap)
		}
		fmt.Printf("Heading level: %s, Count: %s\n", level, headingCountMap[level])
	}
}

func IsLoginFormAvailable(node *html.Node) bool {
	passwordCount, signupHint, loginHint := analyzeHtmlInputTags(node)
	return passwordCount == 1 && signupHint == false && loginHint == true
}

func analyzeHtmlInputTags(node *html.Node) (int, bool, bool) {
	passwordCount := 0
	signupHint := false
	loginHint := false
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
		if containsLoginHint(name, id, placeholder) {
			loginHint = true
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		analyzeHtmlInputTags(c)
	}
	return passwordCount, signupHint, loginHint
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

func containsLoginHint(values ...string) bool {
	for _, v := range values {
		l := strings.ToLower(v)
		if strings.Contains(l, "login") || strings.Contains(l, "signin") || strings.Contains(l, "log in") {
			return true
		}
	}
	return false
}
