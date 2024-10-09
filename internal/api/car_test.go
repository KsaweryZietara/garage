package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/KsaweryZietara/garage/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCar(t *testing.T) {
	suite := NewSuite(t)
	defer suite.Teardown()

	response := suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/makes"), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	var makes []internal.Make
	suite.ParseResponse(t, response, &makes)
	require.NotEmpty(t, makes)

	response = suite.CallAPI(http.MethodGet, fmt.Sprintf("/api/makes/%v/models", makes[0].ID), []byte{}, nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	var models []internal.Model
	suite.ParseResponse(t, response, &models)
	assert.NotEmpty(t, models)
}
