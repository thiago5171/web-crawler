package adapters

import (
	"backend_template/src/core/domain/crawler"
	"backend_template/src/core/domain/errors"
)

type CrawlerAdapter interface {
	CreateLinks(links crawler.VisitedLinks) (bool, errors.Error)
	ListLinks() ([]crawler.VisitedLinks, errors.Error)
}
