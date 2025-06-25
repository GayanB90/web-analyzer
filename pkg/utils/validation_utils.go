package utils

import (
	"fmt"
	"net/url"
)

func ValidateURL(urlString string) error {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return fmt.Errorf("URL Parse Error: %s", err)
	}
	return nil
}
