package api

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/KsaweryZietara/garage/internal/auth"
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
}

func NewSuite(t *testing.T) *Suite {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	storage, cleanup, err := storage.NewForTests(connString, log)
	require.NoError(t, err)
	auth := auth.New("secret-key")
	api := New(Config{}, log, storage, auth)
	router := http.NewServeMux()
	api.attachRoutes(router)
	server := httptest.NewServer(router)

	return &Suite{
		t:       t,
		logger:  log,
		server:  server,
		client:  server.Client(),
		cleanup: cleanup,
	}
}

func (s *Suite) CallAPI(method string, path string, body []byte) *http.Response {
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", s.server.URL, path), bytes.NewBuffer(body))
	require.NoError(s.t, err)
	response, err := s.client.Do(request)
	require.NoError(s.t, err)
	return response
}

func (s *Suite) Teardown() {
	err := s.cleanup()
	require.NoError(s.t, err)
}