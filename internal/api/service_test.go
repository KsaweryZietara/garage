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
