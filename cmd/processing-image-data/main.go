package main

import (
	"fmt"
	"log"

	"github.com/Sergei3232/Getting-image-data/config"
)

func main() {
	conf, err := config.NenConfig()
	if err != nil {
		log.Panicln(err.Error())
	}
	fmt.Println(conf)
}
