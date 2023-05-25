package services

import (
	"backend_template/src/core/domain/crawler"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/interfaces/usecases"
	"log"
	"strings"
)

type crawlerService struct {
	adapter adapters.CrawlerAdapter
}

func NewCrawlerService(adapter adapters.CrawlerAdapter) usecases.CrawlerUseCase {
	return &crawlerService{adapter}
}

func (cr crawlerService) CreateLink(visitedLink crawler.VisitedLinks) bool {
	created, apiErr := cr.adapter.CreateLinks(visitedLink)
	if apiErr != nil || !created {
		if strings.Contains(apiErr.String(), "duplicate") {
			log.Printf("URL já cadastrada:\n %s \n", visitedLink.Url())
			return false
		}
		log.Printf("Erro ao cadastrar URL:\n %v \n", apiErr.String())
		return false
	}

	return true
}
