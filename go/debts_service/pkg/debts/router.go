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

	router.Path("/{group_id}").Methods("POST").HandlerFunc(handlers.CreateDebt)
	router.Path("/{group_id}/sent").Methods("GET").HandlerFunc(handlers.GetSentDebts)
	router.Path("/{group_id}/received").Methods("GET").HandlerFunc(handlers.GetReceivedDebts)
	router.Path("/{debt_id}/accept").Methods("PATCH").HandlerFunc(handlers.AcceptDebt)
	router.Path("/{debt_id}/reject").Methods("PATCH").HandlerFunc(handlers.RejectDebt)
	router.Path("/{debt_id}/cancel").Methods("PATCH").HandlerFunc(handlers.CancelDebt)

	router.Use(JwtMiddleware)

	// override default gorilla 405 handler
	router.MethodNotAllowedHandler = handlers.MethodNotAllowedHandler()

	return router
}
