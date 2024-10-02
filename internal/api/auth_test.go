package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOwnerRegisterAndLoginEndpoints(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	employee := internal.CreateEmployeeDTO{
		Name:            "John",
		Surname:         "Doe",
		Email:           "john.doe@example.com",
		Password:        "Password123",
		ConfirmPassword: "Password123",
	}
	employeeJSON, err := json.Marshal(employee)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/employees/register", employeeJSON, nil)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	loginDTO := internal.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "Password123",
	}
	loginJSON, err := json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/employees/login", loginJSON, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	loginDTO = internal.LoginDTO{
		Email:    "wrong.mail@example.com",
		Password: "Password123",
	}
	loginJSON, err = json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/employees/login", loginJSON, nil)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	loginDTO = internal.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "WrongPassword",
	}
	loginJSON, err = json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/employees/login", loginJSON, nil)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestMechanicRegisterAndLoginEndpoints(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	employee, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "",
			Surname:   "",
			Email:     "john.doe@example.com",
			Password:  "",
			Role:      internal.MechanicRole,
			GarageID:  nil,
			Confirmed: false,
		})
	assert.NoError(t, err)

	codeID := uuid.New().String()
	_, err = suite.api.storage.ConfirmationCodes().Insert(
		internal.ConfirmationCode{
			ID:         codeID,
			EmployeeID: employee.ID,
		})
	assert.NoError(t, err)

	employeeDTO := internal.CreateEmployeeDTO{
		Name:            "John",
		Surname:         "Doe",
		Email:           "john.doe@example.com",
		Password:        "Password123",
		ConfirmPassword: "Password123",
	}
	employeeJSON, err := json.Marshal(employeeDTO)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/employees/register/"+codeID, employeeJSON, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	loginDTO := internal.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "Password123",
	}
	loginJSON, err := json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/employees/login", loginJSON, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCustomerRegisterAndLoginEndpoints(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	customer := internal.CreateCustomerDTO{
		Email:           "john.doe@example.com",
		Password:        "Password123",
		ConfirmPassword: "Password123",
	}
	customerJSON, err := json.Marshal(customer)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/customers/register", customerJSON, nil)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	loginDTO := internal.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "Password123",
	}
	loginJSON, err := json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/customers/login", loginJSON, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	loginDTO = internal.LoginDTO{
		Email:    "wrong.mail@example.com",
		Password: "Password123",
	}
	loginJSON, err = json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/customers/login", loginJSON, nil)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	loginDTO = internal.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "WrongPassword",
	}
	loginJSON, err = json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/customers/login", loginJSON, nil)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}
