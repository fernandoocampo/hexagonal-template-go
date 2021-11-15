package people_test

import (
	"context"
	"testing"

	"github.com/fernandoocampo/hexagonal-template-go/internal/adapters/anydb"
	"github.com/fernandoocampo/hexagonal-template-go/internal/people"
)

func TestCreatePerson(t *testing.T) {
	// GIVEN
	newPersonName := "Fernando"
	newPerson := people.NewPerson{
		Name: newPersonName,
	}
	anyDBConnection := &AnyDBConnectionMock{}
	peopleRepository := anydb.NewClient(anyDBConnection)
	peopleService := people.NewService(peopleRepository)
	peopleEndpoints := people.NewEndpoints(peopleService).CreatePerson()
	ctx := context.TODO()

	// WHEN
	result, err := peopleEndpoints(ctx, newPerson)

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
	if anyDBConnection.data["id"] == "" {
		t.Errorf("person to save in the repository must have an ID, but got empty")
	}
	if anyDBConnection.data["name"] != newPersonName {
		t.Errorf("person to save in the repository must be: %q, but got: %q", newPersonName, anyDBConnection.data["name"])
	}
}

// AnyDBConnectionMock simulates a hypotetical external library.
type AnyDBConnectionMock struct {
	data map[string]interface{}
}

// Persist hypotetical persist method.
func (a *AnyDBConnectionMock) Persist(ctx context.Context, data map[string]interface{}) error {
	a.data = data
	return nil
}
