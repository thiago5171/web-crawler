package handlers

import (
	"backend_template/src/core/domain/crawler"
	"backend_template/src/core/interfaces/usecases"
	"backend_template/src/infra/mail"
	"backend_template/src/ui/api/handlers/dto/request"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"time"
)

type CrawlerHandler interface {
	SearchLinks(echo.Context) error
}

type crawlerHandler struct {
	service usecases.CrawlerUseCase
}

func NewCrawlerHandler(service usecases.CrawlerUseCase) CrawlerHandler {
	return &crawlerHandler{service}
}

var (
	links             []crawler.VisitedLinks
	hasSentEmail      bool
	beforeLastTwoLink *html.Node
)

// SearchLinks
// @ID Crawler.SearchLinks
// @Summary Buscador de urls a partir da url enviada
// @Description Ao Prencher os o body da requisição com a url base, email que o usario quer receber os links encontrados e a quantidade de links que deverá ser encontrado
// @Description OBS: Recomendamos  preencher o campo de quantidade de links com até 150, estaremos trabalhando para aumentar exponencialmente esse número
// @Accept json
// @Param json body  request.SearchLink true "JSON com todos os dados necessários para que seja possivel realizar a buscas das urls"
// @Produce json
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /api/search-link [post]
func (cr crawlerHandler) SearchLinks(context echo.Context) error {
	var body request.SearchLink
	if err := context.Bind(&body); err != nil {
		return unsupportedMediaTypeError
	}
	cr.NavigateLinks(body.Url, body.Email, body.NumberLinks)
	return context.JSON(http.StatusCreated, "Links cadastrados e enviados ao email com sucesso")
}

func (cr crawlerHandler) NavigateLinks(url, email string, numberLinks int) {
	parsedUrl := cr.validateUrl(url)

	if !cr.linkJumper(parsedUrl, email, numberLinks) {
		links = []crawler.VisitedLinks{}
		return
	}
}

func (cr crawlerHandler) validateUrl(url string) *html.Node {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("status diferente de 200: %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return doc
}

func (cr crawlerHandler) linkJumper(node *html.Node, email string, numberLinks int) bool {
	if len(links) == numberLinks {
		return true
	}
	fmt.Println(len(links))
	cr.extractLink(node, email, numberLinks)

	if len(links) == numberLinks {
		if !hasSentEmail {
			go mail.SendListLinksEmail(links, email)
			hasSentEmail = true
		}
		return false
	} else {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			cr.linkJumper(c, email, numberLinks)
		}
	}

	return true
}

func (cr crawlerHandler) extractLink(node *html.Node, email string, numberLinks int) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" || link.Scheme == "mailto" {
				continue
			}
			fmt.Println(link.String())
			if !linq.From(links).Contains(link.String()) {
				visitedLink := crawler.New(nil, link.String(), link.Host, time.Now())
				created := cr.service.CreateLink(visitedLink)

				if created {
					links = append(links, visitedLink)
				} else {
					continue
				}
			}
			cr.NavigateLinks(link.String(), email, numberLinks)
		}
	}
}
