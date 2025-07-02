package service

import (
	"github.com/GayanB90/go-web-analyzer/pkg/utils"
	"github.com/sirupsen/logrus"
)

type LexicalUrlValidationService struct{}

func (service *LexicalUrlValidationService) ValidateUrl(url string) (err error) {
	logrus.Infof("Parsing the URL %s", url)
	err = utils.ValidateURL(url)
	if err != nil {
		logrus.Infof("Failed to parse the URL %s, error: %v", url, err.Error())
		return err
	}
	return nil
}
