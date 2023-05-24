package query

type CrawlerQueryBuilder interface {
	Create() string
	List() string
}

type crawlerQueryBuilder struct{}

func (*crawlerQueryBuilder) List() string {
	return `SELECT id,url, website, checked_date FROM links`
}

func (*crawlerQueryBuilder) Create() string {
	return `INSERT INTO links (url, website, checked_date) VALUES ($1, $2, $3);`
}

func Crawler() CrawlerQueryBuilder {
	return &crawlerQueryBuilder{}
}
