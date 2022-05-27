package avatar

import (
	// std lib
	"fmt"
	"os"

	// Internal
	"github.com/coding-kiko/avatar_service/pkg/log"

	// Third party
	"github.com/streadway/amqp"
)

var (
	avatarQueue = os.Getenv("AVATAR_QUEUE")
)

// Receives new avatar or user update from the users or groups service
func NewAvatarQueueConsumer(service Service, conn *amqp.Connection, logger log.Logger) {
	ch, err := conn.Channel()
	if err != nil {
		logger.Error("queue.go", "NewAvatarQueueConsumer", err.Error())
		panic(err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		avatarQueue, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		logger.Error("queue.go", "NewAvatarQueueConsumer", err.Error())
		panic(err.Error())
	}
	logger.Info("queue.go", "NewAvatarQueueConsumer", "avatarQueue declared")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logger.Error("queue.go", "NewAvatarQueueConsumer", err.Error())
		panic(err.Error())
	}

	logger.Info("queue.go", "NewAvatarQueueConsumer", "avatarQueue: start listening for messages")
	for d := range msgs {
		file := File{Data: d.Body, Filename: fmt.Sprintf("%s", d.Headers["filename"])}
		_ = service.ProcessAvatar(file)

	}
	fmt.Println("exits for loop for some reason")

}
