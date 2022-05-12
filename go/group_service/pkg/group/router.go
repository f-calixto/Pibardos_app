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
	// router.Path("/{id}").Methods("GET").HandlerFunc(handlers.GetGroup)
	router.Path("/").Methods("POST").HandlerFunc(handlers.CreateGroup)
	// router.Path("/generate_code").Methods("POST").HandlerFunc(handlers.GenerateAccessCode)
	// router.Path("/join").Methods("POST").HandlerFunc(handlers.JoinGroup)
	// router.Path("/{id}").Methods("PUT").HandlerFunc(handlers.UpdateGroup)
	router.Use(JwtMiddleware)
	return router
}
