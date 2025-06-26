package service

type UrlValidationService interface {
	ValidateUrl(url string) (err error)
}
