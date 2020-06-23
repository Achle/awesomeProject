package main

import (
	"awesomeProject/rabbitmq"
	"log"
)

func main() {
	mq := rabbitmq.SimpleMoudle("Test")
	mq.SimplePublisher("hello world!")
	log.Println("Publish success")
}
