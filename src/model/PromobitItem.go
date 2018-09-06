package model

//PromobitItem Define um item encontrado na pagina
type PromobitItem struct {
	Nome  string `json:"Nome"`
	Link  string `json:"Link"`
	Preco string `json:"Preco"`
	Loja  string `json:"Loja"`
}

//PromobitPromo Guarda as promoções do sitepromobit
type PromobitPromo struct {
	Oferta []PromobitItem `json:"Oferta"`
}
