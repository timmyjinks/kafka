package main

import (
	"github.com/timmyjinks/simple-kafka/consumer"
	"github.com/timmyjinks/simple-kafka/producer"
)

func main() {
	topic := "message"
	partition := 0
	c := consumer.NewConsumerService(topic, partition)
	p := producer.NewProducerService(topic, partition)

	c.Start()

	app := application{
		Consumer: c,
		Producer: p,
	}

	app.Run(":8081")
}
