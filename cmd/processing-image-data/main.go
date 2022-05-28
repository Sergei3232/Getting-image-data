package main

import (
	"github.com/Sergei3232/Getting-image-data/internal/app/processor"
	"log"

	"github.com/Sergei3232/Getting-image-data/config"
)

func main() {
	conf, err := config.NenConfig()
	if err != nil {
		log.Panicln(err.Error())
	}

	processorClient, err := processor.NewProcessor(conf)
	if err != nil {
		log.Panicln(err)
	}
	processorClient.Run()
}
