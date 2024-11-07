package api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/mail"
	"github.com/KsaweryZietara/garage/internal/validate"

	"github.com/google/uuid"
)

const (
	base64Prefix = "data:image/"
)

func (a *API) CreateGarage(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreateGarageDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateGarageDTO(dto)
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
		a.handleError(writer, err, 400)
		return
	}

	_, err = a.storage.Garages().GetByOwnerID(owner.ID)
	if err == nil {
		a.handleError(writer, errors.New("cannot create more than one garage"), 400)
		return
	}

	garage := internal.NewGarage(dto, owner.ID)
	garage, err = a.storage.Garages().Insert(garage)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	for _, serviceDTO := range dto.Services {
		service := internal.NewService(serviceDTO, garage.ID)
		_, err = a.storage.Services().Insert(service)
		if err != nil {
			a.log.Error(err.Error())
		}
	}

	for _, employeeEmail := range dto.EmployeeEmails {
		employee, err := a.storage.Employees().Insert(
			internal.Employee{
				Email:     employeeEmail,
				Role:      internal.MechanicRole,
				GarageID:  &garage.ID,
				Confirmed: false,
			})
		if err != nil {
			a.log.Error(err.Error())
			continue
		}
		code, err := a.storage.ConfirmationCodes().Insert(
			internal.ConfirmationCode{
				ID:         uuid.New().String(),
				EmployeeID: employee.ID,
			})
		if err != nil {
			a.log.Error(err.Error())
			continue
		}
		if err = a.mail.Send(
			employeeEmail,
			"Rejestracja",
			mail.NewEmployeeTemplate,
			mail.NewEmployee{
				GarageName: garage.Name,
				Code:       code.ID,
			},
		); err != nil {
			a.log.Error(err.Error())
		}
	}

	a.sendResponse(writer, nil, 201)
}

func (a *API) UpdateGarage(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreateGarageDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateGarageDTO(dto)
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

	garage.Name = dto.Name
	garage.City = dto.City
	garage.Street = dto.Street
	garage.Number = dto.Number
	garage.PostalCode = dto.PostalCode
	garage.PhoneNumber = dto.PhoneNumber
	garage.Latitude = dto.Latitude
	garage.Longitude = dto.Longitude

	err = a.storage.Garages().Update(garage)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 200)
}

func (a *API) GetEmployeeGarage(writer http.ResponseWriter, request *http.Request) {
	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	employee, err := a.storage.Employees().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 401)
		return
	}

	var garage internal.Garage
	switch employee.Role {
	case internal.OwnerRole:
		garage, err = a.storage.Garages().GetByOwnerID(employee.ID)
	case internal.MechanicRole:
		garage, err = a.storage.Garages().GetByID(*employee.GarageID)
	}
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	a.sendResponse(writer, internal.NewGarageDTO(garage), 200)
}

func (a *API) ListGarages(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	query := queryParams.Get("query")

	pageStr := queryParams.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	latitudeStr := queryParams.Get("latitude")
	latitude, err := strconv.ParseFloat(latitudeStr, 64)

	longitudeStr := queryParams.Get("longitude")
	longitude, err := strconv.ParseFloat(longitudeStr, 64)

	sortBy := queryParams.Get("sortBy")
	if sortBy != "rating" && sortBy != "distance" {
		sortBy = "rating"
	}

	garages, err := a.storage.Garages().List(page, query, latitude, longitude, sortBy)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, internal.NewGarageDTOs(garages), 200)
}

func (a *API) GetGarage(writer http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	garage, err := a.storage.Garages().GetByID(id)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	a.sendResponse(writer, internal.NewGarageDTO(garage), 200)
}

func (a *API) UpdateLogo(writer http.ResponseWriter, request *http.Request) {
	var dto internal.LogoDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	if len(dto.Base64Logo) == 0 {
		a.sendResponse(writer, nil, 400)
		return
	}

	if strings.HasPrefix(dto.Base64Logo, base64Prefix) {
		separatorIndex := strings.Index(dto.Base64Logo, ",")
		if separatorIndex == -1 {
			a.handleError(writer, fmt.Errorf("invalid base64 format"), 400)
			return
		}

		dto.Base64Logo = dto.Base64Logo[separatorIndex+1:]
	}

	decodedLogo, err := base64.StdEncoding.DecodeString(dto.Base64Logo)
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

	if err = a.storage.Garages().UpdateLogo(garage.ID, decodedLogo); err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 200)
}
