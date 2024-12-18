package storage

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/storage/postgres"

	_ "github.com/lib/pq"
)

type IStorage interface {
	Employees() Employees
	Garages() Garages
	Services() Services
	ConfirmationCodes() ConfirmationCodes
	Customers() Customers
	Appointments() Appointments
	Cars() Cars
}

type Employees interface {
	Insert(employee internal.Employee) (internal.Employee, error)
	GetByEmail(email string) (internal.Employee, error)
	Update(employee internal.Employee) error
	ListConfirmedByGarageID(garageID int) ([]internal.Employee, error)
	ListByGarageID(garageID int) ([]internal.Employee, error)
	GetConfirmedByID(ID int) (internal.Employee, error)
	GetByID(ID int) (internal.Employee, error)
	Delete(ID int) error
	UpdateProfilePicture(ID int, profilePicture []byte) error
}

type Garages interface {
	Insert(garage internal.Garage) (internal.Garage, error)
	GetByOwnerID(employeeID int) (internal.Garage, error)
	GetByID(ID int) (internal.Garage, error)
	List(page int, query string, latitude, longitude float64, sortBy string) ([]internal.Garage, error)
	Update(garage internal.Garage) error
	UpdateLogo(ID int, logo []byte) error
}

type Services interface {
	Insert(service internal.Service) (internal.Service, error)
	ListByGarageID(garageID int) ([]internal.Service, error)
	GetByID(ID int) (internal.Service, error)
	Delete(ID int) error
}

type ConfirmationCodes interface {
	Insert(code internal.ConfirmationCode) (internal.ConfirmationCode, error)
	GetByID(ID string) (internal.ConfirmationCode, error)
	DeleteByEmployeeID(ID int) error
}

type Customers interface {
	Insert(customer internal.Customer) (internal.Customer, error)
	GetByEmail(email string) (internal.Customer, error)
}

type Appointments interface {
	Insert(appointment internal.Appointment) (internal.Appointment, error)
	GetByTimeSlot(slot internal.TimeSlot, employeeID int) ([]internal.Appointment, error)
	GetByEmployeeID(employeeID int, date time.Time) ([]internal.Appointment, error)
	GetByGarageID(garageID int, date time.Time) ([]internal.Appointment, error)
	GetByCustomerID(customerID int) ([]internal.Appointment, error)
	GetByID(ID int) (internal.Appointment, error)
	Update(appointment internal.Appointment) error
	ListByGarageID(garageID int) ([]internal.Appointment, error)
	Delete(ID int) error
}

type Cars interface {
	ListMakes() ([]internal.Make, error)
	ListModels(makeID int) ([]internal.Model, error)
	GetByModelID(modelID int) (internal.Car, error)
}

type Storage struct {
	employees         Employees
	garages           Garages
	services          Services
	confirmationCodes ConfirmationCodes
	customers         Customers
	appointments      Appointments
	cars              Cars
}

func New(url string, log *slog.Logger) (Storage, error) {
	connection, err := postgres.WaitForDatabaseAccess(url, postgres.RetryCount, log)
	if err != nil {
		return Storage{}, err
	}

	err = postgres.RunMigrations(connection, postgres.Up, "../../../../")
	if err != nil {
		return Storage{}, err
	}

	return Storage{
		employees:         postgres.NewEmployee(connection),
		garages:           postgres.NewGarage(connection),
		services:          postgres.NewService(connection),
		confirmationCodes: postgres.NewConfirmationCode(connection),
		customers:         postgres.NewCustomer(connection),
		appointments:      postgres.NewAppointment(connection),
		cars:              postgres.NewCar(connection),
	}, nil
}

func NewForTests(url string, log *slog.Logger) (Storage, func() error, error) {
	connection, err := postgres.WaitForDatabaseAccess(url, postgres.RetryCount, log)
	if err != nil {
		return Storage{}, nil, err
	}

	err = postgres.RunMigrations(connection, postgres.Up, "../../../")
	if err != nil {
		return Storage{}, nil, err
	}

	cleanup := func() error {
		err = postgres.RunMigrations(connection, postgres.Down, "../../../")
		if err != nil {
			return fmt.Errorf("failed to clear DB tables: %w", err)
		}
		return nil
	}

	return Storage{
		employees:         postgres.NewEmployee(connection),
		garages:           postgres.NewGarage(connection),
		services:          postgres.NewService(connection),
		confirmationCodes: postgres.NewConfirmationCode(connection),
		customers:         postgres.NewCustomer(connection),
		appointments:      postgres.NewAppointment(connection),
		cars:              postgres.NewCar(connection),
	}, cleanup, nil
}

func (s Storage) Employees() Employees {
	return s.employees
}

func (s Storage) Garages() Garages {
	return s.garages
}

func (s Storage) Services() Services {
	return s.services
}

func (s Storage) ConfirmationCodes() ConfirmationCodes {
	return s.confirmationCodes
}

func (s Storage) Customers() Customers {
	return s.customers
}

func (s Storage) Appointments() Appointments {
	return s.appointments
}

func (s Storage) Cars() Cars {
	return s.cars
}
