package validate

import (
	"errors"
	"regexp"
	"time"
	"unicode"

	"github.com/KsaweryZietara/garage/internal"
)

func CreateEmployeeDTO(dto internal.CreateEmployeeDTO, validateEmail bool) error {
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

func CreateGarageDTO(dto internal.CreateGarageDTO) error {
	if dto.Name == "" || dto.City == "" || dto.Street == "" || dto.Number == "" || dto.PostalCode == "" || dto.PhoneNumber == "" {
		return errors.New("fields cannot be empty")
	}

	if len(dto.Name) > 255 || len(dto.City) > 255 || len(dto.Street) > 255 {
		return errors.New("name, city and street cannot have more than 255 characters")
	}

	if len(dto.Number) > 15 || len(dto.PostalCode) > 15 || len(dto.PhoneNumber) > 15 {
		return errors.New("number, postal code and phone number cannot have more than 15 characters")
	}

	if dto.Latitude == 0 && dto.Longitude == 0 {
		return errors.New("coordinates cannot be zeros")
	}

	if dto.Latitude < -90 || dto.Latitude > 90 {
		return errors.New("latitude must be between -90 and 90 degrees")
	}

	if dto.Longitude < -180 || dto.Longitude > 180 {
		return errors.New("longitude must be between -180 and 180 degrees")
	}

	if !isPostalCode(dto.PostalCode) {
		return errors.New("invalid postal code format")
	}

	if !isPhoneNumber(dto.PhoneNumber) {
		return errors.New("invalid phone number format")
	}

	for _, service := range dto.Services {
		if err := CreateServiceDTO(service); err != nil {
			return err
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

func CreateCustomerDTO(dto internal.CreateCustomerDTO) error {
	if dto.Password == "" || dto.ConfirmPassword == "" {
		return errors.New("fields cannot be empty")
	}

	if len(dto.Email) > 255 || len(dto.Password) > 255 || len(dto.ConfirmPassword) > 255 {
		return errors.New("fields cannot have more than 255 characters")
	}

	if dto.Password != dto.ConfirmPassword {
		return errors.New("passwords must be identical")
	}

	if !isEmail(dto.Email) {
		return errors.New("invalid email format")
	}

	if !isPassword(dto.Password) {
		return errors.New("password must have at least one number, one capital letter and be at least 8 characters long")
	}

	return nil
}

func CreateAppointmentDTO(dto internal.CreateAppointmentDTO) error {
	if dto.StartTime.IsZero() || dto.EndTime.IsZero() {
		return errors.New("start time and end time cannot be empty")
	}

	if dto.StartTime.Before(time.Now()) {
		return errors.New("start time cannot be in the past")
	}

	if dto.EndTime.Before(dto.StartTime) {
		return errors.New("end time must be after start time")
	}

	if dto.ServiceID <= 0 {
		return errors.New("service ID must be greater than zero")
	}

	if dto.EmployeeID <= 0 {
		return errors.New("employee ID must be greater than zero")
	}

	if dto.ModelID <= 0 {
		return errors.New("model ID must be greater than zero")
	}

	return nil
}

func CreateReviewDTO(dto internal.CreateReviewDTO) error {
	if dto.Rating < 1 || dto.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	if len(dto.Comment) > 2048 {
		return errors.New("comment must not exceed 2048 characters")
	}

	return nil
}

func CreateServiceDTO(dto internal.ServiceDTO) error {
	if dto.Name == "" {
		return errors.New("service name cannot be empty")
	}

	if len(dto.Name) > 255 {
		return errors.New("service name cannot have more than 255 characters")
	}

	if dto.Time <= 0 {
		return errors.New("service time must be greater than zero")
	}

	if dto.Price <= 0 {
		return errors.New("service price must be greater than zero")
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
