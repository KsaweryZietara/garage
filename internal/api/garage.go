package api

import (
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
)

func (a *API) GetEmployeeGarage(writer http.ResponseWriter, request *http.Request) {
	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	employee, err := a.storage.Employees().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	var garage internal.Garage
	switch employee.Role {
	case internal.Owner:
		garage, err = a.storage.Garages().GetByOwnerID(employee.ID)
	case internal.Mechanic:
		garage, err = a.storage.Garages().GetByID(*employee.GarageID)
	}
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	a.sendResponse(writer, internal.NewGarageDTO(garage), 200)
}
