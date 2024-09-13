package api

import (
	"net/http"
	"strconv"

	"github.com/KsaweryZietara/garage/internal"
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
