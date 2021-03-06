package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func SimpleMoudle(queue string) *RMQ {
	return NewRabbitmq(queue, "", "")
}

func (r *RMQ) SimplePublisher(message string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	r.failed(err, "队列申请失败")
	err = r.channel.Publish(
		r.ExchangeName,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message)})
	r.failed(err, "信息发布失败")
}
func (r *RMQ) SimpleConsumer() {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	r.failed(err, "队列申请失败")
	megs, err := r.channel.Consume(
		r.QueueName,
		"test",
		true,
		false,
		false,
		false,
		nil)
	r.failed(err, "队列申请失败")

	do := make(chan bool)
	go func() {
		for d := range megs {
			log.Printf("Recive:%s", d.Body)
		}
	}()
	log.Println("Waiting Info!......")
	<-do
}
