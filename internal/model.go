package internal

type Employee struct {
	ID       int
	Name     string
	Surname  string
	Email    string
	Password string
}

func NewEmployee(dto RegisterDTO) Employee {
	return Employee{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
