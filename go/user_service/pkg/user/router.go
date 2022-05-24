package user

import (
	// Internal
	"github.com/coding-kiko/user_service/pkg/log"

	// Third party
	"github.com/gorilla/mux"
)

func NewRouter(handlers Handlers, logger log.Logger) *mux.Router {
	router := mux.NewRouter()
	logger.Info("router.go", "NewRouter", "Initializing handlers")

	router.Path("/{id}").Methods("PATCH").HandlerFunc(handlers.UpdateUser)
	router.Path("/{id}/avatar").Methods("PATCH").HandlerFunc(handlers.UpdateUserAvatar)
	// router.Path("/{id}/groups").Methods("GET").HandlerFunc(handler.GetUserGroups)
	// router.Path("/{id}").Methods("GET").HandlerFunc(handler.GetUser)
	router.Use(JwtMiddleware)

	return router
}
