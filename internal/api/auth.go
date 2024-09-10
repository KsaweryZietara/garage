package api

import (
	"encoding/json"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/auth"
	"github.com/KsaweryZietara/garage/internal/validate"
)

func (a *API) RegisterOwner(writer http.ResponseWriter, request *http.Request) {
	var dto internal.RegisterDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.RegisterDTO(dto, true)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employee := internal.NewEmployee(dto, internal.Owner)

	hash, err := auth.HashPassword(dto.Password)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}
	employee.Password = hash

	_, err = a.storage.Employees().Insert(employee)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}

func (a *API) RegisterMechanic(writer http.ResponseWriter, request *http.Request) {
	var dto internal.RegisterDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.RegisterDTO(dto, false)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employee := internal.NewEmployee(dto, internal.Mechanic)

	codeID := request.PathValue("code")
	code, err := a.storage.ConfirmationCodes().GetByID(codeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}
	employee.ID = code.EmployeeID

	hash, err := auth.HashPassword(dto.Password)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}
	employee.Password = hash

	if err = a.storage.Employees().Update(employee); err != nil {
		a.handleError(writer, err, 500)
		return
	}

	if err = a.storage.ConfirmationCodes().DeleteByID(codeID); err != nil {
		a.log.Error(err.Error())
	}

	a.sendResponse(writer, nil, 200)
}

func (a *API) Login(writer http.ResponseWriter, request *http.Request) {
	var dto internal.LoginDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.LoginDTO(dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employee, err := a.storage.Employees().GetByEmail(dto.Email)
	if err != nil {
		a.sendResponse(writer, nil, 401)
		return
	}

	if !auth.VerifyPassword(dto.Password, employee.Password) {
		a.sendResponse(writer, nil, 401)
		return
	}

	token, err := a.auth.CreateToken(employee.Email, employee.Role)
	if err != nil {
		a.sendResponse(writer, nil, 401)
		return
	}

	a.sendResponse(writer, token, 200)
}
