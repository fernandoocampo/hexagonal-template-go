package people

import (
	"context"
	"errors"
	"log"

	"github.com/fernandoocampo/hexagonal-template-go/internal/adapters/storage"
)

// Repository define people repository behavior.
type Repository interface {
	CreatePerson(ctx context.Context, newPerson storage.InsertPersonCommand) error
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
	// then save the person
	personToSave := createPerson.toInsertPersonCommand()
	err := p.repository.CreatePerson(ctx, personToSave)
	if err != nil {
		log.Println("msg", "creating a person failed", "level", "ERROR", "error", err, "data", personToSave)
		return errors.New("new person could not be stored")
	}
	return nil
}
