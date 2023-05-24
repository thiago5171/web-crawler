package postgres

import (
	"backend_template/src/core/domain/crawler"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/infra/repository"
	"backend_template/src/infra/repository/postgres/query"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type crawlerPostgresRepository struct{}

func NewCrawlerPostgresRepository() adapters.CrawlerAdapter {
	return &crawlerPostgresRepository{}
}

func (c *crawlerPostgresRepository) CreateLinks(link crawler.VisitedLinks) (bool, errors.Error) {
	err := defaultExecQuery(query.Crawler().Create(), link.Url(), link.Website(), link.CheckedDate())
	if err != nil {
		logger.Error().Msg(err.String())
		return false, err
	}
	return true, nil
}

func (c *crawlerPostgresRepository) ListLinks() ([]crawler.VisitedLinks, errors.Error) {
	rows, err := repository.Queryx(query.Crawler().List())
	if err != nil {
		return nil, err
	}
	var links = []crawler.VisitedLinks{}
	for rows.Next() {
		var serializedCrawler = map[string]interface{}{}
		rows.MapScan(serializedCrawler)
		link, err := newCrawlerFromMapRows(serializedCrawler)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, errors.NewUnexpected()
		}
		links = append(links, link)
	}
	return links, nil
}

func newCrawlerFromMapRows(data map[string]interface{}) (crawler.VisitedLinks, errors.Error) {
	var err error
	var id uuid.UUID
	var url = fmt.Sprint(data["url"])
	var website = fmt.Sprint(data["website"])
	var checkedDate = fmt.Sprint(data["checked_date"])
	convertDate, err := time.Parse("2006-01-02", checkedDate)
	if err != nil {
		return nil, errors.NewUnexpected()
	}
	id, err = uuid.Parse(string(data["id"].([]uint8)))
	if err != nil {
		return nil, errors.NewUnexpected()
	}
	return crawler.New(&id, url, website, convertDate), nil
}
