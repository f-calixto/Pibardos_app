// package main

import (
	// std lib
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// Internal
	"github.com/coding-kiko/debts_service/pkg/debts"
	"github.com/coding-kiko/debts_service/pkg/log"

	// Third Party
	_ "github.com/lib/pq"
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

	// initialize repository layer
	repository := debts.NewRepository(postgresDb, logger)

	// initialize service layer
	service := debts.NewService(repository, logger)

	// initialize handlers layer
	handlers := debts.NewHandler(service, logger)

	// start mux and listening
	router := debts.NewRouter(handlers, logger)
	addr := fmt.Sprintf("0.0.0.0:%s", ApiPort)
	logger.Info("main.go", "main", "Started listening on "+addr)
	err = http.ListenAndServe(addr, router)

	logger.Error("main.go", "main", err.Error())
	os.Exit(1)
}
