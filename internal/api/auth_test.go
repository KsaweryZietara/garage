package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/stretchr/testify/assert"
)

func TestRegisterEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	employee := internal.RegisterDTO{
		Name:            "John",
		Surname:         "Doe",
		Email:           "john.doe@example.com",
		Password:        "Password123",
		ConfirmPassword: "Password123",
		Role:            "MECHANIC",
	}

	employeeJSON, _ := json.Marshal(employee)

	response := suite.CallAPI(http.MethodPost, "/api/business/register", employeeJSON)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}
