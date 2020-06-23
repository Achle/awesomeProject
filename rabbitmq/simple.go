package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func SimpleMoudle(queue string) *RMQ {
	return NewRabbitmq(queue, "", "")
}

func (r *RMQ) SimplePublisher(messgage string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	r.failed(err, "隊列申請失敗")
	err = r.channel.Publish(
		r.ExchangeName,
		r.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messgage)})
	r.failed(err, "信息發佈失敗")
}
func (r *RMQ) SimpleConsumer() {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	r.failed(err, "隊列申請失敗")
	megs, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	r.failed(err, "信息獲取失敗")

	do := make(chan bool)
	go func() {

		for d := range megs {
			log.Println(d.Body)
		}
	}()
	log.Println("Waiting Info!......")
	<-do
}
