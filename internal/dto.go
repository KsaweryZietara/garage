package internal

import "time"

type Error struct {
	Message string `json:"message"`
}

type Token struct {
	JWT string `json:"jwt"`
}

type CreateEmployeeDTO struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateGarageDTO struct {
	Name           string       `json:"name"`
	City           string       `json:"city"`
	Street         string       `json:"street"`
	Number         string       `json:"number"`
	PostalCode     string       `json:"postalCode"`
	PhoneNumber    string       `json:"phoneNumber"`
	Services       []ServiceDTO `json:"services"`
	EmployeeEmails []string     `json:"employeeEmails"`
}

type ServiceDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Time  int    `json:"time"`
	Price int    `json:"price"`
}

func NewServiceDTO(service Service) ServiceDTO {
	return ServiceDTO{
		ID:    service.ID,
		Name:  service.Name,
		Time:  service.Time,
		Price: service.Price,
	}
}

func NewServiceDTOs(services []Service) []ServiceDTO {
	serviceDTOs := make([]ServiceDTO, len(services))
	for i, service := range services {
		serviceDTOs[i] = NewServiceDTO(service)
	}
	return serviceDTOs
}

type GarageDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Street      string `json:"street"`
	Number      string `json:"number"`
	PostalCode  string `json:"postalCode"`
	PhoneNumber string `json:"phoneNumber"`
}

func NewGarageDTO(garage Garage) GarageDTO {
	return GarageDTO{
		ID:          garage.ID,
		Name:        garage.Name,
		City:        garage.City,
		Street:      garage.Street,
		Number:      garage.Number,
		PostalCode:  garage.PostalCode,
		PhoneNumber: garage.PhoneNumber,
	}
}

func NewGarageDTOs(garages []Garage) []GarageDTO {
	garageDTOs := make([]GarageDTO, len(garages))
	for i, garage := range garages {
		garageDTOs[i] = NewGarageDTO(garage)
	}
	return garageDTOs
}

type EmployeeDTO struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func NewEmployeeDTO(employee Employee) EmployeeDTO {
	return EmployeeDTO{
		ID:      employee.ID,
		Name:    employee.Name,
		Surname: employee.Surname,
	}
}

func NewEmployeeDTOs(employees []Employee) []EmployeeDTO {
	employeeDTOs := make([]EmployeeDTO, len(employees))
	for i, employee := range employees {
		employeeDTOs[i] = NewEmployeeDTO(employee)
	}
	return employeeDTOs
}

type CreateCustomerDTO struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type CreateAppointmentDTO struct {
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	ServiceID  int       `json:"serviceId"`
	EmployeeID int       `json:"employeeId"`
}

type AppointmentDTO struct {
	ID        int          `json:"id"`
	StartTime time.Time    `json:"startTime"`
	EndTime   time.Time    `json:"endTime"`
	Service   ServiceDTO   `json:"service"`
	Employee  *EmployeeDTO `json:"employee,omitempty"`
	Garage    *GarageDTO   `json:"garage,omitempty"`
}

func NewAppointmentDTO(appointment Appointment, service Service, employee Employee, garage Garage) AppointmentDTO {
	employeeDTO := NewEmployeeDTO(employee)
	garageDTO := NewGarageDTO(garage)
	return AppointmentDTO{
		ID:        appointment.ID,
		StartTime: appointment.StartTime,
		EndTime:   appointment.EndTime,
		Service:   NewServiceDTO(service),
		Employee:  &employeeDTO,
		Garage:    &garageDTO,
	}
}

type CustomerAppointmentDTOs struct {
	Upcoming   []AppointmentDTO `json:"upcoming"`
	InProgress []AppointmentDTO `json:"inProgress"`
	Completed  []AppointmentDTO `json:"completed"`
}

func NewCustomerAppointmentDTOs(appointments []AppointmentDTO) CustomerAppointmentDTOs {
	var result CustomerAppointmentDTOs
	now := time.Now()

	for _, appointment := range appointments {
		switch {
		case appointment.StartTime.After(now):
			result.Upcoming = append(result.Upcoming, appointment)

		case appointment.StartTime.Before(now) && appointment.EndTime.After(now):
			result.InProgress = append(result.InProgress, appointment)

		case appointment.EndTime.Before(now):
			result.Completed = append(result.Completed, appointment)
		}
	}

	return result
}
