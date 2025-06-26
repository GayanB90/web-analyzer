package service

import (
	"log"
	"net/http"
)

type HttpUrlValidationService struct{}

func (service *HttpUrlValidationService) ValidateUrl(url string) (err error) {
	log.Printf("invoking GET request for URL %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	log.Printf("finished invoking GET request for URL %s", url)
	defer resp.Body.Close()
	return nil
}
