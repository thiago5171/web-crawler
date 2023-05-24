package services

import (
	"backend_template/src/core/domain/crawler"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/interfaces/usecases"
	"backend_template/src/infra/mail"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type crawlerService struct {
	adapter adapters.CrawlerAdapter
}

func NewCrawlerService(adapter adapters.CrawlerAdapter) usecases.CrawlerUseCase {
	return &crawlerService{adapter}
}

var (
	links        []crawler.VisitedLinks
	hasSentEmail bool
)

func (cr crawlerService) NavigateLinks(url string) {
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

	if !cr.extractLinks(doc) {
		links = []crawler.VisitedLinks{}
		return
	}

}

func (cr crawlerService) extractLinks(node *html.Node) bool {
	if len(links) == 50 {
		return true
	}
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

				created, apiErr := cr.adapter.CreateLinks(visitedLink)
				if apiErr != nil || !created {
					if strings.Contains(apiErr.String(), "duplicate") {
						log.Printf("URL já cadastrada:\n %s \n", link.String())
						continue
					}
					log.Printf("Erro ao cadastrar URL:\n %v \n", apiErr)
					continue
				}
				if created {
					links = append(links, visitedLink)
				}
			}
			cr.NavigateLinks(link.String())
		}
	}
	if len(links) == 50 {
		if !hasSentEmail {
			go mail.SendListLinksEmail(links, "tgc1@aluno.ifal.edu.br")
			hasSentEmail = true
		}
		return false
	} else {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			cr.extractLinks(c)
		}
	}

	return true
}

//package services
//
//import (
//	"backend_template/src/core/domain/crawler"
//	"backend_template/src/core/interfaces/adapters"
//	"backend_template/src/core/interfaces/usecases"
//	"backend_template/src/infra/mail"
//	"fmt"
//	"golang.org/x/net/html"
//	"log"
//	"net/http"
//	"net/url"
//	"time"
//)
//
//type crawlerService struct {
//	adapter adapters.CrawlerAdapter
//}
//
//func NewCrawlerService(
//	adapter adapters.CrawlerAdapter,
//
//) usecases.CrawlerUseCase {
//	return &crawlerService{adapter}
//}
//
//var (
//	links   []crawler.VisitedLinks
//	visited map[string]bool = map[string]bool{}
//)
//
//func (cr crawlerService) NavigateLinks(url string) {
//	if ok := visited[url]; ok {
//		return
//	}
//	visited[url] = true
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Println(err)
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		log.Printf("status diferente de 200: %d", resp.StatusCode)
//	}
//
//	doc, err := html.Parse(resp.Body)
//	if err != nil {
//		log.Println(err)
//	}
//
//	if finalizedSearch := cr.extractLinks(doc); finalizedSearch {
//		return
//	}
//}
//
//func (cr crawlerService) extractLinks(node *html.Node) bool {
//	if node.Type == html.ElementNode && node.Data == "a" {
//		for _, attr := range node.Attr {
//			if attr.Key != "href" {
//				continue
//			}
//			link, err := url.Parse(attr.Val)
//			if err != nil || link.Scheme == "" || link.Scheme == "mailto" {
//				continue
//			}
//			fmt.Println(link.String())
//
//			visitedLink := crawler.New(nil, link.String(), link.Host, time.Now())
//			links = append(links, visitedLink)
//
//			created, apiErr := cr.adapter.CreateLinks(visitedLink)
//			if apiErr != nil || !created {
//				log.Printf("Erro ao cadastrar URL ou já foi cadastrada: %s \n", link.String())
//				continue
//			}
//
//			if len(links) == 50 {
//				go mail.SendListLinksEmail(links, "tgc1@aluno.ifal.edu.br")
//				links = []crawler.VisitedLinks{}
//				visited = map[string]bool{}
//				fmt.Println("\n\n\n\n\n\n\n AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA \n\n\n\n\n\n\n")
//				return true
//			}
//			cr.NavigateLinks(link.String())
//		}
//	}
//	for c := node.FirstChild; c != nil; c = c.NextSibling {
//		cr.extractLinks(c)
//	}
//	return false
//}
