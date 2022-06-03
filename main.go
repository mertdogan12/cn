package main

import (
	"flag"
	"fmt"
	"log"
)

var id *string

func main() {
	id = flag.String("-s", "", "Computer id")
	flag.Parse()

	if len(*id) == 0 {
		log.Fatalln("Id ist nicht gegeben")
	}

	fmt.Println(*id)
}
