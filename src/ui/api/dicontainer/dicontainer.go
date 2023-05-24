package dicontainer

import (
	"backend_template/src/core/interfaces/usecases"
	"backend_template/src/core/services"
	"backend_template/src/infra/repository/postgres"
	"backend_template/src/infra/repository/redis"
)

func AccountUseCase() usecases.AccountUseCase {
	repo := postgres.NewAccountRepository()
	return services.NewAccountService(repo)
}
func CrawlerUseCase() usecases.CrawlerUseCase {
	repo := postgres.NewCrawlerPostgresRepository()
	return services.NewCrawlerService(repo)
}

func AuthUseCase() usecases.AuthUseCase {
	repo := postgres.NewAuthPostgresRepository()
	sessionRepo := redis.NewSessionRepository()
	passwordResetRepo := redis.NewPasswordResetRepository()
	return services.NewAuthService(repo, sessionRepo, passwordResetRepo)
}

func ResourcesUseCase() usecases.ResourcesUseCase {
	repo := postgres.NewResourcesPostgresAdapter()
	return services.NewResourcesService(repo)
}
