package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestInsertGarage(t *testing.T) {
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
	garage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)
	newGarage.ID = garage.ID

	assert.Equal(t, newGarage, garage)
}
