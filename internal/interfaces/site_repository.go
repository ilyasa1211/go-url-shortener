package interfaces

import "github.com/ilyasa1211/go-url-shortener/internal/entities"

type SiteRepository interface {
	All() *[]entities.Site
	FindByAlias(aliasUrl string) (*entities.Site, error)
	Create(site *entities.Site) error
	UpdateByAlias(aliasUrl string, targetUrl string) error
	DeleteByAlias(aliasUrl string) error
}
