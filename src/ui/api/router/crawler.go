package router

import (
	"backend_template/src/ui/api/dicontainer"
	"backend_template/src/ui/api/handlers"
	"github.com/labstack/echo/v4"
)

type crawlerRouter struct {
	handler handlers.CrawlerHandler
}

func NewCrawlerRouter() Router {
	service := dicontainer.CrawlerUseCase()
	handler := handlers.NewCrawlerHandler(service)
	return &crawlerRouter{handler}
}

func (c *crawlerRouter) Load(group *echo.Group) {
	group.POST("/search-link", c.handler.SearchLinks)

}
