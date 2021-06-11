package people

import "github.com/google/uuid"

// NewPerson contains new person data.
type NewPerson interface {
	Name() string
}

// toCreatePerson transform the given new person data to a create person
func toCreatePerson(newPerson NewPerson) *createPerson {
	return &createPerson{
		id:   uuid.New().String(),
		name: newPerson.Name(),
	}
}

// createPerson contains data to create a person
type createPerson struct {
	id   string
	name string
}

// ID returns the given person id.
func (c *createPerson) ID() string {
	return c.id
}

// Name returns the given person name.
func (c *createPerson) Name() string {
	return c.name
}

// toSavePersonCommand transform the given create person command
// to an save person command.
func (c *createPerson) toSavePersonCommand() SavePersonCommand {
	return c
}
