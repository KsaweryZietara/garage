package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)

	newEmployee := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Role:     "OWNER",
	}

	err := employeeRepo.Insert(newEmployee)
	assert.NoError(t, err)

	retrievedEmployee, err := employeeRepo.GetByEmail(newEmployee.Email)
	assert.NoError(t, err)

	assert.Equal(t, newEmployee.Name, retrievedEmployee.Name)
	assert.Equal(t, newEmployee.Surname, retrievedEmployee.Surname)
	assert.Equal(t, newEmployee.Email, retrievedEmployee.Email)
	assert.Equal(t, newEmployee.Password, retrievedEmployee.Password)
	assert.Equal(t, newEmployee.Role, retrievedEmployee.Role)
}
