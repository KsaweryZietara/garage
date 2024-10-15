package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/validate"
)

func (a *API) ListServices(writer http.ResponseWriter, request *http.Request) {
	garageIDStr := request.PathValue("id")
	garageID, err := strconv.Atoi(garageIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	services, err := a.storage.Services().ListByGarageID(garageID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, internal.NewServiceDTOs(services), 200)
}

func (a *API) GetService(writer http.ResponseWriter, request *http.Request) {
	serviceIDStr := request.PathValue("id")
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	service, err := a.storage.Services().GetByID(serviceID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, internal.NewServiceDTO(service), 200)
}

func (a *API) CreateService(writer http.ResponseWriter, request *http.Request) {
	var dto internal.ServiceDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateServiceDTO(dto)
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

	service := internal.NewService(dto, garage.ID)
	_, err = a.storage.Services().Insert(service)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}
