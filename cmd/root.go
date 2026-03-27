package main

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/timmyjinks/simple-kafka/consumer"
	"github.com/timmyjinks/simple-kafka/producer"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		portFlag, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Println(err)
		}

		port := ":" + strconv.Itoa(portFlag)
		topic := "message"
		partition := 0

		c := consumer.NewConsumerService(topic, partition)
		p := producer.NewProducerService(topic, partition)

		c.Start()

		app := application{
			Consumer: c,
			Producer: p,
		}

		app.Run(port)
	},
}

func init() {
	rootCmd.Flags().IntP("port", "p", 8080, "port to listen on")
}
