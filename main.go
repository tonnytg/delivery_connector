package main

import (
	"log"

	"github.com/tonnytg/delivery_connector/internal/infra/event_consumer"
)

func main() {
	log.Println("start event delivery")

	ed := event_consumer.NewEventDelivery() // Mudando o nome da variável
	ed.Start()
}
