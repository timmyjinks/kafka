package main

import (
	"log"
)

func main() {
	if err := commandCmd.Execute(); err != nil {
		log.Println(err)
	}
}
