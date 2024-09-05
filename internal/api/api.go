package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/KsaweryZietara/garage/internal"
	"github.com/KsaweryZietara/garage/internal/auth"
	"github.com/KsaweryZietara/garage/internal/storage"

	"github.com/rs/cors"
)

const (
	bearerPrefix = "Bearer "
	emailKey     = "email"
)

type Config struct {
	Port string
}

type API struct {
	server  *http.Server
	log     *slog.Logger
	storage storage.Storage
	auth    *auth.Auth
}

func New(cfg Config, log *slog.Logger, storage storage.Storage, auth *auth.Auth) *API {
	return &API{
		server: &http.Server{
			Addr: fmt.Sprintf(":%s", cfg.Port),
		},
		log:     log,
		storage: storage,
		auth:    auth,
	}
}

func (a *API) Start() {
	router := http.NewServeMux()

	a.attachRoutes(router)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	a.server.Handler = c.Handler(router)

	a.log.Info("starting garage")
	log.Fatal(a.server.ListenAndServe())
}

func (a *API) attachRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /api/business/register", a.Register)
	router.HandleFunc("POST /api/business/login", a.Login)
	router.Handle("POST /api/business/creator", a.authMiddleware(http.HandlerFunc(a.Creator), internal.Owner))
	router.Handle("GET /api/garages", a.authMiddleware(http.HandlerFunc(a.GetGarage), internal.Owner))
}

func (a *API) authMiddleware(next http.Handler, role internal.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			a.sendResponse(w, nil, 401)
			return
		}
		if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			a.sendResponse(w, nil, 401)
			return
		}

		token := authHeader[len(bearerPrefix):]

		email, tokenRole, err := a.auth.VerifyToken(token)
		if err != nil {
			a.sendResponse(w, nil, 401)
			return
		}
		if tokenRole != role {
			a.sendResponse(w, nil, 401)
			return
		}

		ctx := context.WithValue(r.Context(), emailKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *API) emailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(emailKey).(string)
	return email, ok
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
