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
	Role            Role   `json:"role"`
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
	Name  string `json:"name"`
	Time  int    `json:"time"`
	Price int    `json:"price"`
}
