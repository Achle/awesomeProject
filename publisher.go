package main

import (
	"awesomeProject/rabbitmq"
	"log"
	"strconv"
)

func main() {
	mq := rabbitmq.SimpleMoudle("Test")
	for i := 0; i < 100; i++ {
		mq.SimplePublisher("message: " + strconv.Itoa(i))
	}

	log.Println("Publish success")

}
