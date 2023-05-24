package usecases

type CrawlerUseCase interface {
	NavigateLinks(url string)
}
