package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mertdogan12/cn/internal/api"
	"github.com/mertdogan12/cn/internal/conf"
	"github.com/mertdogan12/cn/internal/database"
)

func main() {
	godotenv.Load()
	conf.Init()

	client, err := database.Connect()
	if err != nil {
		panic(err)
	}

	defer database.Disconnect(client)
	fmt.Println("Server started at", *conf.Port)

	http.HandleFunc("/getname", api.GetName)
	http.ListenAndServe(":"+*conf.Port, nil)
}
