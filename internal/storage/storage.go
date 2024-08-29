package storage

import (
	"fmt"
	"log/slog"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/storage/postgres"
	_ "github.com/lib/pq"
)

type IStorage interface {
	Employees() Employees
}

type Employees interface {
	Insert(employee internal.Employee) error
	GetByEmail(email string) (internal.Employee, error)
}

type Storage struct {
	employees Employees
}

func New(url string, log *slog.Logger) (Storage, error) {
	connection, err := postgres.WaitForDatabaseAccess(url, postgres.RetryCount, log)
	if err != nil {
		return Storage{}, err
	}

	err = postgres.RunMigrations(connection, postgres.Up)
	if err != nil {
		return Storage{}, err
	}

	return Storage{
		employees: postgres.NewEmployee(connection),
	}, nil
}

func NewForTests(url string, log *slog.Logger) (Storage, func() error, error) {
	connection, err := postgres.WaitForDatabaseAccess(url, postgres.RetryCount, log)
	if err != nil {
		return Storage{}, nil, err
	}

	err = postgres.RunMigrations(connection, postgres.Up)
	if err != nil {
		return Storage{}, nil, err
	}

	cleanup := func() error {
		err = postgres.RunMigrations(connection, postgres.Down)
		if err != nil {
			return fmt.Errorf("failed to clear DB tables: %w", err)
		}
		return nil
	}

	return Storage{
		employees: postgres.NewEmployee(connection),
	}, cleanup, nil
}

func (s Storage) Employees() Employees {
	return s.employees
}
