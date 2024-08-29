package api

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/storage"
)

type Config struct {
	Port string
}

type API struct {
	server  *http.Server
	log     *slog.Logger
	storage storage.Storage
}

func New(cfg Config, log *slog.Logger, storage storage.Storage) *API {
	return &API{
		server: &http.Server{
			Addr: fmt.Sprintf(":%s", cfg.Port),
		},
		log:     log,
		storage: storage,
	}
}

func (a *API) Start() {
	router := http.NewServeMux()

	a.attachRoutes(router)
	a.server.Handler = router

	a.log.Info("starting garage")
	log.Fatal(a.server.ListenAndServe())
}

func (a *API) attachRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /api/business/register", a.Register)
}

func (a *API) sendResponse(writer http.ResponseWriter, response interface{}, HTTPStatusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(HTTPStatusCode)
	if response == nil {
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		a.log.Error("unable to marshal response", "error", response)
		http.Error(writer, "unable to marshal response", http.StatusInternalServerError)
		return
	}

	_, err = writer.Write(responseJson)
	if err != nil {
		a.log.Error("unable to write response", "error", string(responseJson))
		http.Error(writer, "unable to write response", http.StatusInternalServerError)
	}
}

func (a *API) handleError(writer http.ResponseWriter, err error, HTTPStatusCode int) {
	a.log.Error("error occurred", "error", err)
	errorResponse := internal.Error{
		Message: err.Error(),
	}
	a.sendResponse(writer, errorResponse, HTTPStatusCode)
}
