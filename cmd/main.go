package main

import (
	"sync"
	"time"

	"github.com/timmyjinks/message-queue/rabbitmq/consumer"
	"github.com/timmyjinks/message-queue/rabbitmq/producer"
)

func main() {
	var wg sync.WaitGroup
	reciever := consumer.NewReciever()
	sender := producer.NewSender()

	wg.Go(func() {
		reciever.Start()
	})

	wg.Add(1)
	go func() {
		for {
			for i := range 100 {
				sender.Send("message" + string(i))
			}
			time.Sleep(time.Second * 10)
		}
	}()
	wg.Done()

	defer reciever.Close()
	defer sender.Close()

	wg.Wait()

	// if err := rootCmd.Execute(); err != nil {
	// 	log.Println(err)
	// }
}
