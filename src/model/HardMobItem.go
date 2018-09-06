package model

//PromobitItem Define um item encontrado na pagina
type HardMobItem struct {
	Nome  string `json:"Nome"`
	Link  string `json:"Link"`
	Preco string `json:"Preco"`
}

//PromobitPromo Guarda as promoções do sitepromobit
type HardMobPromo struct {
	Oferta []HardMobItem `json:"Oferta"`
}
