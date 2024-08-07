package rabbitmq

import (
	"mailer-rabbitmq/internal/mailer"

	"github.com/streadway/amqp"
)

type RabbitMqConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  Config
	mailer  mailer.Mailer
}

func NewRabbitMQConsumer(config Config) (*RabbitMqConsumer, error) {
	conn, err := amqp.Dial(config.URI)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = channel.ExchangeDeclare(
		config.Exchange,
		config.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &RabbitMqConsumer{
		conn:    conn,
		channel: channel,
		config:  config,
		mailer:  mailer.Mailer{},
	}, nil

}
