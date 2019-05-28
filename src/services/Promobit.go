package services

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/brunobog/PromoHunter/src/model"
)

var (
	promobitURL = "https://www.promobit.com.br/"
)

//PromoBitScrap do scrap in promobit
func PromoBitScrap() (Promocoes model.PromobitPromo) {

	resp, err := http.Get(promobitURL)
	if err != nil {
		log.Println("[PromoBitScrap] - Não foi possivel conectar ao promobit")
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// digite a tag e .nome da classe
	doc.Find("div .pr-tl-card").Each(func(i int, s *goquery.Selection) {
		// For each item...
		item := model.PromobitItem{}
		item.Nome = strings.Replace(strings.Replace(s.Find("a").Text(), "\n", "", -1), "\t", "", -1)
		item.Preco = s.Find("div .price").Text()
		item.Loja = s.Find("div .where").Text()
		link, _ := s.Find("a").Attr("href")
		item.Link = "www.promobit.com.br" + strings.Replace(link, " ", "", -1)
		Promocoes.Oferta = append(Promocoes.Oferta, item)
	})
	return
}

// func compareAndSendPromobit(desejo string, item model.PromobitItem, caminho string) {
// 	log.Println(item.Nome, item.Link)
// 	if strings.Contains(strings.ToUpper(item.Nome), strings.ToUpper(desejo)) {
// 		log.Println(item.Nome, item.Link)
// 		item, _ := json.Marshal(item)
// 		orquestrador.Add(1)
// 		go doRequest(item, caminho)
// 	}
// 	orquestrador.Done()
// }

// func filtrarDesejoPromobit(desejo string, caminho string) (err error) {
// 	for _, item := range PromoBitScrap().Oferta {
// 		orquestrador.Add(1)
// 		go compareAndSendPromobit(desejo, item, caminho)
// 	}
// 	orquestrador.Done()
// 	return
// }
