package main

import (
	"net/http"

	"github.com/go-chi/chi"
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
	r := chi.NewRouter()
	h := handlers.ProvideHandler(svc)
	h.Router(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error().Err(err).Msg("Error")
	}
}
