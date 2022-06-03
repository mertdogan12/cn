package main

import (
	"flag"
	"fmt"
	"log"
)

var id *string
	"github.com/mertdogan12/cn/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	id = flag.String("-s", "", "Computer id")
	flag.Parse()

	if len(*id) == 0 {
	client, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer database.Disconnect(client)
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
