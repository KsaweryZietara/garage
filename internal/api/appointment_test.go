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

	mechanic, err := suite.api.storage.Employees().Insert(
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
		ModelID:    1,
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

	ownerToken, err := suite.api.auth.CreateToken("email", internal.OwnerRole)
	require.NoError(t, err)
	response = suite.CallAPI(http.MethodDelete, fmt.Sprintf("/api/services/%v", service.ID), []byte{}, &ownerToken)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	newAppointment := internal.CreateAppointmentDTO{
		StartTime:  time.Date(2030, 9, 25, 11, 0, 0, 0, time.UTC),
		EndTime:    time.Date(2030, 9, 25, 13, 0, 0, 0, time.UTC),
		ServiceID:  service.ID,
		EmployeeID: mechanic.ID,
		ModelID:    1,
	}
	newAppointmentJSON, err := json.Marshal(newAppointment)
	require.NoError(t, err)
	response = suite.CallAPI(http.MethodPost, "/api/appointments", newAppointmentJSON, token)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestGetEmployeeAndCustomerAppointmentsEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	customer, err := suite.api.storage.Customers().Insert(
		internal.Customer{
			Email:    "email",
			Password: "password",
		})
	assert.NoError(t, err)

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

	mechanic1, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "name1",
			Surname:   "surname1",
			Email:     "email2",
			Password:  "password",
			Role:      internal.MechanicRole,
			GarageID:  &garage.ID,
			Confirmed: true,
		})
	assert.NoError(t, err)

	mechanic2, err := suite.api.storage.Employees().Insert(
		internal.Employee{
			Name:      "name2",
			Surname:   "surname2",
			Email:     "email3",
			Password:  "password",
			Role:      internal.MechanicRole,
			GarageID:  &garage.ID,
			Confirmed: true,
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

	_, err = suite.api.storage.Appointments().Insert(
		internal.Appointment{
			StartTime:  time.Date(2024, 9, 23, 14, 0, 0, 0, time.UTC),
			EndTime:    time.Date(2024, 9, 25, 12, 0, 0, 0, time.UTC),
			ServiceID:  service.ID,
			EmployeeID: mechanic2.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		})
	assert.NoError(t, err)

	_, err = suite.api.storage.Appointments().Insert(
		internal.Appointment{
			StartTime:  time.Date(2024, 9, 22, 15, 0, 0, 0, time.UTC),
			EndTime:    time.Date(2024, 9, 23, 14, 0, 0, 0, time.UTC),
			ServiceID:  service.ID,
			EmployeeID: mechanic1.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		})
	assert.NoError(t, err)

	var appointmentDTOs []internal.AppointmentDTO

	token, err := suite.api.auth.CreateToken(mechanic1.Email, internal.MechanicRole)
	assert.NoError(t, err)
	response := suite.CallAPI(http.MethodGet, "/api/employees/appointments?date=2024-09-23", []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	suite.ParseResponse(t, response, &appointmentDTOs)
	assert.Len(t, appointmentDTOs, 1)
	assert.Equal(t, 23, appointmentDTOs[0].StartTime.Day())
	assert.Equal(t, 8, appointmentDTOs[0].StartTime.Hour())
	assert.Equal(t, 23, appointmentDTOs[0].EndTime.Day())
	assert.Equal(t, 14, appointmentDTOs[0].EndTime.Hour())
	assert.Nil(t, appointmentDTOs[0].Employee)

	token, err = suite.api.auth.CreateToken(owner.Email, internal.OwnerRole)
	assert.NoError(t, err)
	response = suite.CallAPI(http.MethodGet, "/api/employees/appointments?date=2024-09-23", []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	suite.ParseResponse(t, response, &appointmentDTOs)
	assert.Len(t, appointmentDTOs, 2)
	assert.Equal(t, 23, appointmentDTOs[0].StartTime.Day())
	assert.Equal(t, 8, appointmentDTOs[0].StartTime.Hour())
	assert.Equal(t, 23, appointmentDTOs[0].EndTime.Day())
	assert.Equal(t, 14, appointmentDTOs[0].EndTime.Hour())
	assert.NotNil(t, appointmentDTOs[0].Employee)
	assert.Equal(t, appointmentDTOs[0].Employee.Name, mechanic1.Name)
	assert.Equal(t, 23, appointmentDTOs[1].StartTime.Day())
	assert.Equal(t, 14, appointmentDTOs[1].StartTime.Hour())
	assert.Equal(t, 23, appointmentDTOs[1].EndTime.Day())
	assert.Equal(t, 16, appointmentDTOs[1].EndTime.Hour())
	assert.NotNil(t, appointmentDTOs[1].Employee)
	assert.Equal(t, appointmentDTOs[1].Employee.Name, mechanic2.Name)

	var customerAppointments internal.CustomerAppointmentDTOs
	token, err = suite.api.auth.CreateToken(customer.Email, internal.CustomerRole)
	assert.NoError(t, err)
	response = suite.CallAPI(http.MethodGet, "/api/customers/appointments", []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	suite.ParseResponse(t, response, &customerAppointments)
	assert.Len(t, customerAppointments.Upcoming, 0)
	assert.Len(t, customerAppointments.InProgress, 0)
	assert.Len(t, customerAppointments.Completed, 2)
}

func TestDeleteAppointmentEndpoint(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	customer, err := suite.api.storage.Customers().Insert(internal.Customer{
		Email:    "john.doe@example.com",
		Password: "Password123",
	})
	assert.NoError(t, err)

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

	mechanic, err := suite.api.storage.Employees().Insert(
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

	service, err := suite.api.storage.Services().Insert(
		internal.Service{
			Name:     "name",
			Time:     2,
			Price:    10,
			GarageID: garage.ID,
		})
	assert.NoError(t, err)

	appointment, err := suite.api.storage.Appointments().Insert(internal.Appointment{
		StartTime:  time.Now().Add(46 * time.Hour),
		EndTime:    time.Now().Add(48 * time.Hour),
		ServiceID:  service.ID,
		EmployeeID: mechanic.ID,
		CustomerID: customer.ID,
		ModelID:    1,
	})
	assert.NoError(t, err)

	token, err := suite.api.auth.CreateToken("john.doe@example.com", internal.CustomerRole)
	require.NoError(t, err)
	response := suite.CallAPI(http.MethodDelete, fmt.Sprintf("/api/appointments/%v", appointment.ID), []byte{}, &token)
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

func TestAppointmentsWithWorkingHours(t *testing.T) {
	appointments := []internal.Appointment{
		{
			StartTime: time.Date(2024, 9, 25, 14, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 9, 26, 11, 0, 0, 0, time.UTC),
		},
		{
			StartTime: time.Date(2024, 9, 26, 11, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 9, 26, 14, 0, 0, 0, time.UTC),
		},
		{
			StartTime: time.Date(2024, 9, 26, 14, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 9, 27, 12, 0, 0, 0, time.UTC),
		},
	}

	date := time.Date(2024, 9, 26, 0, 0, 0, 0, time.UTC)

	updatedAppointments := appointmentsWithWorkingHours(appointments, date)

	require.Len(t, updatedAppointments, len(appointments))

	assert.Equal(t, time.Date(2024, 9, 26, 8, 0, 0, 0, time.UTC), updatedAppointments[0].StartTime)
	assert.Equal(t, time.Date(2024, 9, 26, 11, 0, 0, 0, time.UTC), updatedAppointments[0].EndTime)

	assert.Equal(t, time.Date(2024, 9, 26, 11, 0, 0, 0, time.UTC), updatedAppointments[1].StartTime)
	assert.Equal(t, time.Date(2024, 9, 26, 14, 0, 0, 0, time.UTC), updatedAppointments[1].EndTime)

	assert.Equal(t, time.Date(2024, 9, 26, 14, 0, 0, 0, time.UTC), updatedAppointments[2].StartTime)
	assert.Equal(t, time.Date(2024, 9, 26, 16, 0, 0, 0, time.UTC), updatedAppointments[2].EndTime)
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		date     time.Time
		expected bool
	}{
		{time.Date(2024, 9, 23, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2024, 9, 25, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2024, 9, 27, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2024, 9, 28, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2024, 9, 29, 0, 0, 0, 0, time.UTC), true},
	}

	for _, test := range tests {
		result := isWeekend(test.date)
		assert.Equal(t, test.expected, result)
	}
}
