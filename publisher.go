package main

import (
	"awesomeProject/rabbitmq"
	"log"
	"strconv"
)

func main() {
	mq := rabbitmq.SubMoudle("Test")
	for i := 0; i < 100; i++ {
		mq.SubPublisher("message: " + strconv.Itoa(i))
	}

	log.Println("Publish success")

}
