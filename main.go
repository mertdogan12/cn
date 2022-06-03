package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/mertdogan12/cn/internal/conf"
	"github.com/mertdogan12/cn/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	godotenv.Load()
	conf.Init()

	client, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer database.Disconnect(client)

	if len(*conf.ID) == 0 {
		log.Fatalln("Id ist nicht gegeben")
	}

	name, err := database.GetName(*conf.ID)
	if err == mongo.ErrNoDocuments {
		log.Fatalln("Konnte keinen Name unter der Id", *conf.ID, "finden")
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(name)
}
