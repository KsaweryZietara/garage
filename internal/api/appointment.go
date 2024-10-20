package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/validate"
)

const (
	openingTime = 8
	closingTime = 16
)

func (a *API) CreateAppointment(writer http.ResponseWriter, request *http.Request) {
	var dto internal.CreateAppointmentDTO
	err := json.NewDecoder(request.Body).Decode(&dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	err = validate.CreateAppointmentDTO(dto)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	customer, err := a.storage.Customers().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 401)
		return
	}

	employee, err := a.storage.Employees().GetConfirmedByID(dto.EmployeeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	service, err := a.storage.Services().GetByID(dto.ServiceID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	if service.IsDeleted {
		a.sendResponse(writer, nil, 404)
		return
	}

	slotFound := false
	for _, slot := range createTimeSlots(dto.StartTime, service.Time) {
		if slot.StartTime.Equal(dto.StartTime) && slot.EndTime.Equal(dto.EndTime) {
			slotFound = true
			break
		}
	}
	if !slotFound {
		a.handleError(writer, errors.New("time slot not available"), 400)
		return
	}

	appointments, err := a.storage.Appointments().GetByTimeSlot(internal.TimeSlot{
		StartTime: dto.StartTime,
		EndTime:   dto.EndTime,
	}, employee.ID)
	if err != nil || len(appointments) != 0 {
		a.handleError(writer, errors.New("time slot not available"), 400)
		return
	}

	appointment := internal.NewAppointment(dto, customer.ID)
	_, err = a.storage.Appointments().Insert(appointment)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	a.sendResponse(writer, nil, 201)
}

func (a *API) GetAvailableSlots(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()

	serviceIDStr := queryParams.Get("serviceId")
	serviceID, err := strconv.Atoi(serviceIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}
	service, err := a.storage.Services().GetByID(serviceID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	employeeIDStr := queryParams.Get("employeeId")
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}
	employee, err := a.storage.Employees().GetConfirmedByID(employeeID)
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	dateStr := queryParams.Get("date")
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	var timeSlots []internal.TimeSlot
	for _, timeSlot := range createTimeSlots(date, service.Time) {
		appointments, err := a.storage.Appointments().GetByTimeSlot(timeSlot, employee.ID)
		if err == nil && len(appointments) == 0 {
			timeSlots = append(timeSlots, timeSlot)
		}
	}

	a.sendResponse(writer, timeSlots, 200)
}

func (a *API) GetEmployeeAppointments(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()

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

	dateStr := queryParams.Get("date")
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		a.handleError(writer, err, 400)
		return
	}

	if isWeekend(date) {
		a.sendResponse(writer, []internal.AppointmentDTO{}, 200)
		return
	}

	var appointments []internal.Appointment
	switch employee.Role {
	case internal.OwnerRole:
		garage, err := a.storage.Garages().GetByOwnerID(employee.ID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		appointments, err = a.storage.Appointments().GetByGarageID(garage.ID, date)
	case internal.MechanicRole:
		appointments, err = a.storage.Appointments().GetByEmployeeID(employee.ID, date)
	}
	if err != nil {
		a.handleError(writer, err, 404)
		return
	}

	appointments = appointmentsWithWorkingHours(appointments, date)

	appointmentDTOs := make([]internal.AppointmentDTO, len(appointments))
	for i, appointment := range appointments {
		service, err := a.storage.Services().GetByID(appointment.ServiceID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		car, err := a.storage.Cars().GetByModelID(appointment.ModelID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		appointmentDTOs[i] = internal.AppointmentDTO{
			ID:        appointment.ID,
			StartTime: appointment.StartTime,
			EndTime:   appointment.EndTime,
			Service:   internal.NewServiceDTO(service),
			Car:       car,
		}
		if employee.Role == internal.OwnerRole {
			mechanic, err := a.storage.Employees().GetConfirmedByID(appointment.EmployeeID)
			mechanicDTO := internal.NewEmployeeDTO(mechanic, false)
			appointmentDTOs[i].Employee = &mechanicDTO
			if err != nil {
				a.handleError(writer, err, 500)
				return
			}
		}
	}

	a.sendResponse(writer, appointmentDTOs, 200)
}

func (a *API) GetCustomerAppointments(writer http.ResponseWriter, request *http.Request) {
	email, ok := a.emailFromContext(request.Context())
	if !ok {
		a.sendResponse(writer, nil, 401)
		return
	}

	customer, err := a.storage.Customers().GetByEmail(email)
	if err != nil {
		a.handleError(writer, err, 401)
		return
	}

	appointments, err := a.storage.Appointments().GetByCustomerID(customer.ID)
	if err != nil {
		a.handleError(writer, err, 500)
		return
	}

	appointmentDTOs := make([]internal.AppointmentDTO, len(appointments))
	for i, appointment := range appointments {
		service, err := a.storage.Services().GetByID(appointment.ServiceID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		employee, err := a.storage.Employees().GetConfirmedByID(appointment.EmployeeID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		garage, err := a.storage.Garages().GetByID(*employee.GarageID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		car, err := a.storage.Cars().GetByModelID(appointment.ModelID)
		if err != nil {
			a.handleError(writer, err, 404)
			return
		}
		appointmentDTOs[i] = internal.NewAppointmentDTO(appointment, service, employee, garage, car)
	}

	a.sendResponse(writer, internal.NewCustomerAppointmentDTOs(appointmentDTOs), 200)
}

func createTimeSlots(date time.Time, serviceDuration int) []internal.TimeSlot {
	var timeSlots []internal.TimeSlot

	if isWeekend(date) {
		return timeSlots
	}

	startTime := time.Date(date.Year(), date.Month(), date.Day(), openingTime, 0, 0, 0, date.Location())

	for startTime.Hour() < closingTime+1 {
		endTime := startTime
		timeLeft := time.Duration(serviceDuration) * time.Hour

		for timeLeft > 0 {
			if endTime.Hour() == closingTime {
				if endTime.Weekday() == time.Friday {
					endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day()+3, openingTime, 0, 0, 0, endTime.Location())
				} else {
					endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day()+1, openingTime, 0, 0, 0, endTime.Location())
				}
			}
			endTime = endTime.Add(time.Hour)
			timeLeft = timeLeft - time.Hour
		}

		timeSlot := internal.TimeSlot{
			StartTime: startTime,
			EndTime:   endTime,
		}

		timeSlots = append(timeSlots, timeSlot)

		startTime = startTime.Add(time.Hour)
	}

	return timeSlots
}

func appointmentsWithWorkingHours(appointments []internal.Appointment, date time.Time) []internal.Appointment {
	for i := range appointments {
		if !appointments[i].StartTime.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
			appointments[i].StartTime = time.Date(date.Year(), date.Month(), date.Day(), 8, 0, 0, 0, date.Location())
		}
		if !appointments[i].EndTime.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
			appointments[i].EndTime = time.Date(date.Year(), date.Month(), date.Day(), 16, 0, 0, 0, date.Location())
		}
	}
	return appointments
}

func isWeekend(date time.Time) bool {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}
	return false
}
