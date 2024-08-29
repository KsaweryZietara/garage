package main

import (
	"log/slog"
	"os"

	"github.com/KsaweryZietara/garage/internal/api"
	"github.com/KsaweryZietara/garage/internal/storage"
	"github.com/KsaweryZietara/garage/internal/storage/postgres"
)

func main() {
	server := api.Config{
		Port: "8080",
	}
	database := postgres.Config{
		User:     "postgres",
		Password: "password",
		Host:     "localhost",
		Port:     "5432",
		Name:     "garage",
		SSLMode:  "disable",
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	storage, err := storage.New(database.ConnectionURL(), log)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	api := api.New(server, log, storage)

	api.Start()
}
