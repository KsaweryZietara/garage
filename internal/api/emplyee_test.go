package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/stretchr/testify/assert"
)

func TestGetEmployeesEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	owner, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:     "name",
			Surname:  "surname",
			Email:    "email",
			Password: "password",
			Role:     internal.OwnerRole,
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

	employee, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:     "name",
			Surname:  "surname",
			Email:    "email2",
			Password: "password",
			Role:     internal.MechanicRole,
			GarageID: &garage.ID,
		})
	assert.NoError(t, err)

	response := suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/garages/%v/employees", garage.ID), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var employeeDTOs []internal.EmployeeDTO
	suite.ParseResponse(t, response, &employeeDTOs)

	assert.Equal(t, 1, len(employeeDTOs))
	assert.Equal(t, employee.ID, employeeDTOs[0].ID)
	assert.Equal(t, employee.Name, employeeDTOs[0].Name)
	assert.Equal(t, employee.Surname, employeeDTOs[0].Surname)
}
