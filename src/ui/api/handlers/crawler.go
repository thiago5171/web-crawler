package handlers

import (
	"backend_template/src/core/interfaces/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CrawlerHandler interface {
	SearchLinks(echo.Context) error
}

type crawlerHandler struct {
	service usecases.CrawlerUseCase
}

func (c crawlerHandler) SearchLinks(context echo.Context) error {

	c.service.NavigateLinks("https://changelly.com/blog/pt-br/top-10-das-melhores-criptomoedas-para-investir-em-2020/#:~:text=Existem%20novas%20moedas%20na%20lista%2C%20mas%20as%20cinco,ROI%20e%20t%C3%AAm%20o%20maior%20potencial%20de%20crescimento.")
	return context.JSON(http.StatusOK, "")
}

func NewCrawlerHandler(service usecases.CrawlerUseCase) CrawlerHandler {
	return &crawlerHandler{service}
}
