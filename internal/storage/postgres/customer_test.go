package postgres

import (
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
)

func TestCustomer(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	customerRepo := NewCustomer(connection)

	newCustomer := internal.Customer{
		Email:    "test@test.com",
		Password: "password123",
	}
	_, err := customerRepo.Insert(newCustomer)
	assert.NoError(t, err)

	retrievedCustomer, err := customerRepo.GetByEmail(newCustomer.Email)
	assert.NoError(t, err)
	assert.Equal(t, newCustomer.Email, retrievedCustomer.Email)
	assert.Equal(t, newCustomer.Password, retrievedCustomer.Password)
}
