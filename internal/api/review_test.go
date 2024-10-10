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

func TestCreateAndDeleteReviewEndpoint(t *testing.T) {
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

	mechanic, err := suite.api.storage.Employees().Insert(
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

	service, err := suite.api.storage.Services().Insert(
		internal.Service{
			Name:     "name",
			Time:     2,
			Price:    10,
			GarageID: garage.ID,
		})
	assert.NoError(t, err)

	appointment, err := suite.api.storage.Appointments().Insert(
		internal.Appointment{
			StartTime:  time.Now().Add(-3 * time.Hour),
			EndTime:    time.Now().Add(-2 * time.Hour),
			ServiceID:  service.ID,
			EmployeeID: mechanic.ID,
			CustomerID: customer.ID,
			ModelID:    1,
		})
	assert.NoError(t, err)

	token, err := suite.api.auth.CreateToken(customer.Email, internal.CustomerRole)
	assert.NoError(t, err)

	review := internal.CreateReviewDTO{
		Rating:  5,
		Comment: "comment",
	}
	reviewJSON, err := json.Marshal(review)
	require.NoError(t, err)
	response := suite.CallAPI(http.MethodPut, fmt.Sprintf("/api/appointments/%v/reviews", appointment.ID), reviewJSON, &token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	var customerAppointments1 internal.CustomerAppointmentDTOs
	response = suite.CallAPI(http.MethodGet, "/api/customers/appointments", []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	suite.ParseResponse(t, response, &customerAppointments1)
	assert.Len(t, customerAppointments1.Completed, 1)
	assert.Equal(t, review.Rating, *customerAppointments1.Completed[0].Rating)
	assert.Equal(t, review.Comment, *customerAppointments1.Completed[0].Comment)

	newReview := internal.CreateReviewDTO{
		Rating:  2,
		Comment: "newComment",
	}
	newReviewJSON, err := json.Marshal(newReview)
	require.NoError(t, err)
	response = suite.CallAPI(http.MethodPut, fmt.Sprintf("/api/appointments/%v/reviews", appointment.ID), newReviewJSON, &token)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	var customerAppointments2 internal.CustomerAppointmentDTOs
	response = suite.CallAPI(http.MethodGet, "/api/customers/appointments", []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	suite.ParseResponse(t, response, &customerAppointments2)
	assert.Len(t, customerAppointments2.Completed, 1)
	assert.Equal(t, newReview.Rating, *customerAppointments2.Completed[0].Rating)
	assert.Equal(t, newReview.Comment, *customerAppointments2.Completed[0].Comment)

	response = suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/garages/%v/reviews", garage.ID), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	var reviewDTOs []internal.ReviewDTO
	suite.ParseResponse(t, response, &reviewDTOs)
	assert.Len(t, reviewDTOs, 1)
	assert.Equal(t, reviewDTOs[0].ID, appointment.ID)
	assert.Equal(t, reviewDTOs[0].Service, service.Name)
	assert.Equal(t, reviewDTOs[0].Employee.Name, mechanic.Name)
	assert.Equal(t, reviewDTOs[0].Employee.Surname, mechanic.Surname)
	assert.Equal(t, reviewDTOs[0].Rating, 2)
	assert.Equal(t, *reviewDTOs[0].Comment, "newComment")

	response = suite.CallAPI(http.MethodDelete, fmt.Sprintf("/api/appointments/%v/reviews", appointment.ID), []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var customerAppointments3 internal.CustomerAppointmentDTOs
	response = suite.CallAPI(http.MethodGet, "/api/customers/appointments", []byte{}, &token)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	suite.ParseResponse(t, response, &customerAppointments3)
	assert.Len(t, customerAppointments3.Completed, 1)
	assert.Nil(t, customerAppointments3.Completed[0].Rating)
	assert.Nil(t, customerAppointments3.Completed[0].Comment)
}
