package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/KsaweryZietara/garage/internal/api"
	"github.com/KsaweryZietara/garage/internal/auth"
	"github.com/KsaweryZietara/garage/internal/mail"
	"github.com/KsaweryZietara/garage/internal/storage"
	"github.com/KsaweryZietara/garage/internal/storage/postgres"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server   api.Config      `env:", prefix=SERVER_"`
	Postgres postgres.Config `env:", prefix=POSTGRES_"`
	Mail     mail.Config     `env:", prefix=MAIL_"`
	AuthKey  string          `env:"AUTH_KEY"`
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.Background()
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	storage, err := storage.New(cfg.Postgres.ConnectionURL(), log)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	auth := auth.New(cfg.AuthKey)

	mail := mail.New(cfg.Mail)

	api := api.New(cfg.Server, log, storage, auth, mail)
	api.Start()
}
