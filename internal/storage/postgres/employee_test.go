package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestEmployee(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	garageRepo := NewGarage(connection)

	newEmployee := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "test@test.com",
		Password: "password123",
		Role:     internal.OwnerRole,
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

	newEmployee2 := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "test2@test.com",
		Password: "password123",
		Role:     internal.MechanicRole,
		GarageID: &garage.ID,
	}
	_, err = employeeRepo.Insert(newEmployee2)
	assert.NoError(t, err)

	newEmployee3 := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "test3@test.com",
		Password: "password123",
		Role:     internal.MechanicRole,
		GarageID: &garage.ID,
	}
	_, err = employeeRepo.Insert(newEmployee3)
	assert.NoError(t, err)

	retrievedEmployee, err := employeeRepo.GetByEmail(newEmployee.Email)
	assert.NoError(t, err)
	assert.Equal(t, newEmployee.Name, retrievedEmployee.Name)
	assert.Equal(t, newEmployee.Surname, retrievedEmployee.Surname)
	assert.Equal(t, newEmployee.Email, retrievedEmployee.Email)
	assert.Equal(t, newEmployee.Password, retrievedEmployee.Password)
	assert.Equal(t, newEmployee.Role, retrievedEmployee.Role)

	employee.Name = "newName"
	employee.Surname = "newSurname"
	employee.Password = "newPassword"
	err = employeeRepo.Update(employee)
	assert.NoError(t, err)

	updatedEmployee, err := employeeRepo.GetByEmail(newEmployee.Email)
	assert.NoError(t, err)
	assert.Equal(t, employee.Name, updatedEmployee.Name)
	assert.Equal(t, employee.Surname, updatedEmployee.Surname)
	assert.Equal(t, employee.Password, updatedEmployee.Password)

	employees, err := employeeRepo.ListByGarageID(garage.ID)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(employees))
}
