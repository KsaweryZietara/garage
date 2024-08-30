package internal

type Error struct {
	Message string `json:"message"`
}

type RegisterDTO struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Role            Role   `json:"role"`
}
