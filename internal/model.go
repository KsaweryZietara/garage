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
	OwnerID     int
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
	}
}

type Service struct {
	ID       int
	Name     string
	Time     int
	Price    int
	GarageID int
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
	ServiceID  int
	EmployeeID int
	CustomerID int
}

func NewAppointment(dto CreateAppointmentDTO, customerID int) Appointment {
	return Appointment{
		StartTime:  dto.StartTime,
		EndTime:    dto.EndTime,
		ServiceID:  dto.ServiceID,
		EmployeeID: dto.EmployeeID,
		CustomerID: customerID,
	}
}

type TimeSlot struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
