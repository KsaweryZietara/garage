package api

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAppointmentEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	token := suite.CreateCustomer(t,
		internal.Customer{
			Email:    "john.doe@example.com",
			Password: "Password123",
		})

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

	mechanic, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:     "name",
			Surname:  "surname",
			Email:    "email2",
			Password: "password",
			Role:     internal.MechanicRole,
			GarageID: &garage.ID,
		})
	assert.NoError(t, err)

	service, err := suite.api.storage.Services().Insert(
		internal.Service{
			Name:     "name",
			Time:     30,
			Price:    10,
			GarageID: garage.ID,
		})
	assert.NoError(t, err)

	appointment := internal.CreateAppointmentDTO{
		StartTime:  time.Now(),
		EndTime:    time.Now().Add(3 * time.Hour),
		ServiceID:  service.ID,
		EmployeeID: mechanic.ID,
	}
	appointmentJSON, err := json.Marshal(appointment)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/appointments", appointmentJSON, token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}
