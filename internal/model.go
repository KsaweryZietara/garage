package internal

type Role string

const (
	Owner    Role = "OWNER"
	Mechanic Role = "MECHANIC"
)

type Employee struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
	Role     Role
	GarageID *int
}

func NewEmployee(dto RegisterDTO, role Role) Employee {
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

func NewGarage(dto CreatorDTO, ownerID int) Garage {
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
