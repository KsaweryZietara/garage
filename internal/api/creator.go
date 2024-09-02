package api

import (
	"encoding/json"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/validate"
)

func (a *API) Creator(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreatorDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreatorDTO(dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
	}

	employee, err := a.storage.Employees().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	garage := internal.NewGarage(dto, employee.ID)
	garage, err = a.storage.Garages().Insert(garage)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	for _, serviceDTO := range dto.Services {
		service := internal.NewService(serviceDTO, garage.ID)
		_, err = a.storage.Services().Insert(service)
		if err != nil {
			a.log.Error(err.Error())
		}
	}

	for _, employeeEmail := range dto.EmployeeEmails {
		_, err = a.storage.Employees().Insert(
			internal.Employee{
				Email:    employeeEmail,
				Role:     internal.Mechanic,
				GarageID: &garage.ID,
			})
		if err != nil {
			a.log.Error(err.Error())
		}
	}

	a.sendResponse(writer, nil, 201)
}
