package main

import (
	"flag"
	"fmt"
	"log"
)

var id *int

func main() {
	id = flag.Int("t", -1, "Id vom Rechner")
	flag.Parse()

	if *id == -1 {
		log.Fatalln("Id ist nicht gegeben")
	}

	fmt.Println(*id)
}
