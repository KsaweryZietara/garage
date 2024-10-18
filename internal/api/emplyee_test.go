package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEmployeesEndpoint(t *testing.T) {
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

	employee, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "name",
			Surname:   "surname",
			Email:     "email2",
			Password:  "password",
			Role:      internal.MechanicRole,
			GarageID:  &garage.ID,
			Confirmed: true,
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

	token, err := suite.api.auth.CreateToken("email", internal.OwnerRole)
	require.NoError(t, err)

	response = suite.CallAPI(http.MethodGet, "/api/employees", []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	suite.ParseResponse(t, response, &employeeDTOs)

	assert.Equal(t, 1, len(employeeDTOs))
	assert.Equal(t, employee.ID, employeeDTOs[0].ID)
	assert.Equal(t, employee.Name, employeeDTOs[0].Name)
	assert.Equal(t, employee.Surname, employeeDTOs[0].Surname)
}

func TestGetEmployeeEndpoint(t *testing.T) {
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

	employee, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "name",
			Surname:   "surname",
			Email:     "email2",
			Password:  "password",
			Role:      internal.MechanicRole,
			GarageID:  &garage.ID,
			Confirmed: true,
		})
	assert.NoError(t, err)

	response := suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/employees/%v", owner.ID), []byte{}, nil)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	response = suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/employees/%v", employee.ID), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var employeeDTO internal.EmployeeDTO
	suite.ParseResponse(t, response, &employeeDTO)

	assert.Equal(t, employee.ID, employeeDTO.ID)
	assert.Equal(t, employee.Name, employeeDTO.Name)
	assert.Equal(t, employee.Surname, employeeDTO.Surname)
}
