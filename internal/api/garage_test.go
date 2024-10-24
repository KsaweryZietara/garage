package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGarageEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	token := suite.CreateEmployee(t,
		internal.Employee{
			Name:      "John",
			Surname:   "Doe",
			Email:     "john.doe@example.com",
			Password:  "Password123",
			Role:      internal.OwnerRole,
			Confirmed: true,
		})

	garage := internal.CreateGarageDTO{
		Name:        "John's Garage",
		City:        "San Francisco",
		Street:      "Market Street",
		Number:      "123",
		PostalCode:  "94-103",
		PhoneNumber: "123456789",
		Latitude:    10,
		Longitude:   10,
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
	garageJSON, err := json.Marshal(garage)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/garages", garageJSON, token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	garage.PhoneNumber = "987654321"
	garageJSON, err = json.Marshal(garage)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPost, "/api/garages", garageJSON, token)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	response = suite.CallAPI(http.MethodPost, "/api/garages", garageJSON, nil)
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	updatedGarage := internal.CreateGarageDTO{
		Name:        "new name",
		City:        "new city",
		Street:      "new street",
		Number:      "111",
		PostalCode:  "99-999",
		PhoneNumber: "999999999",
		Latitude:    20,
		Longitude:   20,
	}
	updatedGarageJSON, err := json.Marshal(updatedGarage)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodPut, "/api/garages", updatedGarageJSON, token)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	response = suite.CallAPI(http.MethodGet, "/api/employees/garages", []byte{}, token)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var garageDTO internal.GarageDTO
	suite.ParseResponse(t, response, &garageDTO)
	assert.Equal(t, "new name", garageDTO.Name)

	logo := internal.LogoDTO{Base64Logo: "logo"}
	logoJSON, err := json.Marshal(logo)
	require.NoError(t, err)
	response = suite.CallAPI(http.MethodPost, "/api/garages/logo", logoJSON, token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestOwnerGetGarageEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	token := suite.CreateEmployee(t,
		internal.Employee{
			Name:      "John",
			Surname:   "Doe",
			Email:     "john.doe@example.com",
			Password:  "Password123",
			Role:      internal.OwnerRole,
			Confirmed: true,
		})

	garage := internal.CreateGarageDTO{
		Name:           "John's Garage",
		City:           "San Francisco",
		Street:         "Market Street",
		Number:         "123",
		PostalCode:     "94-103",
		PhoneNumber:    "123456789",
		Latitude:       10,
		Longitude:      10,
		Services:       []internal.ServiceDTO{},
		EmployeeEmails: []string{},
	}
	garageJSON, err := json.Marshal(garage)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/garages", garageJSON, token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	response = suite.CallAPI(http.MethodGet, "/api/employees/garages", []byte{}, token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestMechanicGetGarageEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	owner, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "name",
			Surname:   "surname",
			Email:     "email",
			Password:  "password",
			Role:      internal.OwnerRole,
			GarageID:  nil,
			Confirmed: true,
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
			Latitude:    10,
			Longitude:   10,
		})
	assert.NoError(t, err)

	token := suite.CreateEmployee(t,
		internal.Employee{
			Name:      "John",
			Surname:   "Doe",
			Email:     "john.doe@example.com",
			Password:  "Password123",
			Role:      internal.MechanicRole,
			GarageID:  &garage.ID,
			Confirmed: true,
		})

	response := suite.CallAPI(http.MethodGet, "/api/employees/garages", []byte{}, token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestGetGaragesEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	owner, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "name",
			Surname:   "surname",
			Email:     "email",
			Password:  "password",
			Role:      internal.OwnerRole,
			GarageID:  nil,
			Confirmed: true,
		})
	assert.NoError(t, err)

	_, err = suite.api.storage.Garages().Insert(
		internal.Garage{
			Name:        "name",
			City:        "city",
			Street:      "street",
			Number:      "number",
			PostalCode:  "postalCode",
			PhoneNumber: "phoneNumber",
			OwnerID:     owner.ID,
			Latitude:    10,
			Longitude:   10,
		})
	assert.NoError(t, err)

	response := suite.CallAPI(http.MethodGet, "/api/garages?query=name&page=1", []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var garageDTOs []internal.GarageDTO
	suite.ParseResponse(t, response, &garageDTOs)

	assert.Equal(t, 1, len(garageDTOs))
	assert.Equal(t, "name", garageDTOs[0].Name)
}

func TestGetGarageEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	owner, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "name",
			Surname:   "surname",
			Email:     "email",
			Password:  "password",
			Role:      internal.OwnerRole,
			GarageID:  nil,
			Confirmed: true,
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
			Latitude:    10,
			Longitude:   10,
		})
	assert.NoError(t, err)

	response := suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/garages/%v", garage.ID), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var garageDTO internal.GarageDTO
	suite.ParseResponse(t, response, &garageDTO)

	assert.Equal(t, "name", garageDTO.Name)
}
