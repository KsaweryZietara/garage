package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterAndLoginEndpoints(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	employee := internal.RegisterDTO{
		Name:            "John",
		Surname:         "Doe",
		Email:           "john.doe@example.com",
		Password:        "Password123",
		ConfirmPassword: "Password123",
		Role:            "OWNER",
	}
	employeeJSON, err := json.Marshal(employee)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/business/register", employeeJSON)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	loginDTO := internal.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "Password123",
	}
	loginJSON, err := json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/business/login", loginJSON)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	loginDTO = internal.LoginDTO{
		Email:    "wrong.mail@example.com",
		Password: "Password123",
	}
	loginJSON, err = json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/business/login", loginJSON)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	loginDTO = internal.LoginDTO{
		Email:    "john.doe@example.com",
		Password: "WrongPassword",
	}
	loginJSON, err = json.Marshal(loginDTO)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/business/login", loginJSON)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}
