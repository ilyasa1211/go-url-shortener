package services

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-url-shortener/internal/dto"
	"github.com/ilyasa1211/go-url-shortener/internal/entities"
	"github.com/ilyasa1211/go-url-shortener/internal/interfaces"
)

type SiteService struct {
	rs interfaces.SiteRepository
}

func NewSiteService(rs interfaces.SiteRepository) *SiteService {
	return &SiteService{rs}
}

func (s *SiteService) FindAll() *[]entities.Site {
	return s.rs.All()
}
func (s *SiteService) Create(r *http.Request) error {
	var csr dto.CreateSiteRequest

	if err := json.NewDecoder(r.Body).Decode(&csr); err != nil {
		return err
	}

	return s.rs.Create(&entities.Site{
		AliasUrl:  csr.AliasUrl,
		TargetUrl: csr.TargetUrl,
	})
}

func (s *SiteService) FindByAlias(r *http.Request) (string, error) {
	aliasUrl := r.PathValue("aliasUrl")

	site, err := s.rs.FindByAlias(aliasUrl)

	if err != nil {
		return "", err
	}

	return site.TargetUrl, nil
}
func (s *SiteService) UpdateByAlias(r *http.Request) error {
	aliasUrl := r.PathValue("aliasUrl")
	var usr dto.UpdateSiteRequest

	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		return err
	}

	return s.rs.UpdateByAlias(aliasUrl, usr.TargetUrl)
}
func (s *SiteService) DeleteByAlias(r *http.Request) error {
	aliasUrl := r.PathValue("aliasUrl")

	return s.rs.DeleteByAlias(aliasUrl)
}
