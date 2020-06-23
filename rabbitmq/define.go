package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

const MYURL = "amqp://rabbit:rabbit@0.0.0.0:5672/achelous"

type RMQ struct {
	URL          string
	conn         *amqp.Connection
	channel      *amqp.Channel
	QueueName    string
	ExchangeName string
	RoutingKey   string
}

func (r *RMQ) failed(err error, message string) {

	if err != nil {
		log.Printf("%s:%s", message, err)
	}
}

func (r *RMQ) Destory() {
	err := r.channel.Close()
	r.failed(err, "信道關閉時便")
	err = r.conn.Close()
	r.failed(err, "連接關閉失敗")
}

func NewRabbitmq(queue, exchange, key string) *RMQ {

	rmq := &RMQ{
		URL:          MYURL,
		QueueName:    queue,
		ExchangeName: exchange,
		RoutingKey:   key}

	var err error = nil
	rmq.conn, err = amqp.Dial(rmq.URL)
	rmq.failed(err, "連接失敗")
	rmq.channel, err = rmq.conn.Channel()
	rmq.failed(err, "新到創建失敗")

	return rmq
}
