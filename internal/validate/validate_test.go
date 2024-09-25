package validate

import (
	"strings"
	"testing"
	"time"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployeeDTO(t *testing.T) {
	t.Run("should return error when any field is empty", func(t *testing.T) {
		dto := internal.CreateEmployeeDTO{
			Name:            "",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateEmployeeDTO(dto, true)
		assert.EqualError(t, err, "fields cannot be empty")
	})

	t.Run("should return error when any field exceeds 255 characters", func(t *testing.T) {
		longString := strings.Repeat("a", 256)
		dto := internal.CreateEmployeeDTO{
			Name:            longString,
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateEmployeeDTO(dto, true)
		assert.EqualError(t, err, "fields cannot have more than 255 characters")
	})

	t.Run("should return error when name or surname contains numbers", func(t *testing.T) {
		dto := internal.CreateEmployeeDTO{
			Name:            "John123",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateEmployeeDTO(dto, true)
		assert.EqualError(t, err, "name and surname cannot contain numbers")
	})

	t.Run("should return error for invalid email format", func(t *testing.T) {
		dto := internal.CreateEmployeeDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "johnexample.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateEmployeeDTO(dto, true)
		assert.EqualError(t, err, "invalid email format")
	})

	t.Run("should return error for invalid password", func(t *testing.T) {
		dto := internal.CreateEmployeeDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "password",
			ConfirmPassword: "password",
		}
		err := CreateEmployeeDTO(dto, true)
		assert.EqualError(t, err, "password must have at least one number, one capital letter and be at least 8 characters long")
	})

	t.Run("should return error for different passwords", func(t *testing.T) {
		dto := internal.CreateEmployeeDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "password",
			ConfirmPassword: "password2",
		}
		err := CreateEmployeeDTO(dto, true)
		assert.EqualError(t, err, "passwords must be identical")
	})

	t.Run("should pass with valid input", func(t *testing.T) {
		dto := internal.CreateEmployeeDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateEmployeeDTO(dto, true)
		assert.NoError(t, err)
	})

	t.Run("should pass with valid input without email", func(t *testing.T) {
		dto := internal.CreateEmployeeDTO{
			Name:            "John",
			Surname:         "Smith",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateEmployeeDTO(dto, false)
		assert.NoError(t, err)
	})
}

func TestLoginDTO(t *testing.T) {
	t.Run("should return error when field is empty", func(t *testing.T) {
		dto := internal.LoginDTO{
			Email:    "",
			Password: "Password1",
		}
		err := LoginDTO(dto)
		assert.EqualError(t, err, "fields cannot be empty")
	})

	t.Run("should return error for invalid email format", func(t *testing.T) {
		dto := internal.LoginDTO{
			Email:    "johnexample.com",
			Password: "Password1",
		}
		err := LoginDTO(dto)
		assert.EqualError(t, err, "invalid email format")
	})

	t.Run("should pass with valid input", func(t *testing.T) {
		dto := internal.LoginDTO{
			Email:    "john@example.com",
			Password: "Password1",
		}
		err := LoginDTO(dto)
		assert.NoError(t, err)
	})
}

func TestIsAlpha(t *testing.T) {
	t.Run("should return true with only alphabetic characters", func(t *testing.T) {
		result := isAlpha("HelloWorld")
		assert.True(t, result)
	})

	t.Run("should return false with digits included", func(t *testing.T) {
		result := isAlpha("Hello123")
		assert.False(t, result)
	})

	t.Run("should return false with special characters included", func(t *testing.T) {
		result := isAlpha("Hello@World")
		assert.False(t, result)
	})

	t.Run("should return false with spaces included", func(t *testing.T) {
		result := isAlpha("Hello World")
		assert.False(t, result)
	})
}

func TestCreateGarageDTO(t *testing.T) {
	t.Run("should return error when any field is empty", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "123456789",
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "fields cannot be empty")
	})

	t.Run("should return error when any field exceeds character limit", func(t *testing.T) {
		longString := strings.Repeat("a", 256)
		dto := internal.CreateGarageDTO{
			Name:        longString,
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "123456789",
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "name, city and street cannot have more than 255 characters")
	})

	t.Run("should return error for invalid postal code format", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "Name",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12345",
			PhoneNumber: "123456789",
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "invalid postal code format")
	})

	t.Run("should return error for invalid phone number format", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "Name",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "12345678901",
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "invalid phone number format")
	})

	t.Run("should return error when service name is empty", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "Name",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "123456789",
			Services: []internal.ServiceDTO{
				{Name: "", Time: 1, Price: 1},
			},
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "service name cannot be empty")
	})

	t.Run("should return error when service time is zero", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "Name",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "123456789",
			Services: []internal.ServiceDTO{
				{Name: "Service", Time: 0, Price: 1},
			},
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "service time must be greater than zero")
	})

	t.Run("should return error when service price is zero", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "Name",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "123456789",
			Services: []internal.ServiceDTO{
				{Name: "Service", Time: 1, Price: 0},
			},
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "service price must be greater than zero")
	})

	t.Run("should return error for invalid email format", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "Name",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "123456789",
			EmployeeEmails: []string{
				"johnexample.com",
			},
		}
		err := CreateGarageDTO(dto)
		assert.EqualError(t, err, "invalid email format")
	})

	t.Run("should pass with valid input", func(t *testing.T) {
		dto := internal.CreateGarageDTO{
			Name:        "Name",
			City:        "City",
			Street:      "Street",
			Number:      "Number",
			PostalCode:  "12-345",
			PhoneNumber: "123456789",
			Services: []internal.ServiceDTO{
				{Name: "Service", Time: 1, Price: 1},
			},
			EmployeeEmails: []string{
				"john@example.com",
			},
		}
		err := CreateGarageDTO(dto)
		assert.NoError(t, err)
	})
}

