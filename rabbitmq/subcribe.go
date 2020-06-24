package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func SubMoudle(exchange string) *RMQ {
	return NewRabbitmq("", exchange, "")
}
func (r *RMQ) SubPublisher(message string) {
	err := r.channel.ExchangeDeclarePassive(
		r.ExchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil)
	r.failed(err, "交换机创建失败")
	err = r.channel.Publish(
		r.ExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message)})
	r.failed(err, "信息发布失败")
}
func (r *RMQ) SubConsumer() {
	err := r.channel.ExchangeDeclarePassive(
		r.ExchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil)
	r.failed(err, "交换机创建失败")
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	r.failed(err, "队列申请失败")
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.ExchangeName,
		false,
		nil)
	r.failed(err, "队列绑定失败")
	megs, err := r.channel.Consume(
		q.Name,
		"",
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
