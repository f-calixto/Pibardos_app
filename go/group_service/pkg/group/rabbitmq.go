package group

import (
	// Internal

	"os"

	// Internal
	"github.com/coding-kiko/group_service/pkg/errors"
	"github.com/coding-kiko/group_service/pkg/log"

	// Third party
	"github.com/streadway/amqp"
)

var (
	AvatarQueue = os.Getenv("AVATAR_QUEUE")
)

// Describes any producer for a rabbit queue
type rabbitProducer struct {
	conn   *amqp.Connection
	logger log.Logger
}

type RabbitProducer interface {
	AvatarQueue(file File) error
}

func NewRabbitProducer(conn *amqp.Connection, logger log.Logger) RabbitProducer {
	return &rabbitProducer{
		conn:   conn,
		logger: logger,
	}
}

// send new avatar to be stored in the avatar service
func (r *rabbitProducer) AvatarQueue(file File) error {
	// create rabbit connection channel
	ch, err := r.conn.Channel()
	if err != nil {
		return errors.NewRabbitError("failed to create channel")
	}
	defer ch.Close()

	// declare queue
	q, err := ch.QueueDeclare(AvatarQueue, true, false, false, false, nil)
	if err != nil {
		return errors.NewRabbitError("failed to declare a queue")
	}

	headers := make(map[string]interface{})
	headers["filename"] = file.Name
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Headers:      headers,
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         file.Data,
		})
	if err != nil {
		return errors.NewRabbitError("failed to publish message")
	}
	r.logger.Info("rabbitmq.go", "AvatarQueue", "avatar published")
	return nil
}
