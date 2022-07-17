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
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func main() {
	logger := log.NewLogger()

	viper.SetConfigFile("prod.env")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("main.go", "main", err.Error())
		panic(err)
	}

	var (
		// Api
		ApiPort = viper.Get("API_PORT").(string)
		// Rabbit
		RabbitMQUser = viper.Get("RABBITMQ_USER").(string)
		RabbitMQPwd  = viper.Get("RABBITMQ_PWD").(string)
		RabbitMQHost = viper.Get("RABBITMQ_HOST").(string)
		RabbitMQPort = viper.Get("RABBITMQ_PORT").(string)
		// files folder absolute path
		StaticPath = viper.Get("STATIC_PATH").(string)
	)

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
