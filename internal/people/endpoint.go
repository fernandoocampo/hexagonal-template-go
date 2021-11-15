package people

import (
	"context"
	"errors"
	"log"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a wrapper for endpoints
type Endpoints struct {
	service *Service
}

// NewEndpoints create a new people endpoints.
func NewEndpoints(service *Service) *Endpoints {
	return &Endpoints{
		service: service,
	}
}

// CreatePerson creates an endpoint function to create people
func (e *Endpoints) CreatePerson() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Println("msg", "creating a new person", "method", "people.Endpoints.CreatePerson", "person", request, "level", "DEBUG")
		newPerson, ok := request.(NewPerson)
		if !ok {
			log.Println(
				"msg", "the given request is not a new person object",
				"method", "people.Endpoints.CreatePerson",
				"level", "ERROR",
			)
			return "false", errors.New("invalid create person request")
		}

		err := e.service.CreatePerson(ctx, newPerson)
		if err != nil {
			log.Println(
				"msg", "the given new person could not be created",
				"method", "people.Endpoints.CreatePerson",
				"error", err,
				"person", newPerson,
				"level", "ERROR",
			)
			return "false", errors.New("person could not be created")
		}
		return "true", nil
	}
}
