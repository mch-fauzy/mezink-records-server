package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mezink-records-server/configs"
	"github.com/mezink-records-server/infras"
	"github.com/mezink-records-server/internal/handlers"
	"github.com/mezink-records-server/internal/repository"
	"github.com/mezink-records-server/internal/service"

	"github.com/mezink-records-server/shared/logger"
	"github.com/rs/zerolog/log"
)

var config *configs.Config

func main() {

	// Initialize zerolog logger
	logger.InitLogger()

	// Initialize config
	config = configs.Get()

	// Create a new database connection
	db := infras.ProvideMySQLConn(config)

	// Initialize the repository with the database connection
	repo := repository.ProvideRepositoryMySql(db)

	// Initialize the service with the repository
	svc := service.ProvideService(repo)

	// Initialize the router
	// ToDo: change below into function
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := handlers.ProvideHandler(svc)
	h.Router(r)

	log.Info().Str("port", config.Server.Port).Msg("Starting up HTTP server.")
	err := http.ListenAndServe(":"+config.Server.Port, r)
	if err != nil {
		log.Error().Err(err).Msg("Error")
	}
}
