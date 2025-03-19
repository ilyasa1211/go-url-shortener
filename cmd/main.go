package main

import (
	"log"
	"net/http"

	"github.com/ilyasa1211/go-url-shortener/internal/infrastructure/database/sqlite"
	handler "github.com/ilyasa1211/go-url-shortener/internal/infrastructure/http"
	"github.com/ilyasa1211/go-url-shortener/internal/services"
)

func main() {
	db := sqlite.Connect()

	siteRepo := sqlite.NewSiteRepository(db)
	siteService := services.NewSiteService(siteRepo)
	siteHandler := handler.NewSiteHandler(siteService)

	http.HandleFunc("GET /sites", siteHandler.Index)
	http.HandleFunc("GET /sites/{aliasUrl}", siteHandler.Show)
	http.HandleFunc("POST /sites", siteHandler.Create)
	http.HandleFunc("PUT /sites/{aliasUrl}", siteHandler.Update)
	http.HandleFunc("DELETE /sites/{aliasUrl}", siteHandler.Delete)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
