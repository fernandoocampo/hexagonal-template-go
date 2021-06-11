package people

import (
	"context"
	"errors"
	"log"
)

// SavePersonCommand defines an insert a person command.
type SavePersonCommand interface {
	ID() string
	Name() string
}

// Repository persist Person data.
type Repository interface {
	CreatePerson(ctx context.Context, person SavePersonCommand) error
}

// Service person service
type Service struct {
	repository Repository
}

// NewService creates a new Person service
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CreatePerson create a new person
func (p *Service) CreatePerson(ctx context.Context, newPerson NewPerson) error {
	log.Println("msg", "creating a person", "level", "DEBUG", "method", "person.Service.CreatePerson", "data", newPerson)
	createPerson := toCreatePerson(newPerson)
	// do any business validation here
	personToSave := createPerson.toSavePersonCommand()
	err := p.repository.CreatePerson(ctx, personToSave)
	if err != nil {
		log.Println("msg", "creating a person failed", "level", "ERROR", "error", err, "data", personToSave)
		return errors.New("new person could not be stored")
	}
	return nil
}
