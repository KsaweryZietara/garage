package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCar(t *testing.T) {
	cleanup := NewSuite(t)
	defer cleanup()

	carRepo := NewCar(connection)

	makes, err := carRepo.ListMakes()
	assert.NoError(t, err)
	require.NotEmpty(t, makes)

	models, err := carRepo.ListModels(makes[0].ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, models)

	car, err := carRepo.GetByModelID(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, car.Make)
	assert.NotEmpty(t, car.Model)
}
