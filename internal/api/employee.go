package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KsaweryZietara/garage/internal"
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

	a.sendResponse(writer, internal.NewEmployeeDTOs(employees), 200)
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

	a.sendResponse(writer, internal.NewEmployeeDTOs(employees), 200)
}

func (a *API) GetEmployee(writer http.ResponseWriter, request *http.Request) {
	employeeIDStr := request.PathValue("id")
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employee, err := a.storage.Employees().GetByID(employeeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	if employee.Role == internal.OwnerRole {
		a.handleError(writer, fmt.Errorf("employee not found"), 404)
		return
	}

	a.sendResponse(writer, internal.NewEmployeeDTO(employee), 200)
}
