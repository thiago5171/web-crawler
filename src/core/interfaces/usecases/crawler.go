package usecases

import "backend_template/src/core/domain/crawler"

type CrawlerUseCase interface {
	CreateLink(visitedLink crawler.VisitedLinks) bool
}
