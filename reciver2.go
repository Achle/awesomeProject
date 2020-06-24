package main

import "awesomeProject/rabbitmq"

func main() {
	mq := rabbitmq.SubMoudle("Test2")
	mq.SubConsumer()
}
