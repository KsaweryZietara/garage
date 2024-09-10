package validate

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/KsaweryZietara/garage/internal"
)

func RegisterDTO(dto internal.RegisterDTO, validateEmail bool) error {
	if dto.Name == "" || dto.Surname == "" || dto.Password == "" || dto.ConfirmPassword == "" {
		return errors.New("fields cannot be empty")
	}

	if len(dto.Name) > 255 || len(dto.Surname) > 255 || len(dto.Email) > 255 || len(dto.Password) > 255 || len(dto.ConfirmPassword) > 255 {
		return errors.New("fields cannot have more than 255 characters")
	}

	if dto.Password != dto.ConfirmPassword {
		return errors.New("passwords must be identical")
	}

	if !isAlpha(dto.Name) || !isAlpha(dto.Surname) {
		return errors.New("name and surname cannot contain numbers")
	}

	if !isEmail(dto.Email) && validateEmail {
		return errors.New("invalid email format")
	}

	if !isPassword(dto.Password) {
		return errors.New("password must have at least one number, one capital letter and be at least 8 characters long")
	}

	return nil
}

func LoginDTO(dto internal.LoginDTO) error {
	if dto.Email == "" || dto.Password == "" {
		return errors.New("fields cannot be empty")
	}

	if !isEmail(dto.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

func CreatorDTO(dto internal.CreatorDTO) error {
	if dto.Name == "" || dto.City == "" || dto.Street == "" || dto.Number == "" || dto.PostalCode == "" || dto.PhoneNumber == "" {
		return errors.New("fields cannot be empty")
	}

	if len(dto.Name) > 255 || len(dto.City) > 255 || len(dto.Street) > 255 {
		return errors.New("name, city and street cannot have more than 255 characters")
	}

	if len(dto.Number) > 15 || len(dto.PostalCode) > 15 || len(dto.PhoneNumber) > 15 {
		return errors.New("number, postal code and phone number cannot have more than 15 characters")
	}

	if !isPostalCode(dto.PostalCode) {
		return errors.New("invalid postal code format")
	}

	if !isPhoneNumber(dto.PhoneNumber) {
		return errors.New("invalid phone number format")
	}

	for _, service := range dto.Services {
		if service.Name == "" {
			return errors.New("service name cannot be empty")
		}
		if len(service.Name) > 255 {
			return errors.New("service name cannot have more than 255 characters")
		}
		if service.Time <= 0 {
			return errors.New("service time must be greater than zero")
		}
		if service.Price <= 0 {
			return errors.New("service price must be greater than zero")
		}
	}

	for _, email := range dto.EmployeeEmails {
		if email == "" {
			return errors.New("email cannot be empty")
		}
		if len(email) > 255 {
			return errors.New("email cannot have more than 255 characters")
		}
		if !isEmail(email) {
			return errors.New("invalid email format")
		}
	}

	return nil
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isEmail(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(s)
}

func isPassword(s string) bool {
	hasDigit := false
	hasUpper := false

	if len(s) < 8 {
		return false
	}

	for _, r := range s {
		switch {
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsUpper(r):
			hasUpper = true
		}
	}

	return hasDigit && hasUpper
}

func isPostalCode(s string) bool {
	re := regexp.MustCompile(`^\d{2}-\d{3}$`)
	return re.MatchString(s)
}

func isPhoneNumber(s string) bool {
	re := regexp.MustCompile(`^\d{9}$`)
	return re.MatchString(s)
}
