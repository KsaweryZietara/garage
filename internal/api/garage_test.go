package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOwnerGetGarageEndpoint(t *testing.T) {
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
		Name:           "John's Garage",
		City:           "San Francisco",
		Street:         "Market Street",
		Number:         "123",
		PostalCode:     "94-103",
		PhoneNumber:    "123456789",
		Services:       []internal.ServiceDTO{},
		EmployeeEmails: []string{},
	}
	creatorJSON, err := json.Marshal(creator)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/business/creator", creatorJSON, token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	response = suite.CallAPI(http.MethodGet, "/api/garages", []byte{}, token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestMechanicGetGarageEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	owner, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:     "name",
			Surname:  "surname",
			Email:    "email",
			Password: "password",
			Role:     internal.Owner,
			GarageID: nil,
		})
	assert.NoError(t, err)

	garage, err := suite.api.storage.Garages().Insert(
		internal.Garage{
			Name:        "name",
			City:        "city",
			Street:      "street",
			Number:      "number",
			PostalCode:  "postalCode",
			PhoneNumber: "phoneNumber",
			OwnerID:     owner.ID,
		})
	assert.NoError(t, err)

	token := suite.CreateEmployee(t,
		internal.Employee{
			Name:     "John",
			Surname:  "Doe",
			Email:    "john.doe@example.com",
			Password: "Password123",
			Role:     internal.Mechanic,
			GarageID: &garage.ID,
		})

	response := suite.CallAPI(http.MethodGet, "/api/garages", []byte{}, token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
