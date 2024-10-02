package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestConfirmationCode(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	employeeRepo := NewEmployee(connection)
	confirmationCodeRepo := NewConfirmationCode(connection)

	newEmployee := internal.Employee{
		Name:      "John",
		Surname:   "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		Role:      "OWNER",
		Confirmed: false,
	}
	employee, err := employeeRepo.Insert(newEmployee)
	assert.NoError(t, err)

	newConfirmationCode := internal.ConfirmationCode{
		ID:         uuid.New().String(),
		EmployeeID: employee.ID,
	}
	createdConfirmationCode, err := confirmationCodeRepo.Insert(newConfirmationCode)
	assert.NoError(t, err)
	assert.Equal(t, newConfirmationCode.ID, createdConfirmationCode.ID)
	assert.Equal(t, newConfirmationCode.EmployeeID, createdConfirmationCode.EmployeeID)

	confirmationCode, err := confirmationCodeRepo.GetByID(newConfirmationCode.ID)
	assert.NoError(t, err)
	assert.Equal(t, newConfirmationCode.ID, confirmationCode.ID)
	assert.Equal(t, newConfirmationCode.EmployeeID, confirmationCode.EmployeeID)

	err = confirmationCodeRepo.DeleteByID(newConfirmationCode.ID)
	assert.NoError(t, err)

	_, err = confirmationCodeRepo.GetByID(newConfirmationCode.ID)
	assert.EqualError(t, err, "dbr: not found")
}
