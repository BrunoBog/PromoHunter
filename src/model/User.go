package model

//User representa os desejos de um Usuario
type User struct {
	Email   string   `json:"Email" bson:"Email"`
	Desejos []Desejo `json:"Desejos" bson:"Desejos"`
}

func (d *Desejo) BuscarDesejos(email string) (err error) {

	return
}

func (d *Desejo) AddDesejos(email string) (err error) {

	return
}
