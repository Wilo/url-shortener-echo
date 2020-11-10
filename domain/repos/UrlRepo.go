package repos

import "github.com/Wilo/url-shortener-echo/domain/models"

type UrlRepo interface {
	FindByHash(Hash string) (models.Url, error)
	Insert(insert models.Url) (models.Url, error)
	GetAll() ([]models.Url, error)
}
