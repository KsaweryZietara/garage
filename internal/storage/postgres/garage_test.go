package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestGarage(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)

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
	createdGarage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	garage, err := garageRepo.GetByOwnerID(employee.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdGarage, garage)

	garage, err = garageRepo.GetByID(createdGarage.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdGarage, garage)
}
