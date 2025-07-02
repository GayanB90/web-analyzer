package service

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type HttpUrlValidationService struct{}

func (service *HttpUrlValidationService) ValidateUrl(url string) (err error) {
	logrus.Infof("invoking GET request for URL %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	logrus.Infof("finished invoking GET request for URL %s", url)
	defer resp.Body.Close()
	return nil
}
