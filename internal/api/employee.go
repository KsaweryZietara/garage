package api

import (
	"net/http"
	"strconv"

	"github.com/KsaweryZietara/garage/internal"
)

func (a *API) ListEmployees(writer http.ResponseWriter, request *http.Request) {
	garageIDStr := request.PathValue("id")
	garageID, err := strconv.Atoi(garageIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	employees, err := a.storage.Employees().ListByGarageID(garageID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, internal.NewEmployeeDTOs(employees), 200)
}
