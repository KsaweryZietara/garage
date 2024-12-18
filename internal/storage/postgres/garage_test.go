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
		Name:      "John",
		Surname:   "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Role:      "OWNER",
		Confirmed: true,
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
		Latitude:    10,
		Longitude:   10,
	}
	createdGarage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	garage, err := garageRepo.GetByOwnerID(employee.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdGarage, garage)

	garage, err = garageRepo.GetByID(createdGarage.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdGarage, garage)

	updatedGarage := internal.Garage{
		ID:          createdGarage.ID,
		Name:        "new name",
		City:        "new city",
		Street:      "new street",
		Number:      "new number",
		PostalCode:  "99-999",
		PhoneNumber: "999999999",
		OwnerID:     employee.ID,
		Latitude:    20,
		Longitude:   20,
	}
	err = garageRepo.Update(updatedGarage)
	assert.NoError(t, err)

	garage, err = garageRepo.GetByID(createdGarage.ID)
	assert.NoError(t, err)
	assert.Equal(t, garage, updatedGarage)

	logo := []byte("logo")
	err = garageRepo.UpdateLogo(createdGarage.ID, logo)
	assert.NoError(t, err)

	garage, err = garageRepo.GetByID(createdGarage.ID)
	assert.NoError(t, err)
	assert.Equal(t, logo, garage.Logo)
}

func TestListGarage(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)
	serviceRepo := NewService(connection)

	newEmployee := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Role:      "OWNER",
		Confirmed: true,
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
		Latitude:    10,
		Longitude:   10,
	}
	garage, err := garageRepo.Insert(newGarage)
	assert.NoError(t, err)

	newService := internal.Service{
		Name:     "Test Service",
		Time:     10,
		Price:    10,
		GarageID: garage.ID,
	}
	_, err = serviceRepo.Insert(newService)
	assert.NoError(t, err)

	testCases := []struct {
		query         string
		expectedCount int
	}{
		{"Test Garage", 1},
		{"Test G", 1},
		{"age", 1},
		{"test garage", 1},
		{"invalid garage", 0},
		{"Test Service", 1},
		{"Test Se", 1},
		{"vice", 1},
		{"test service", 1},
		{"invalid service", 0},
	}
	for _, tc := range testCases {
		garages, err := garageRepo.List(1, tc.query, 0, 0, "rating")
		assert.NoError(t, err)
		assert.Equal(t, tc.expectedCount, len(garages))
	}
}
