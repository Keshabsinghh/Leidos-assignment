package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter is the function which creates the mux router according to the paths and their handlers
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)

		router.
			Path(route.Pattern).
			Methods(route.Method).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/").HandlerFunc(Hello).Methods(http.MethodGet)
	return router
}
