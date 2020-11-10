package controllers

import (
	"net/http"

	"github.com/Wilo/url-shortener-echo/domain/models"
	"github.com/Wilo/url-shortener-echo/domain/services"
	"github.com/labstack/echo/v4"
)

type UrlController struct {
	urlService services.DomainUrlServices
}

// NewRequestController constructor for RequestController
func NewUrlController(urlService services.DomainUrlServices) UrlController {
	return UrlController{
		urlService: urlService,
	}
}

// InitUrlControllerRouting define the routes for the controller endpoints
func (u *UrlController) InitUrlControllerRouting(e *echo.Echo) {
	e.GET("/:hash", u.findByHashEndpoint)
	e.GET("/all", u.getAllEndpoint)
	e.POST("/", u.insertEndpoint)
}

func (u *UrlController) findByHashEndpoint(ctx echo.Context) error {
	hash := ctx.Param("hash")
	if hash != "" {
		url, err := u.urlService.FindByHash(hash)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		return ctx.Redirect(http.StatusTemporaryRedirect, url.Link)
		//return ctx.JSON(http.StatusOK, url)
	}
	return ctx.JSON(http.StatusBadRequest, "Url not found!")
}

func (u *UrlController) getAllEndpoint(ctx echo.Context) error {
	urls, err := u.urlService.GetAll()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, urls)
}

func (u *UrlController) insertEndpoint(ctx echo.Context) error {
	url := new(models.Url)
	if err := ctx.Bind(url); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	created, err := u.urlService.Create(*url)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusCreated, created)
}
