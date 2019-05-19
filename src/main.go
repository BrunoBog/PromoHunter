package main

import (
	"sync"
	// "github.com/PuerkitoBio/goquery"
	// "strings"
	// "encoding/json"
)

var (
	orquestrador sync.WaitGroup
)

//addDesejo  -pendente
func addDesejo(json []byte) {
	return
}

func main() {

	orquestrador.Add(1)
	hm := Hardmob
	go hm.filtrarDesejoHardMob"Luva", "http://requestbin.fullcontact.com/zp13d1zp") // item - destino

	orquestrador.Add(1)
	go filtrarDesejoPromobit("Luva", "http://requestbin.fullcontact.com/zp13d1zp") // item - destino
	orquestrador.Wait()

}
