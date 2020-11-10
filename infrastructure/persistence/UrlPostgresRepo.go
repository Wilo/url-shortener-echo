package persistence

import (
	"errors"

	"github.com/Wilo/url-shortener-echo/domain/models"
	"github.com/Wilo/url-shortener-echo/domain/repos"
	"gorm.io/gorm"
)

// UrlPostgresRepo implementation of repo with postgres db
type UrlPostgresRepo struct {
	connection *gorm.DB
}

// NewUrlPostgresRepo constructor function of UrlPostgresRepo
func NewUrlPostgresRepo(connection *gorm.DB) repos.UrlRepo {
	return &UrlPostgresRepo{
		connection: connection,
	}
}

// FindByHash implementation with postgres
func (u *UrlPostgresRepo) FindByHash(Hash string) (url models.Url, err error) {
	err = u.connection.Where("hash = ?", Hash).Take(&url).Error
	if err != nil {
		return models.Url{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Url{}, errors.New("Url not found")
	}
	return url, err
}

// GetAll implementation with postgres
func (u *UrlPostgresRepo) GetAll() (urls []models.Url, err error) {
	err = u.connection.Find(&urls).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("urls not found")
	}
	return urls, err
}

// Insert create new
func (u *UrlPostgresRepo) Insert(url models.Url) (models.Url, error) {
	err := u.connection.Create(&url).Error
	if err != nil {
		return models.Url{}, err
	}
	return url, err
}
