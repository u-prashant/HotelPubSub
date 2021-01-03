package messagequeue

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// RabbitMqConfig are the configurations required to connect to
// rabbit message queue
type RabbitMqConfig struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

// RabbitMq contains the client using which a message can be produced or consumed
type RabbitMq struct {
	client *amqp.Connection
}

// GetNewRabbitMqObject creates a new rabbit mq object
func GetNewRabbitMqObject() *RabbitMq {
	mq := &RabbitMq{}
	return mq
}

// Initialise initialises the connection to the rabbitmq queue
func (mq *RabbitMq) Initialise(config interface{}) error {
	rabbitMqConfig := config.(RabbitMqConfig)
	url := "amqp://" + rabbitMqConfig.Address + ":" + rabbitMqConfig.Port
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Errorf("error in connecting to [%s] - err[%s]", url, err.Error())
		return err
	}
	log.Infof("successfully connected to [%s]", url)
	mq.client = conn
	return nil
}

// Publish publishes the data onto the queue
func (mq *RabbitMq) Publish(queueName string, data []byte) error {
	ch, err := mq.client.Channel()
	if err != nil {
		log.Errorf("error in creating channel - [%s]", err.Error())
		return err
	}
	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Errorf("error in creating queue - [%s]", err.Error())
		return err
	}
	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})
	if err != nil {
		log.Errorf("error in publishing message - [%s]", err.Error())
		return err
	}
	log.Infof("successfully published to queue [%s]", queueName)
	return nil
}

// Subscribe subscribes the message for the queue name provided and
// calls the handler func in case the message is received onto the queue
func (mq *RabbitMq) Subscribe(queueName string, handler Handler) error {
	ch, err := mq.client.Channel()
	if err != nil {
		log.Errorf("error in creating channel - [%s]", err.Error())
		return err
	}
	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Errorf("error in creating queue - [%s]", err.Error())
		return err
	}
	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Errorf("error registering consumer - [%s]", err.Error())
		return err
	}
	log.Infof("successfully subsribed to [%s]", queueName)
	for msg := range msgs {
		ack := handler(msg.Body)
		if ack {
			msg.Ack(ack)
		}
	}
	return nil
}
