package producer

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/timmyjinks/message-queue/rabbitmq/util"
)

type Sender struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewSender() *Sender {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")

	return &Sender{
		Conn:    conn,
		Channel: ch,
	}
}

func (s *Sender) Send(body string) {
	q, err := s.Channel.QueueDeclare(
		"hello", // name
		true,    // durability
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	util.FailOnError(err, "Failed to declare a queue")

	log.Println("sending")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func (s *Sender) Close() {
	defer s.Conn.Close()
	defer s.Channel.Close()
}
