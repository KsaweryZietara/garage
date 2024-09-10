package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertConfirmationCode(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	confirmationCodeRepo := NewConfirmationCode(connection)

	newEmployee := internal.Employee{
		Name:     "John",
		Surname:  "Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Role:     "OWNER",
	}
	employee, err := employeeRepo.Insert(newEmployee)
	assert.NoError(t, err)

	newConfirmationCode := internal.ConfirmationCode{
		ID:         uuid.New().String(),
		EmployeeID: employee.ID,
	}
	confirmationCode, err := confirmationCodeRepo.Insert(newConfirmationCode)
	assert.NoError(t, err)
	assert.Equal(t, newConfirmationCode.ID, confirmationCode.ID)
	assert.Equal(t, newConfirmationCode.EmployeeID, confirmationCode.EmployeeID)
}
