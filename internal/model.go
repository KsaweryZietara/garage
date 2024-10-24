package internal

import "time"

type Role string

const (
	OwnerRole    Role = "OWNER"
	MechanicRole Role = "MECHANIC"
	CustomerRole Role = "CUSTOMER"
)

type Employee struct {
	ID        int
	Name      string
	Surname   string
	Email     string
	Password  string
	Role      Role
	GarageID  *int
	Confirmed bool
	IsDeleted bool
}

func NewEmployee(dto CreateEmployeeDTO, role Role) Employee {
	return Employee{
		Name:    dto.Name,
		Surname: dto.Surname,
		Email:   dto.Email,
		Role:    role,
	}
}

type Garage struct {
	ID          int
	Name        string
	City        string
	Street      string
	Number      string
	PostalCode  string
	PhoneNumber string
	Latitude    float64
	Longitude   float64
	OwnerID     int
	Rating      float64
	Distance    float64
	Logo        []byte
}

func NewGarage(dto CreateGarageDTO, ownerID int) Garage {
	return Garage{
		Name:        dto.Name,
		City:        dto.City,
		Street:      dto.Street,
		Number:      dto.Number,
		PostalCode:  dto.PostalCode,
		PhoneNumber: dto.PhoneNumber,
		OwnerID:     ownerID,
		Latitude:    dto.Latitude,
		Longitude:   dto.Longitude,
	}
}

type Service struct {
	ID        int
	Name      string
	Time      int
	Price     int
	IsDeleted bool
	GarageID  int
}

func NewService(dto ServiceDTO, garageID int) Service {
	return Service{
		Name:     dto.Name,
		Time:     dto.Time,
		Price:    dto.Price,
		GarageID: garageID,
	}
}

type ConfirmationCode struct {
	ID         string
	EmployeeID int
}

type Customer struct {
	ID       int
	Email    string
	Password string
}

func NewCustomer(dto CreateCustomerDTO) Customer {
	return Customer{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

type Appointment struct {
	ID         int
	StartTime  time.Time
	EndTime    time.Time
	Rating     *int
	Comment    *string
	ServiceID  int
	EmployeeID int
	CustomerID int
	ModelID    int
}

func NewAppointment(dto CreateAppointmentDTO, customerID int) Appointment {
	return Appointment{
		StartTime:  dto.StartTime,
		EndTime:    dto.EndTime,
		ServiceID:  dto.ServiceID,
		EmployeeID: dto.EmployeeID,
		CustomerID: customerID,
		ModelID:    dto.ModelID,
	}
}

type TimeSlot struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type Make struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	MakeID int    `json:"makeId"`
}

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}
