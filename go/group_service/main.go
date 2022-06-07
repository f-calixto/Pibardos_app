package main

import (
	// std lib
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// Internal
	"github.com/coding-kiko/group_service/pkg/group"
	"github.com/coding-kiko/group_service/pkg/log"

	// Third Party
	_ "github.com/lib/pq"
)

var (
	ApiPort           = os.Getenv("API_PORT")
	PostgresDB        = os.Getenv("POSTGRES_DB")
	PostgresContainer = os.Getenv("POSTGRES_CONTAINER")
	PostgresPort      = os.Getenv("POSTGRES_PORT")
	PostgresPwd       = os.Getenv("POSTGRES_PWD")
)

func main() {
	logger := log.NewLogger()

	// create postgres connection
	postgresConnString := fmt.Sprintf("postgres://postgres:%s@%s:%s/%s?sslmode=disable", PostgresPwd, PostgresContainer, PostgresPort, PostgresDB)
	postgresDb, err := sql.Open("postgres", postgresConnString)
	if err != nil {
		panic(err)
	}
	defer postgresDb.Close()

	// initialize repository layer
	repository := group.NewRepository(postgresDb, logger)

	// initialize service layer
	service := group.NewService(repository, logger)

	// initialize handlers layer
	handlers := group.NewHandler(service, logger)

	// start mux and listening
	router := group.NewRouter(handlers, logger)
	addr := fmt.Sprintf("0.0.0.0:%s", ApiPort)
	logger.Info("main.go", "main", "Started listening on "+addr)
	err = http.ListenAndServe(addr, router)

	logger.Error("main.go", "main", err.Error())
	os.Exit(1)
}
