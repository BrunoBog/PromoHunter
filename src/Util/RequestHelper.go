package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//RequestHelper do requests ...
type RequestHelper struct {
}

//DoRequest Do requests do destny api
func DoRequest(json []byte, caminho string) {

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

	if err != nil {
		println("[doRequest], deu erro ao tentar executar o request")
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		log.Println(resposta.Status)

		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			log.Println("[doRequest] Erro ao ler o corpo da resposta")
			return
		}
		log.Println(string(corpo))
	}

}
