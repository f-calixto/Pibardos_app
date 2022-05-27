package group

import (
	// Internal
	"github.com/coding-kiko/group_service/pkg/log"

	// Third party
	"github.com/gorilla/mux"
)

func NewRouter(handlers Handlers, logger log.Logger) *mux.Router {
	router := mux.NewRouter()
	logger.Info("router.go", "NewRouter", "Initializing handlers")

	router.Path("/").Methods("POST").HandlerFunc(handlers.CreateGroup)

	// router.Use(JwtMiddleware)
	return router
}
