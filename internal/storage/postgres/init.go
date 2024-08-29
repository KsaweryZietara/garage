package postgres

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"runtime"
	"slices"
	"strings"
	"time"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
)

const (
	connectionURLFormat = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	RetryCount          = 10
	timeout             = 300 * time.Millisecond
)

type migrationOrder int

const (
	Up migrationOrder = iota
	Down
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func (cfg *Config) ConnectionURL() string {
	return fmt.Sprintf(connectionURLFormat, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
}

func WaitForDatabaseAccess(connString string, retryCount int, log *slog.Logger) (*dbr.Connection, error) {
	var connection *dbr.Connection
	var err error

	for ; retryCount > 0; retryCount-- {
		connection, err = dbr.Open("postgres", connString, nil)
		if err != nil {
			return nil, fmt.Errorf("invalid connection string: %w", err)
		}

		err = connection.Ping()
		if err == nil {
			return connection, nil
		}
		log.Warn("database connection failed: %s", err)

		err = connection.Close()
		if err != nil {
			log.Info("failed to close database ...")
		}

		log.Info("failed to access database, retrying...")
		time.Sleep(timeout)
	}

	return nil, fmt.Errorf("timeout waiting for database access")
}

func RunMigrations(connection *dbr.Connection, order migrationOrder) error {
	_, currentPath, _, _ := runtime.Caller(0)
	migrationsPath := fmt.Sprintf("%s/resources/migrations/", path.Join(path.Dir(currentPath), "../../../"))

	if order != Up && order != Down {
		return fmt.Errorf("unknown migration order")
	}

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("while reading migration data: %w in directory :%s", err, migrationsPath)
	}

	suffix := ""
	if order == Down {
		suffix = "down.sql"
		slices.Reverse(files)
	}

	if order == Up {
		suffix = "up.sql"
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), suffix) {
			content, err := os.ReadFile(migrationsPath + file.Name())
			if err != nil {
				return fmt.Errorf("while reading migration files: %w file: %s", err, file.Name())
			}
			if _, err = connection.Exec(string(content)); err != nil {
				return fmt.Errorf("while applying migration files: %w file: %s", err, file.Name())
			}
		}
	}

	return nil
}
