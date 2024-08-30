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
}

func NewEmployee(dto RegisterDTO) Employee {
	return Employee{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     dto.Role,
	}
}
