package server

import (
	"currency-converter/common/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

//struct for main router
type Router struct {
	*mux.Router
}

//creates a new router
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

//Initialize routes
func (r *Router) InitializeRoutes() {
	r.Router.HandleFunc("/conversion", handlers.ConversionHandler).
		Methods(http.MethodPost).
		Name("conversion")
}
