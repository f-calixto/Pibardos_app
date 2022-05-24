package main

import (
	// std lib
	"fmt"
	"net/http"
	"os"

	// internal
	"github.com/coding-kiko/avatar_service/pkg/avatar"
	"github.com/coding-kiko/avatar_service/pkg/log"

	// third party
	"github.com/streadway/amqp"
)

var (
	// Api
	ApiPort = os.Getenv("API_PORT")
	// Rabbit
	RabbitMQUser = os.Getenv("RABBITMQ_USER")
	RabbitMQPwd  = os.Getenv("RABBITMQ_PWD")
	RabbitMQHost = os.Getenv("RABBITMQ_HOST")
	RabbitMQPort = os.Getenv("RABBITMQ_PORT")
	// files folder absolute path
	StaticPath = os.Getenv("STATIC_PATH")
)

func main() {
	logger := log.NewLogger()

	// initialize avatar Storage layer
	storage := avatar.NewAvatarStorage(StaticPath, logger)

	// initialize service layer
	service := avatar.NewService(storage, logger)

	// rabbitmq connection
	rabbitConnString := fmt.Sprintf("amqp://%s:%s@%s:%s/", RabbitMQUser, RabbitMQPwd, RabbitMQHost, RabbitMQPort)
	rabbitConn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		logger.Error("main.go", "main", err.Error())
		panic(err)
	}
	defer rabbitConn.Close()

	go avatar.NewAvatarQueueConsumer(service, rabbitConn, logger)

	fs := http.FileServer(http.Dir(StaticPath))
	http.Handle("/", fs)
	addr := fmt.Sprintf("0.0.0.0:%s", ApiPort)
	logger.Info("main.go", "main", "Listening on "+addr)
	err = http.ListenAndServe(addr, nil)
	logger.Error("main.go", "main", err.Error())
	os.Exit(1)
}
