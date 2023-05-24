package dicontainer

import (
	"backend_template/src/core/interfaces/usecases"
	"backend_template/src/core/services"
	"backend_template/src/infra/repository/postgres"
)

func CrawlerUseCase() usecases.CrawlerUseCase {
	repo := postgres.NewCrawlerPostgresRepository()
	return services.NewCrawlerService(repo)
}
