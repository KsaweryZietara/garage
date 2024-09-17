package internal

type Error struct {
	Message string `json:"message"`
}

type Token struct {
	JWT string `json:"jwt"`
}

type RegisterDTO struct {
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

type CreatorDTO struct {
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
