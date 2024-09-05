package postgres

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var connection *dbr.Connection

func TestMain(m *testing.M) {
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test-storage"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
	)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	connString, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	connection, err = WaitForDatabaseAccess(connString, RetryCount, log)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	exitValue := m.Run()
	os.Exit(exitValue)
}

func NewSuite(t *testing.T) func() {
	err := RunMigrations(connection, Up, "../../../")
	require.NoError(t, err)
	return func() {
		err = RunMigrations(connection, Down, "../../../")
		require.NoError(t, err)
	}
}
