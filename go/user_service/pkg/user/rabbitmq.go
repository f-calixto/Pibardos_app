package user

import (
	// Internal
	"encoding/json"
	"errors"
	"os"

	"github.com/coding-kiko/user_service/pkg/log"

	// Third party
	"github.com/streadway/amqp"
)

var (
	UsersQueue  = os.Getenv("USERS_QUEUE")
	AvatarQueue = os.Getenv("AVATAR_QUEUE")
)

// Describes any consumer for a rabbit queue

type rabbitConsumer struct {
	conn    *amqp.Connection
	service Service
	logger  log.Logger
}

type RabbitConsumer interface {
	UsersQueue()
}

func NewRabbitConsumer(service Service, conn *amqp.Connection, logger log.Logger) RabbitConsumer {
	return &rabbitConsumer{
		conn:    conn,
		service: service,
		logger:  logger,
	}
}

// Receives new users or user update from the authentication service
func (r *rabbitConsumer) UsersQueue() {
	// create rabbit connection channel
	ch, err := r.conn.Channel()
	if err != nil {
		r.logger.Error("main.go", "main", err.Error())
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		UsersQueue, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		r.logger.Error("rabbitmq.go", "NewUsersQueueConsumer", err.Error())
		panic(err.Error())
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack - TODO manual acknowledgment
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		r.logger.Error("rabbitmq.go", "NewUsersQueueConsumer", err.Error())
		panic(err.Error())
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			r.logger.Info("rabbitmq.go", "UsersQueue", "Received a message: "+string(d.Body))

			req := UpsertUserRequest{}
			err = json.Unmarshal(d.Body, &req)
			if err != nil {
				r.logger.Error("rabbitmq.go", "NewUsersQueueConsumer", err.Error())
				panic(err.Error())
			}
			_, err = r.service.UpsertUser(req)
			if err != nil { // dont akcnowledge TODO
				r.logger.Error("rabbitmq.go", "NewUsersQueueConsumer", err.Error())
				panic(err.Error())
			}
		}
	}()
	<-forever
}

// Describes any producer for a rabbit queue

type rabbitProducer struct {
	conn   *amqp.Connection
	logger log.Logger
}

type RabbitProducer interface {
	AvatarQueue(avatar []byte, filename string) error
}

func NewRabbitProducer(conn *amqp.Connection, logger log.Logger) RabbitProducer {
	return &rabbitProducer{
		conn:   conn,
		logger: logger,
	}
}

// send new avatar to be stored in the avatar service
func (r *rabbitProducer) AvatarQueue(avatar []byte, filename string) error {
	// create rabbit connection channel
	ch, err := r.conn.Channel()
	if err != nil {
		r.logger.Error("main.go", "main", err.Error())
		panic(err)
	}
	defer ch.Close()

	// declare queue
	q, err := ch.QueueDeclare(AvatarQueue, true, false, false, false, nil)
	if err != nil {
		return errors.New("failed to declare a queue")
	}

	headers := make(map[string]interface{})
	headers["filename"] = filename
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Headers:      headers,
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         avatar,
		})
	if err != nil {
		return errors.New("failed to publish message")
	}
	r.logger.Info("rabbitmq.go", "AvatarQueue", "avatar published")
	return nil
}
