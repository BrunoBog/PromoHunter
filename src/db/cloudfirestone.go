package db

import "github.com/brunobog/PromoHunter/src/model"

//https://firebase.google.com/docs/firestore/quickstart?hl=pt-br

func addWish(wish model.Desejo) {
	sa := option.WithCredentialsFile("../key.json")

	fb.app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}

// func main() {

// Use a service account
//sa := option.WithCredentialsFile("../hunterpromo-26348-firebase-adminsdk-ynyho-c8d16c9dda.json")

// fb.app, err := firebase.NewApp(context.Background(), nil, sa)
// if err != nil {
// 	log.Fatalln(err)
// }

// client, err := app.Firestore(context.Background())
// if err != nil {
// 	log.Fatalln(err)
// }

// desejo := model.Desejo{}
// desejo.Nome = "Puff"
// desejo.Nome = "Carrinho"

// user := model.User{}
// user.Email = "bog906@gmail.com"
// user.Desejos[0] = desejo

// //gravarDados(client, user)
// lerDados(client)

// }
