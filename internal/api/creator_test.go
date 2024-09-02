package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatorEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	token := suite.CreateEmployee(t,
		internal.Employee{
			Name:     "John",
			Surname:  "Doe",
			Email:    "john.doe@example.com",
			Password: "Password123",
			Role:     internal.Owner,
		})

	creator := internal.CreatorDTO{
		Name:        "John's Garage",
		City:        "San Francisco",
		Street:      "Market Street",
		Number:      "123",
		PostalCode:  "94-103",
		PhoneNumber: "123456789",
		Services: []internal.ServiceDTO{
			{
				Name:  "Oil Change",
				Time:  30,
				Price: 50,
			},
			{
				Name:  "Tire Rotation",
				Time:  15,
				Price: 25,
			},
		},
		EmployeeEmails: []string{
			"john@example.com",
			"jane@example.com",
		},
	}

	creatorJSON, err := json.Marshal(creator)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/business/creator", creatorJSON, token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	response = suite.CallAPI(http.MethodPost, "/api/business/creator", creatorJSON, nil)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}