func TestIsEmail(t *testing.T) {
	t.Run("should return true with valid email", func(t *testing.T) {
		result := isEmail("test@example.com")
		assert.True(t, result)
	})

	t.Run("should return false with missing at symbol", func(t *testing.T) {
		result := isEmail("testexample.com")
		assert.False(t, result)
	})

	t.Run("should return false with missing domain", func(t *testing.T) {
		result := isEmail("test@")
		assert.False(t, result)
	})

	t.Run("should return false with invalid domain", func(t *testing.T) {
		result := isEmail("test@example")
		assert.False(t, result)
	})

	t.Run("should return false with special characters not allowed", func(t *testing.T) {
		result := isEmail("test@exa$mple.com")
		assert.False(t, result)
	})

	t.Run("should return true with valid email including special characters", func(t *testing.T) {
		result := isEmail("test.email+alex@leetcode.com")
		assert.True(t, result)
	})
}

func TestIsPassword(t *testing.T) {
	t.Run("should return true with valid password", func(t *testing.T) {
		result := isPassword("Password1")
		assert.True(t, result)
	})

	t.Run("should return false with no uppercase", func(t *testing.T) {
		result := isPassword("password1")
		assert.False(t, result)
	})

	t.Run("should return false with no digit", func(t *testing.T) {
		result := isPassword("Password")
		assert.False(t, result)
	})

	t.Run("should return false with less than eight characters", func(t *testing.T) {
		result := isPassword("Pass1")
		assert.False(t, result)
	})
}

func TestIsPostalCode(t *testing.T) {
	t.Run("should return true with valid postal code", func(t *testing.T) {
		result := isPostalCode("12-345")
		assert.True(t, result)
	})

	t.Run("should return false with missing dash", func(t *testing.T) {
		result := isPostalCode("12345")
		assert.False(t, result)
	})

	t.Run("should return false with extra characters", func(t *testing.T) {
		result := isPostalCode("12-3456")
		assert.False(t, result)
	})

	t.Run("should return false with letters included", func(t *testing.T) {
		result := isPostalCode("12-AB5")
		assert.False(t, result)
	})
}

func TestIsPhoneNumber(t *testing.T) {
	t.Run("should return true with valid phone number", func(t *testing.T) {
		result := isPhoneNumber("123456789")
		assert.True(t, result)
	})

	t.Run("should return false with less than 9 digits", func(t *testing.T) {
		result := isPhoneNumber("+12345678")
		assert.False(t, result)
	})

	t.Run("should return false with more than 15 digits", func(t *testing.T) {
		result := isPhoneNumber("+1234567890123456")
		assert.False(t, result)
	})

	t.Run("should return false with letters included", func(t *testing.T) {
		result := isPhoneNumber("+12345678A901")
		assert.False(t, result)
	})
}

