package service

import (
	"log"

	"github.com/GayanB90/go-web-analyzer/pkg/utils"
)

type LexicalUrlValidationService struct{}

func (service *LexicalUrlValidationService) ValidateUrl(url string) (err error) {
	log.Printf("Parsing the URL %s", url)
	err = utils.ValidateURL(url)
	if err != nil {
		log.Printf("Failed to parse the URL %s, error: %v", url, err.Error())
		return err
	}
	return nil
}
