package validate

import (
	"strings"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestRegisterDTO(t *testing.T) {
	t.Run("should return error when any field is empty", func(t *testing.T) {
		dto := internal.RegisterDTO{
			Name:            "",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
			Role:            internal.Mechanic,
		}
		err := RegisterDTO(dto)
		assert.EqualError(t, err, "fields cannot be empty")
	})

	t.Run("should return error when any field exceeds 255 characters", func(t *testing.T) {
		longString := strings.Repeat("a", 256)
		dto := internal.RegisterDTO{
			Name:            longString,
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
			Role:            internal.Mechanic,
		}
		err := RegisterDTO(dto)
		assert.EqualError(t, err, "fields cannot have more than 255 characters")
	})

	t.Run("should return error when name or surname contains numbers", func(t *testing.T) {
		dto := internal.RegisterDTO{
			Name:            "John123",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
			Role:            internal.Owner,
		}
		err := RegisterDTO(dto)
		assert.EqualError(t, err, "name and surname cannot contain numbers")
	})

	t.Run("should return error for invalid email format", func(t *testing.T) {
		dto := internal.RegisterDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "johnexample.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
			Role:            internal.Owner,
		}
		err := RegisterDTO(dto)
		assert.EqualError(t, err, "invalid email format")
	})

	t.Run("should return error for invalid password", func(t *testing.T) {
		dto := internal.RegisterDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "password",
			ConfirmPassword: "password",
			Role:            internal.Mechanic,
		}
		err := RegisterDTO(dto)
		assert.EqualError(t, err, "password must have at least one number, one capital letter and be at least 8 characters long")
	})

	t.Run("should return error for different passwords", func(t *testing.T) {
		dto := internal.RegisterDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "password",
			ConfirmPassword: "password2",
			Role:            internal.Owner,
		}
		err := RegisterDTO(dto)
		assert.EqualError(t, err, "passwords must be identical")
	})

	t.Run("should pass with valid input", func(t *testing.T) {
		dto := internal.RegisterDTO{
			Name:            "John",
			Surname:         "Smith",
			Email:           "john@example.com",
			Password:        "Password1",
			ConfirmPassword: "Password1",
			Role:            internal.Owner,
		}
		err := RegisterDTO(dto)
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
