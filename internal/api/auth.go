package api

import (
	"encoding/json"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/auth"
	"github.com/KsaweryZietara/garage/internal/validate"
)

func (a *API) Register(writer http.ResponseWriter, request *http.Request) {
	var dto internal.RegisterDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.RegisterDTO(dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	hash, err := auth.HashPassword(dto.Password)
	if err != nil {
		a.handleError(writer, err, 400)
	}
	dto.Password = hash

	employee := internal.NewEmployee(dto)
	err = a.storage.Employees().Insert(employee)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}
