package main

//https://firebase.google.com/docs/firestore/quickstart?hl=pt-br

import (
	"cloud.google.com/go/firestore"
	"log"

	"../model"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {

	// Use a service account
	sa := option.WithCredentialsFile("chave.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	desejo := model.Desejo{}
	desejo.Nome = "Puff"
	desejo.Nome = "Carrinho"

	user := model.User{}
	user.Email = "bog906@gmail.com"
	user.Desejos[0] = desejo

	//gravarDados(client, user)
	lerDados(client)

}

// I Don´t know how this works... https://godoc.org/firebase.google.com/go#App
func gravarDados(client *firestore.Client, dados map[string]interface{}) {

	_, _, err := client.Collection("users").Add(context.Background(), dados)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}

}

func lerDados(client *firestore.Client) {
	iter := client.Collection("users").Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		log.Println(doc.Data())
	}
}
