package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/validate"
)

func (a *API) CreateReview(writer http.ResponseWriter, request *http.Request) {
	appointmentIDStr := request.PathValue("id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	var dto internal.CreateReviewDTO
	err = json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateReviewDTO(dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	customer, err := a.storage.Customers().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 401)
		return
	}

	appointment, err := a.storage.Appointments().GetByID(appointmentID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	if customer.ID != appointment.CustomerID {
		a.handleError(writer, fmt.Errorf("customer ID %d does not match appointment customer ID %d", customer.ID, appointment.CustomerID), 400)
		return
	}

	if appointment.EndTime.After(time.Now()) {
		a.handleError(writer, fmt.Errorf("the appointment is not over yet"), 400)
		return
	}

	appointment.Rating = &dto.Rating
	appointment.Comment = &dto.Comment

	err = a.storage.Appointments().Update(appointment)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}

func (a *API) DeleteReview(writer http.ResponseWriter, request *http.Request) {
	appointmentIDStr := request.PathValue("id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	customer, err := a.storage.Customers().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 401)
		return
	}

	appointment, err := a.storage.Appointments().GetByID(appointmentID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	if customer.ID != appointment.CustomerID {
		a.handleError(writer, fmt.Errorf("customer ID %d does not match appointment customer ID %d", customer.ID, appointment.CustomerID), 400)
		return
	}

	if appointment.EndTime.After(time.Now()) {
		a.handleError(writer, fmt.Errorf("the appointment is not over yet"), 400)
		return
	}

	appointment.Rating = nil
	appointment.Comment = nil

	err = a.storage.Appointments().Update(appointment)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 200)
}

func (a *API) ListReviews(writer http.ResponseWriter, request *http.Request) {
	garageIDStr := request.PathValue("id")
	garageID, err := strconv.Atoi(garageIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	appointments, err := a.storage.Appointments().ListByGarageID(garageID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	reviewDTOs := make([]internal.ReviewDTO, len(appointments))
	for i, appointment := range appointments {
		service, err := a.storage.Services().GetByID(appointment.ServiceID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		employee, err := a.storage.Employees().GetByID(appointment.EmployeeID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		reviewDTOs[i] = internal.NewReviewDTO(appointment, service, employee)
	}

	a.sendResponse(writer, reviewDTOs, 200)
}
