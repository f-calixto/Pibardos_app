package debts

import (
	// Internal
	"github.com/coding-kiko/debts_service/pkg/log"

	// Third party
	"github.com/gorilla/mux"
)

func NewRouter(handlers Handlers, logger log.Logger) *mux.Router {
	router := mux.NewRouter()
	logger.Info("router.go", "NewRouter", "Initializing handlers")

	router.Path("/{group_id}").Methods("POST").HandlerFunc(handlers.CreateRequest)
	router.Path("/{group_id}").Methods("GET").HandlerFunc(handlers.GetGroupDebts)
	router.Path("/{request_id}/accept").Methods("PATCH").HandlerFunc(handlers.AcceptRequest)
	router.Path("/{request_id}/reject").Methods("PATCH").HandlerFunc(handlers.RejectRequest)
	router.Path("/{group_id}/received/{user_id}").Methods("GET").HandlerFunc(handlers.GetReceivedRequests)
	router.Path("/{group_id}/sent/{user_id}").Methods("GET").HandlerFunc(handlers.GetSentRequests)
	router.Path("/{group_id}/cancel").Methods("DELETE").HandlerFunc(handlers.CancelDebt)

	router.Use(JwtMiddleware)

	// override default gorilla 405 handler
	router.MethodNotAllowedHandler = handlers.MethodNotAllowedHandler()

	return router
}
