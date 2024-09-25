package api

import (
	"encoding/json"
	"fmt"
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
			Time:     2,
			Price:    10,
			GarageID: garage.ID,
		})
	assert.NoError(t, err)

	appointment := internal.CreateAppointmentDTO{
		StartTime:  time.Date(2030, 9, 24, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2030, 9, 24, 13, 0, 0, 0, time.UTC),
		ServiceID:  service.ID,
		EmployeeID: mechanic.ID,
	}
	appointmentJSON, err := json.Marshal(appointment)
	require.NoError(t, err)

	response := suite.CallAPI(http.MethodPost, "/api/appointments", appointmentJSON, token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	response = suite.CallAPI(
		http.MethodGet,
		fmt.Sprintf("/api/appointments/availableSlots?serviceId=%v&employeeId=%v&date=2024-09-24", service.ID, mechanic.ID),
		[]byte{},
		nil,
	)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCreateTimeSlots(t *testing.T) {
	t.Run("working days", func(t *testing.T) {
		date := time.Date(2024, 9, 23, 0, 0, 0, 0, time.UTC)
		serviceDuration := 10

		timeSlots := createTimeSlots(date, serviceDuration)

		require.Len(t, timeSlots, 9)

		assert.Equal(t, 23, timeSlots[0].StartTime.Day())
		assert.Equal(t, 8, timeSlots[0].StartTime.Hour())
		assert.Equal(t, 24, timeSlots[0].EndTime.Day())
		assert.Equal(t, 10, timeSlots[0].EndTime.Hour())

		assert.Equal(t, 23, timeSlots[3].StartTime.Day())
		assert.Equal(t, 11, timeSlots[3].StartTime.Hour())
		assert.Equal(t, 24, timeSlots[3].EndTime.Day())
		assert.Equal(t, 13, timeSlots[3].EndTime.Hour())

		assert.Equal(t, 23, timeSlots[8].StartTime.Day())
		assert.Equal(t, 16, timeSlots[8].StartTime.Hour())
		assert.Equal(t, 25, timeSlots[8].EndTime.Day())
		assert.Equal(t, 10, timeSlots[8].EndTime.Hour())
	})

	t.Run("weekend", func(t *testing.T) {
		date := time.Date(2024, 9, 26, 0, 0, 0, 0, time.UTC)
		serviceDuration := 15

		timeSlots := createTimeSlots(date, serviceDuration)

		require.Len(t, timeSlots, 9)

		assert.Equal(t, 26, timeSlots[0].StartTime.Day())
		assert.Equal(t, 8, timeSlots[0].StartTime.Hour())
		assert.Equal(t, 27, timeSlots[0].EndTime.Day())
		assert.Equal(t, 15, timeSlots[0].EndTime.Hour())

		assert.Equal(t, 26, timeSlots[3].StartTime.Day())
		assert.Equal(t, 11, timeSlots[3].StartTime.Hour())
		assert.Equal(t, 30, timeSlots[3].EndTime.Day())
		assert.Equal(t, 10, timeSlots[3].EndTime.Hour())

		assert.Equal(t, 26, timeSlots[8].StartTime.Day())
		assert.Equal(t, 16, timeSlots[8].StartTime.Hour())
		assert.Equal(t, 30, timeSlots[8].EndTime.Day())
		assert.Equal(t, 15, timeSlots[8].EndTime.Hour())
	})
}
