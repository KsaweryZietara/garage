package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)
	serviceRepo := NewService(connection)

	newEmployee := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Role:     "OWNER",
	}
	employee, err := employeeRepo.Insert(newEmployee)
	assert.NoError(t, err)

	newGarage := internal.Garage{
		Name:        "Test Garage",
		City:        "Test City",
		Street:      "Test Street",
		Number:      "123",
		PostalCode:  "12345",
		PhoneNumber: "1234567890",
		OwnerID:     employee.ID,
	}
	garage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	newService1 := internal.Service{
		Name:     "Test Service",
		Time:     60,
		Price:    100.0,
		GarageID: garage.ID,
	}
	service1, err := serviceRepo.Insert(newService1)
	assert.NoError(t, err)
	newService2 := internal.Service{
		Name:     "Test Service 2",
		Time:     60,
		Price:    100.0,
		GarageID: garage.ID,
	}
	_, err = serviceRepo.Insert(newService2)
	assert.NoError(t, err)

	services, err := serviceRepo.ListByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(services))

	service, err := serviceRepo.GetByID(service1.ID)
	assert.NoError(t, err)
	assert.Equal(t, service1.ID, service.ID)
	assert.Equal(t, service1.Name, service.Name)
	assert.Equal(t, service1.Time, service.Time)
	assert.Equal(t, service1.Price, service.Price)
}
