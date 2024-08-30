package validate

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/KsaweryZietara/garage/internal"
)

func RegisterDTO(dto internal.RegisterDTO) error {
	if dto.Name == "" || dto.Surname == "" || dto.Email == "" || dto.Password == "" || dto.ConfirmPassword == "" {
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

	if !isEmail(dto.Email) {
		return errors.New("invalid email format")
	}

	if !isPassword(dto.Password) {
		return errors.New("password must have at least one number, one capital letter and be at least 8 characters long")
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
