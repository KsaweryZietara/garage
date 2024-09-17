package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/mail"
	"github.com/KsaweryZietara/garage/internal/validate"

	"github.com/google/uuid"
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
		return
	}

	owner, err := a.storage.Employees().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	_, err = a.storage.Garages().GetByOwnerID(owner.ID)
	if err == nil {
		a.handleError(writer, errors.New("cannot create more than one garage"), 400)
		return
	}

	garage := internal.NewGarage(dto, owner.ID)
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
		employee, err := a.storage.Employees().Insert(
			internal.Employee{
				Email:    employeeEmail,
				Role:     internal.MechanicRole,
				GarageID: &garage.ID,
			})
		if err != nil {
			a.log.Error(err.Error())
			continue
		}
		code, err := a.storage.ConfirmationCodes().Insert(
			internal.ConfirmationCode{
				ID:         uuid.New().String(),
				EmployeeID: employee.ID,
			})
		if err != nil {
			a.log.Error(err.Error())
			continue
		}
		if err = a.mail.Send(
			employeeEmail,
			"Rejestracja",
			mail.NewEmployeeTemplate,
			mail.NewEmployee{
				GarageName: garage.Name,
				Code:       code.ID,
			},
		); err != nil {
			a.log.Error(err.Error())
		}
	}

	a.sendResponse(writer, nil, 201)
}
