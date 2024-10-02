package api

import (
	"encoding/json"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/auth"
	"github.com/KsaweryZietara/garage/internal/validate"
)

func (a *API) CreateOwner(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreateEmployeeDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateEmployeeDTO(dto, true)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employee := internal.NewEmployee(dto, internal.OwnerRole)

	hash, err := auth.HashPassword(dto.Password)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}
	employee.Password = hash
	employee.Confirmed = true

	_, err = a.storage.Employees().Insert(employee)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}

func (a *API) CreateMechanic(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreateEmployeeDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateEmployeeDTO(dto, false)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employee := internal.NewEmployee(dto, internal.MechanicRole)

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
	employee.Confirmed = true

	if err = a.storage.Employees().Update(employee); err != nil {
		a.handleError(writer, err, 500)
		return
	}

	if err = a.storage.ConfirmationCodes().DeleteByID(codeID); err != nil {
		a.log.Error(err.Error())
	}

	a.sendResponse(writer, nil, 200)
}

func (a *API) CreateCustomer(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreateCustomerDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateCustomerDTO(dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	customer := internal.NewCustomer(dto)

	hash, err := auth.HashPassword(dto.Password)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}
	customer.Password = hash

	_, err = a.storage.Customers().Insert(customer)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}

func (a *API) LoginCustomer(writer http.ResponseWriter, request *http.Request) {
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

	customer, err := a.storage.Customers().GetByEmail(dto.Email)
	if err != nil {
		a.sendResponse(writer, nil, 401)
		return
	}

	if !auth.VerifyPassword(dto.Password, customer.Password) {
		a.sendResponse(writer, nil, 401)
		return
	}

	token, err := a.auth.CreateToken(customer.Email, internal.CustomerRole)
	if err != nil {
		a.sendResponse(writer, nil, 401)
		return
	}

	a.sendResponse(writer, token, 200)
}

func (a *API) LoginEmployee(writer http.ResponseWriter, request *http.Request) {
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
