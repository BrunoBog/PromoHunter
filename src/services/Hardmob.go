package services

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/brunobog/PromoHunter/src/model"
)

var (
	url = "http://www.hardmob.com.br/promocoes/"
)

// func filtrarDesejoHardMob(desejo string, caminho string) (err error) {
// 	var orquestrador sync.WaitGroup

// 	for _, item := range HardMobScrap().Oferta {
// 		log.Println(item.Nome, item.Link)
// 		orquestrador.Add(1)
// 		go compareAndSendHardMob(desejo, item, caminho)
// 	}
// 	orquestrador.Done()
// 	return
// }

// HardMobScrap this init a crawler in Hardmob
func HardMobScrap() (Promocoes model.HardMobPromo) {

	resp, err := http.Get(url)
	if err != nil {
		log.Println("[HardMobScrap] - Não foi possivel conectar ao hardmob")
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
		item.Link = url + strings.Replace(link, " ", "", -1)
		index := strings.Index(item.Nome, "R$")
		if index > 0 {
			item.Preco = item.Nome[index:(len(item.Nome) - 1)]
		}
		Promocoes.Oferta = append(Promocoes.Oferta, item)

	})
	return
}

// func compareAndSendHardMob(desejo string, item model.HardMobItem, caminho string) {
// 	if strings.Contains(strings.ToUpper(item.Nome), strings.ToUpper(desejo)) {
// 		log.Println(item.Nome, item.Link)
// 		item, _ := json.Marshal(item)
// 		// orquestrador.Add(1)
// 		// go doRequest(item, caminho)
// 	}
// orquestrador.Done()
// }
