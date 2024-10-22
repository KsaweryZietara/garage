package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/mail"
	"github.com/KsaweryZietara/garage/internal/validate"

	"github.com/google/uuid"
)

func (a *API) ListConfirmedEmployees(writer http.ResponseWriter, request *http.Request) {
	garageIDStr := request.PathValue("id")
	garageID, err := strconv.Atoi(garageIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employees, err := a.storage.Employees().ListConfirmedByGarageID(garageID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, internal.NewEmployeeDTOs(employees, false), 200)
}

func (a *API) ListEmployees(writer http.ResponseWriter, request *http.Request) {
	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	owner, err := a.storage.Employees().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 401)
		return
	}

	garage, err := a.storage.Garages().GetByOwnerID(owner.ID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	employees, err := a.storage.Employees().ListByGarageID(garage.ID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, internal.NewEmployeeDTOs(employees, true), 200)
}

func (a *API) GetEmployee(writer http.ResponseWriter, request *http.Request) {
	employeeIDStr := request.PathValue("id")
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employee, err := a.storage.Employees().GetConfirmedByID(employeeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	if employee.Role == internal.OwnerRole {
		a.handleError(writer, fmt.Errorf("employee not found"), 404)
		return
	}

	a.sendResponse(writer, internal.NewEmployeeDTO(employee, false), 200)
}

func (a *API) CreateEmployee(writer http.ResponseWriter, request *http.Request) {
	var dto internal.EmployeeEmailDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	if !validate.IsEmail(dto.Email) {
		a.sendResponse(writer, nil, 400)
		return
	}

	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	owner, err := a.storage.Employees().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 401)
		return
	}

	garage, err := a.storage.Garages().GetByOwnerID(owner.ID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	employee, err := a.storage.Employees().Insert(
		internal.Employee{
			Email:     dto.Email,
			Role:      internal.MechanicRole,
			GarageID:  &garage.ID,
			Confirmed: false,
		})
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	code, err := a.storage.ConfirmationCodes().Insert(
		internal.ConfirmationCode{
			ID:         uuid.New().String(),
			EmployeeID: employee.ID,
		})
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	if err = a.mail.Send(
		dto.Email,
		"Rejestracja",
		mail.NewEmployeeTemplate,
		mail.NewEmployee{
			GarageName: garage.Name,
			Code:       code.ID,
		},
	); err != nil {
		a.log.Error(err.Error())
	}

	a.sendResponse(writer, nil, 201)
}

func (a *API) ResendConfirmationEmail(writer http.ResponseWriter, request *http.Request) {
	employeeIDStr := request.PathValue("id")
	employeeID, err := strconv.Atoi(employeeIDStr)
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
		a.handleError(writer, err, 401)
		return
	}

	garage, err := a.storage.Garages().GetByOwnerID(owner.ID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	employee, err := a.storage.Employees().GetByID(employeeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	if employee.GarageID != nil && *employee.GarageID != garage.ID {
		a.handleError(writer, errors.New("employee not found"), 404)
		return
	}

	if employee.Confirmed {
		a.handleError(writer, errors.New("employee is already confirmed"), 400)
		return
	}

	code, err := a.storage.ConfirmationCodes().Insert(
		internal.ConfirmationCode{
			ID:         uuid.New().String(),
			EmployeeID: employee.ID,
		})
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	if err = a.mail.Send(
		employee.Email,
		"Rejestracja",
		mail.NewEmployeeTemplate,
		mail.NewEmployee{
			GarageName: garage.Name,
			Code:       code.ID,
		},
	); err != nil {
		a.log.Error(err.Error())
	}

	a.sendResponse(writer, nil, 200)
}

func (a *API) DeleteEmployee(writer http.ResponseWriter, request *http.Request) {
	employeeIDStr := request.PathValue("id")
	employeeID, err := strconv.Atoi(employeeIDStr)
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
		a.handleError(writer, err, 401)
		return
	}

	garage, err := a.storage.Garages().GetByOwnerID(owner.ID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	employee, err := a.storage.Employees().GetByID(employeeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	if employee.GarageID != nil && *employee.GarageID != garage.ID {
		a.handleError(writer, errors.New("employee not found"), 404)
		return
	}

	err = a.storage.Employees().Delete(employee.ID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 200)
}
