package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/stretchr/testify/assert"
)

func TestGetServicesEndpoint(t *testing.T) {
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

	_, err = suite.api.storage.Services().Insert(internal.Service{
		Name:     "name",
		Time:     30,
		Price:    100,
		GarageID: garage.ID,
	})
	assert.NoError(t, err)

	response := suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/garages/%v/services", garage.ID), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var serviceDTOs []internal.ServiceDTO
	suite.ParseResponse(t, response, &serviceDTOs)

	assert.Equal(t, 1, len(serviceDTOs))
	assert.Equal(t, "name", serviceDTOs[0].Name)
	assert.Equal(t, 30, serviceDTOs[0].Time)
	assert.Equal(t, 100, serviceDTOs[0].Price)
}

func TestGetServiceEndpoint(t *testing.T) {
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

	service, err := suite.api.storage.Services().Insert(internal.Service{
		Name:     "name",
		Time:     30,
		Price:    100,
		GarageID: garage.ID,
	})
	assert.NoError(t, err)

	response := suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/services/%v", service.ID), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var serviceDTO internal.ServiceDTO
	suite.ParseResponse(t, response, &serviceDTO)

	assert.Equal(t, service.ID, serviceDTO.ID)
	assert.Equal(t, service.Name, serviceDTO.Name)
	assert.Equal(t, service.Time, serviceDTO.Time)
	assert.Equal(t, service.Price, serviceDTO.Price)
}
