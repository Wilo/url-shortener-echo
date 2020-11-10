package services

import (
	"github.com/Wilo/url-shortener-echo/domain/models"
	"github.com/Wilo/url-shortener-echo/domain/repos"
)

// UrlServices interface definitions
type UrlServices interface {
	FindByHash(Hash string) (models.Url, error)
	Create(insert models.Url) (models.Url, error)
	GetAll() ([]models.Url, error)
}

// DomainUrlServices provided for Request Management
type DomainUrlServices struct {
	urlRepo repos.UrlRepo
}

// NewUrlServices constructor for Request Service
func NewUrlServices(urlRepo repos.UrlRepo) DomainUrlServices {
	return DomainUrlServices{
		urlRepo: urlRepo,
	}
}

// GetAll obtain all data
func (u *DomainUrlServices) GetAll() ([]models.Url, error) {
	return u.urlRepo.GetAll()
}

// Create makes url short
func (u *DomainUrlServices) Create(url models.Url) (models.Url, error) {
	return u.urlRepo.Insert(url)
}

// FindByHash get the specific url by your own hash
func (u *DomainUrlServices) FindByHash(Hash string) (models.Url, error) {
	return u.urlRepo.FindByHash(Hash)
}
