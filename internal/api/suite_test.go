package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/auth"
	"github.com/KsaweryZietara/garage/internal/mail"
	"github.com/KsaweryZietara/garage/internal/storage"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var connString string

func TestMain(m *testing.M) {
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test-api"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
	)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	connString, err = container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	exitValue := m.Run()
	os.Exit(exitValue)
}

type Suite struct {
	t       *testing.T
	logger  *slog.Logger
	server  *httptest.Server
	client  *http.Client
	cleanup func() error
	api     *API
}

func NewSuite(t *testing.T) *Suite {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	storage, cleanup, err := storage.NewForTests(connString, log)
	require.NoError(t, err)
	auth := auth.New("secret-key")
	mail := mail.New(mail.Config{})

	api := New(Config{}, log, storage, auth, mail)

	router := http.NewServeMux()
	api.attachRoutes(router)
	server := httptest.NewServer(router)

	return &Suite{
		t:       t,
		logger:  log,
		server:  server,
		client:  server.Client(),
		cleanup: cleanup,
		api:     api,
	}
}

func (s *Suite) CreateEmployee(t *testing.T, employee internal.Employee) *internal.Token {
	_, err := s.api.storage.Employees().Insert(employee)
	require.NoError(t, err)
	token, err := s.api.auth.CreateToken(employee.Email, employee.Role)
	require.NoError(t, err)
	return &token
}

func (s *Suite) CreateCustomer(t *testing.T, customer internal.Customer) *internal.Token {
	_, err := s.api.storage.Customers().Insert(customer)
	require.NoError(t, err)
	token, err := s.api.auth.CreateToken(customer.Email, internal.CustomerRole)
	require.NoError(t, err)
	return &token
}

func (s *Suite) CallAPI(method string, path string, body []byte, token *internal.Token) *http.Response {
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", s.server.URL, path), bytes.NewBuffer(body))
	require.NoError(s.t, err)
	if token != nil {
		request.Header.Add("Authorization", "Bearer "+token.JWT)
	}
	response, err := s.client.Do(request)
	require.NoError(s.t, err)
	return response
}

func (s *Suite) Teardown() {
	err := s.cleanup()
	require.NoError(s.t, err)
}

func (s *Suite) ParseResponse(t *testing.T, response *http.Response, result interface{}) {
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	err = json.Unmarshal(body, result)
	require.NoError(t, err)
}
