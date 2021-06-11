package web

import (
	"github.com/go-kit/kit/endpoint"
)

// EndpointsCreator define a model to create service endpoints.
type EndpointsCreator interface {
	CreatePerson() endpoint.Endpoint
}

// CreatePersonRequest contains data to create a person.
type CreatePersonRequest struct {
	JName string `json:"name"`
}

// Name returns the name of the create person request.
func (c CreatePersonRequest) Name() string {
	return c.JName
}
