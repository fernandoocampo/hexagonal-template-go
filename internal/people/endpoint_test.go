package people_test

import (
	"context"
	"testing"

	"github.com/fernandoocampo/myapp/internal/people"
)

func TestCreatePerson(t *testing.T) {
	// GIVEN
	newPersonName := "Fernando"
	newPerson := newPerson{
		name: newPersonName,
	}
	peopleRepository := &repositoryMock{}
	peopleService := people.NewService(peopleRepository)
	peopleEndpoints := people.NewEndpoints(peopleService)
	ctx := context.TODO()

	// WHEN
	result, err := peopleEndpoints.CreatePerson(ctx, newPerson)

	// THEN
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	resultValue, ok := result.(string)
	if !ok {
		t.Errorf("unexpected result type: %+v", resultValue)
	}
	if resultValue != "true" {
		t.Errorf("unexpected result: %s", resultValue)
	}
	if peopleRepository.personToSave.ID() == "" {
		t.Errorf("person to save in the repository must have an ID, but got empty")
	}
	if peopleRepository.personToSave.Name() != newPersonName {
		t.Errorf("person to save in the repository must be: %q, but got: %q", newPersonName, peopleRepository.personToSave.Name())
	}
}

// newPerson defines new person data.
type newPerson struct {
	name string
}

// Name returns the new person name.
func (n newPerson) Name() string {
	return n.name
}

// repositoryMock defines a mock for people repository
type repositoryMock struct {
	personToSave people.SavePersonCommand
}

func (r *repositoryMock) CreatePerson(ctx context.Context, person people.SavePersonCommand) error {
	r.personToSave = person
	return nil
}
