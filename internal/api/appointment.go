package api

import (
	"encoding/json"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/validate"
)

func (a *API) CreateAppointment(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreateAppointmentDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateAppointmentDTO(dto)
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

	_, err = a.storage.Employees().GetByID(dto.EmployeeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	_, err = a.storage.Services().GetByID(dto.ServiceID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	appointment := internal.NewAppointment(dto, customer.ID)
	_, err = a.storage.Appointments().Insert(appointment)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}
