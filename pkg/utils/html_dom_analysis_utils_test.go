package utils_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/GayanB90/go-web-analyzer/pkg/utils"
	"golang.org/x/net/html"
)

func TestExtractHtmlVersion(t *testing.T) {
	testCases := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "HTML5",
			html:     `<!DOCTYPE html><html><head></head><body></body></html>`,
			expected: "HTML5",
		},
		{
			name:     "HTML 4",
			html:     `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">`,
			expected: `HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"`,
		},
		{
			name:     "No Doctype Tag",
			html:     `<html><head></head><body><p></p></body></html>`,
			expected: "Error",
		},
		{
			name:     "Empty Input",
			html:     ``,
			expected: "Error",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := strings.NewReader(testCase.html)
			actual := utils.ExtractHtmlVersion(reader)
			if actual != testCase.expected {
				t.Errorf("expected %v, got %v at test case %v", testCase.expected, actual, testCase.name)
			}
		})
	}
}

func TestExtractHtmlTitle(t *testing.T) {
	testCases := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "SUCCESS",
			html:     `<!DOCTYPE html><html><head><title>test title</title></head><body></body></html>`,
			expected: "test title",
		},
		{
			name:     "No Title Tag",
			html:     `<html><head></head><body><p></p></body></html>`,
			expected: "",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := bytes.NewReader([]byte(testCase.html))
			doc, err := html.Parse(reader)
			if err != nil {
				t.Errorf("error %v", err)
			}
			actual := utils.ExtractHtmlTitleText(doc)
			if actual != testCase.expected {
				t.Errorf("expected %v, got %v at test case %v", testCase.expected, actual, testCase.name)
			}
		})
	}
}

func TestExtractHyperlinks(t *testing.T) {
	testCases := []struct {
		name     string
		html     string
		expected []string
	}{
		{
			name:     "SUCCESS",
			html:     `<!DOCTYPE html><html><head></head><body><a href="www.google.com"></a></body></html>`,
			expected: []string{"www.google.com"},
		},
		{
			name:     "No Link Tag",
			html:     `<html><head></head><body><p></p></body></html>`,
			expected: []string{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := bytes.NewReader([]byte(testCase.html))
			doc, err := html.Parse(reader)
			if err != nil {
				t.Errorf("error %v", err)
			}
			hyperlinks := make([]string, 0)
			utils.ExtractHyperlinks(doc, &hyperlinks)
			if !reflect.DeepEqual(hyperlinks, testCase.expected) {
				t.Errorf("expected %v, got %v at test case %v", testCase.expected, hyperlinks, testCase.name)
			}
		})
	}
}

func TestExtractHeadingCount(t *testing.T) {
	testCases := []struct {
		name     string
		html     string
		expected map[string]int
	}{
		{
			name: "SUCCESS",
			html: `<!DOCTYPE html><html><head></head><body><h1>test</h1><h2>test</h2><h2>test2</h2></body></html>`,
			expected: map[string]int{
				"h1": 1,
				"h2": 2,
			},
		},
		{
			name:     "No Heading Tags",
			html:     `<html><head></head><body></body></html>`,
			expected: map[string]int{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := bytes.NewReader([]byte(testCase.html))
			doc, err := html.Parse(reader)
			if err != nil {
				t.Errorf("error %v", err)
			}
			headingCount := make(map[string]int)
			utils.ExtractHeadingCount(doc, headingCount)
			if !reflect.DeepEqual(headingCount, testCase.expected) {
				t.Errorf("expected %v, got %v at test case %v", testCase.expected, headingCount, testCase.name)
			}
		})
	}
}

func TestIsLoginFormAvailable(t *testing.T) {
	testCases := []struct {
		name     string
		html     string
		expected bool
	}{
		{
			name:     "LOGIN Form Present",
			html:     `<!DOCTYPE html><html><head></head><body><form><input type="password" placeholder="password"/><button type="submit"></button></form></body></html>`,
			expected: true,
		},
		{
			name:     "Login Form Absent",
			html:     `<html><head></head><body></body></html>`,
			expected: false,
		},
		{
			name:     "Signup Form Present On Input Count",
			html:     `<!DOCTYPE html><html><head></head><body><form><input type="password" placeholder="password"/><input type="password" placeholder="confirm password"/><button type="submit"></button></form></body></html>`,
			expected: false,
		},
		{
			name:     "Signup Form Present On Signup Hint",
			html:     `<!DOCTYPE html><html><head></head><body><form><input type="password" placeholder="signup"/><input type="password" placeholder="confirm password"/><button type="submit"></button></form></body></html>`,
			expected: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := bytes.NewReader([]byte(testCase.html))
			doc, err := html.Parse(reader)
			if err != nil {
				t.Errorf("error %v", err)
			}
			actual := utils.IsLoginFormAvailable(doc)
			if actual != testCase.expected {
				t.Errorf("expected %v, got %v at test case %v", testCase.expected, actual, testCase.name)
			}
		})
	}
}