func TestCreateCustomerDTO(t *testing.T) {
	t.Run("should return error when any field is empty", func(t *testing.T) {
		dto := internal.CreateCustomerDTO{
			Email:           "john@example.com",
			Password:        "",
			ConfirmPassword: "Password1",
		}
		err := CreateCustomerDTO(dto)
		assert.EqualError(t, err, "fields cannot be empty")
	})

	t.Run("should return error when any field exceeds 255 characters", func(t *testing.T) {
		longString := strings.Repeat("a", 256)
		dto := internal.CreateCustomerDTO{
			Email:           longString,
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateCustomerDTO(dto)
		assert.EqualError(t, err, "fields cannot have more than 255 characters")
	})

	t.Run("should return error for invalid email format", func(t *testing.T) {
		dto := internal.CreateCustomerDTO{
			Email:           "johnexample.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateCustomerDTO(dto)
		assert.EqualError(t, err, "invalid email format")
	})

	t.Run("should return error for invalid password", func(t *testing.T) {
		dto := internal.CreateCustomerDTO{
			Email:           "john@example.com",
			Password:        "password",
			ConfirmPassword: "password",
		}
		err := CreateCustomerDTO(dto)
		assert.EqualError(t, err, "password must have at least one number, one capital letter and be at least 8 characters long")
	})

	t.Run("should return error for different passwords", func(t *testing.T) {
		dto := internal.CreateCustomerDTO{
			Email:           "john@example.com",
			Password:        "password",
			ConfirmPassword: "password2",
		}
		err := CreateCustomerDTO(dto)
		assert.EqualError(t, err, "passwords must be identical")
	})

	t.Run("should pass with valid input", func(t *testing.T) {
		dto := internal.CreateCustomerDTO{
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
		}
		err := CreateCustomerDTO(dto)
		assert.NoError(t, err)
	})
}

func TestCreateAppointmentDTO(t *testing.T) {
	t.Run("should return error when start time or end time is zero", func(t *testing.T) {
		dto := internal.CreateAppointmentDTO{
			StartTime:  time.Time{},
			EndTime:    time.Now(),
			ServiceID:  1,
			EmployeeID: 1,
		}
		err := CreateAppointmentDTO(dto)
		assert.EqualError(t, err, "start time and end time cannot be empty")
	})

	t.Run("should return error when start time is before current date", func(t *testing.T) {
		dto := internal.CreateAppointmentDTO{
			StartTime:  time.Now().Add(-time.Hour),
			EndTime:    time.Now().Add(time.Hour),
			ServiceID:  1,
			EmployeeID: 1,
		}
		err := CreateAppointmentDTO(dto)
		assert.EqualError(t, err, "start time cannot be in the past")
	})

	t.Run("should return error when end time is before start time", func(t *testing.T) {
		dto := internal.CreateAppointmentDTO{
			StartTime:  time.Now().Add(2 * time.Hour),
			EndTime:    time.Now().Add(time.Hour),
			ServiceID:  1,
			EmployeeID: 1,
		}
		err := CreateAppointmentDTO(dto)
		assert.EqualError(t, err, "end time must be after start time")
	})

	t.Run("should return error when service ID is less than or equal to zero", func(t *testing.T) {
		dto := internal.CreateAppointmentDTO{
			StartTime:  time.Now().Add(time.Hour),
			EndTime:    time.Now().Add(2 * time.Hour),
			ServiceID:  0,
			EmployeeID: 1,
		}
		err := CreateAppointmentDTO(dto)
		assert.EqualError(t, err, "service ID must be greater than zero")
	})

	t.Run("should return error when employee ID is less than or equal to zero", func(t *testing.T) {
		dto := internal.CreateAppointmentDTO{
			StartTime:  time.Now().Add(time.Hour),
			EndTime:    time.Now().Add(2 * time.Hour),
			ServiceID:  1,
			EmployeeID: 0,
		}
		err := CreateAppointmentDTO(dto)
		assert.EqualError(t, err, "employee ID must be greater than zero")
	})

	t.Run("should pass with valid input", func(t *testing.T) {
		dto := internal.CreateAppointmentDTO{
			StartTime:  time.Now().Add(time.Hour),
			EndTime:    time.Now().Add(2 * time.Hour),
			ServiceID:  1,
			EmployeeID: 1,
		}
		err := CreateAppointmentDTO(dto)
		assert.NoError(t, err)
	})
}
