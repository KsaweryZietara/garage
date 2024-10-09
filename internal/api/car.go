package api

import (
	"net/http"
	"strconv"
)

func (a *API) ListMakes(writer http.ResponseWriter, _ *http.Request) {
	makes, err := a.storage.Cars().ListMakes()
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, makes, 200)
}

func (a *API) ListModels(writer http.ResponseWriter, request *http.Request) {
	makeIDStr := request.PathValue("id")
	makeID, err := strconv.Atoi(makeIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	models, err := a.storage.Cars().ListModels(makeID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, models, 200)
}
