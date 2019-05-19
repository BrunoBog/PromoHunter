package application

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/brunobog/Promohunter/src/model"
)

var (
	orquestrador sync.WaitGroup
)

//Hardmob Represents a website to scrap
type Hardmob struct {
	URL string `json:"URL"`
}

//HardMobScrap Get all items in start page of Hardmob
func HardMobScrap(hm *Hardmob) (Promocoes model.HardMobPromo) {

	resp, err := http.Get(hm.URL)
	if err != nil {
		log.Println("[HardMobScrap] - Não fpo possivel conectar ao promobit")
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	doc.Find("div .inner").Each(func(i int, s *goquery.Selection) {
		// For each item...
		item := model.HardMobItem{}
		item.Nome = strings.Replace(strings.Replace(s.Find("a").Text(), "\n", "", -1), "\t", "", -1)
		link, _ := s.Find("a").Attr("href")
		item.Link = strings.Replace(link, " ", "", -1)
		index := strings.Index(item.Nome, "R$")
		if index > 0 {
			item.Preco = item.Nome[index:(len(item.Nome) - 1)]
		}
		Promocoes.Oferta = append(Promocoes.Oferta, item)

	})
	return
}

func compareAndSendHardMob(desejo string, item model.HardMobItem, caminho string) {
	if strings.Contains(strings.ToUpper(item.Nome), strings.ToUpper(desejo)) {
		log.Println(item.Nome, item.Link)
		item, _ := json.Marshal(item)

		orquestrador.Add(1)
		go doRequest(item, caminho)
	}
	orquestrador.Done()
}

func filtrarDesejoHardMob(desejo string, caminho string) (err error) {
	for _, item := range HardMobScrap().Oferta {
		log.Println(item.Nome, item.Link)
		orquestrador.Add(1)
		go compareAndSendHardMob(desejo, item, caminho)
	}
	orquestrador.Done()
	return
}
