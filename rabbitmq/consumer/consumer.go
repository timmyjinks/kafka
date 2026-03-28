package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/timmyjinks/message-queue/rabbitmq/util"
	"log"
)

type Reciever struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewReciever() *Reciever {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")

	util.FailOnError(err, "Failed to declare a queue")

	return &Reciever{
		Conn:    conn,
		Channel: ch,
	}
}

func (r *Reciever) Start() {
	q, err := r.Channel.QueueDeclare(
		"hello", // name
		true,    // durability
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)

	msgs, err := r.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func (r *Reciever) Close() {
	defer r.Conn.Close()
	defer r.Channel.Close()
}
