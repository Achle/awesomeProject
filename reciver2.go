package main

import "awesomeProject/rabbitmq"

func main() {
	mq := rabbitmq.SimpleMoudle("Test")
	mq.SimpleConsumer()
}
