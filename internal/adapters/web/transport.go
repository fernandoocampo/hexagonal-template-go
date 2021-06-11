package web

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a factory to create http servers for this project.
func NewHTTPServer(endpoints EndpointsCreator) http.Handler {
	router := mux.NewRouter()
	router.Methods(http.MethodPost).Path("/people").Handler(
		httptransport.NewServer(
			endpoints.CreatePerson(),
			decodeCreatePersonRequest,
			encodeCreateUserResponse,
		),
	)
	return router
}
