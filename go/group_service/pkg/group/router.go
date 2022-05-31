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
	router.Path("/{id}").Methods("GET").HandlerFunc(handlers.GetGroup)
	router.Path("/{id}").Methods("PATCH").HandlerFunc(handlers.UpdateGroup)
	router.Path("/{id}/avatar").Methods("PATCH").HandlerFunc(handlers.UpdateGroupAvatar)
	router.Path("/{id}/generateCode").Methods("POST").HandlerFunc(handlers.GenerateAccessCode)
	router.Path("/join").Methods("POST").HandlerFunc(handlers.JoinGroup)
	router.Path("/{id}/users").Methods("GET").HandlerFunc(handlers.GetGroupMembers)

	router.Use(JwtMiddleware)

	// override default gorilla 405 handler
	router.MethodNotAllowedHandler = handlers.MethodNotAllowedHandler()

	return router
}
