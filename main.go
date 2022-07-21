package main

import (
	"log"

	"github.com/awesome-sphere/as-booking-consumer/db/models"
	"github.com/awesome-sphere/as-booking-consumer/kafka"
)

func main() {

	models.InitDatabase()
	log.Println("Starting kafka...")
	kafka.InitKafkaTopic()
	for {
	}

}
