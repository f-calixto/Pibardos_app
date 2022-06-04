package calendar

import (
	// Internal
	"github.com/coding-kiko/calendar_service/pkg/log"

	// Third party
	"github.com/gorilla/mux"
)

func NewRouter(handlers Handlers, logger log.Logger) *mux.Router {
	router := mux.NewRouter()
	logger.Info("router.go", "NewRouter", "Initializing handlers")

	router.Path("/{group_id}").Methods("POST").HandlerFunc(handlers.NewEvent)
	router.Path("/join/{event_id}").Methods("POST").HandlerFunc(handlers.JoinEvent)
	router.Path("/{group_id}").Methods("GET").HandlerFunc(handlers.GetEvents)
	router.Path("/{event_id}").Methods("DELETE").HandlerFunc(handlers.CancelEvent)
	router.Path("/{group_id}/{event_id}").Methods("PATCH").HandlerFunc(handlers.UpdateEvent)
	router.Use(JwtMiddleware)

	// override default gorilla 405 handler
	router.MethodNotAllowedHandler = handlers.MethodNotAllowedHandler()

	return router
}
