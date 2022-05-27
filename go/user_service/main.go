package main

import (
	// std lib
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// Internal
	"github.com/coding-kiko/user_service/pkg/log"
	"github.com/coding-kiko/user_service/pkg/user"

	// Third Party
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

var (
	// Api
	ApiPort = os.Getenv("API_PORT")
	// Postgres
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPwd  = os.Getenv("POSTGRES_PWD")
	PostgresHost = os.Getenv("POSTGRES_HOST")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	PostgresDB   = os.Getenv("POSTGRES_DB")
	// Rabbit
	RabbitMQUser = os.Getenv("RABBITMQ_USER")
	RabbitMQPwd  = os.Getenv("RABBITMQ_PWD")
	RabbitMQHost = os.Getenv("RABBITMQ_HOST")
	RabbitMQPort = os.Getenv("RABBITMQ_PORT")
)

func main() {
	logger := log.NewLogger()

	// create postgres connection
	postgresConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", PostgresUser, PostgresPwd, PostgresHost, PostgresPort, PostgresDB)
	postgresDb, err := sql.Open("postgres", postgresConnString)
	if err != nil {
		logger.Error("main.go", "main", err.Error())
		panic(err)
	}
	defer postgresDb.Close()

	err = postgresDb.Ping()
	if err != nil {
		logger.Error("main.go", "main", err.Error())
		panic(err)
	}

	// rabbitmq connection
	rabbitConnString := fmt.Sprintf("amqp://%s:%s@%s:%s/", RabbitMQUser, RabbitMQPwd, RabbitMQHost, RabbitMQPort)
	rabbitConn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		logger.Error("main.go", "main", err.Error())
		panic(err)
	}
	defer rabbitConn.Close()

	// initialize repository layer
	repository := user.NewRepository(postgresDb, logger)

	// initialize rabbit producer layer
	rabbitProducer := user.NewRabbitProducer(rabbitConn, logger)

	// initialize service layer
	service := user.NewService(repository, rabbitProducer, logger)

	// initialize rabbit consumer layer
	rabbitConsumer := user.NewRabbitConsumer(service, rabbitConn, logger)

	// listen concurrently to the queue sending users information from authentication service
	go rabbitConsumer.UsersQueue()

	// initialize handlers layer
	handlers := user.NewHandler(service, logger)

	// start mux and listening
	router := user.NewRouter(handlers, logger)
	addr := fmt.Sprintf("0.0.0.0:%s", ApiPort)
	logger.Info("main.go", "main", "Started listening on "+addr)
	err = http.ListenAndServe(addr, router)

	logger.Error("main.go", "main", err.Error())
	os.Exit(1)
}
