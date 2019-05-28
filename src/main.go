package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/brunobog/PromoHunter/src/services"
)

var (
	orquestrador sync.WaitGroup
)

func doRequest(json []byte, caminho string) {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	// estruturando o request com autenticação
	request, err := http.NewRequest("POST", caminho, bytes.NewBuffer(json))
	if err != nil {
		return
		//println("[doRequest], erro ao tentar fazer o request para o requestbin")
	}
	//usuario e senha aqui
	request.SetBasicAuth("fizz", "buzz")
	// para enviar um content type temos que add ele noo request
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	defer resposta.Body.Close()
	if err != nil {
		orquestrador.Done()
		return
		//println("[doRequest], deu erro ao tentar executar o request")
	}
	if resposta.StatusCode == 200 {
		//println(resposta.Status)

		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			orquestrador.Done()
			//println("[doRequest] Erro ao ler o corpo da resposta")
			return
		}
		log.Println(string(corpo))
	}

	orquestrador.Done()
}

//addDesejo  -pendente
func addDesejo(json []byte) {
	return
}

func main() {

	hardMobOffers := services.HardMobScrap().Oferta
	for _, item := range hardMobOffer {
		log.Println(item.Nome, item.Link)
	}

	promobitOffers := services.PromoBitScrap().Oferta
	for _, item := range promobitOffers {
		log.Println(item.Nome, item.Link)
	}

}
