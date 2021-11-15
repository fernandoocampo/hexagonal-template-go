package anydb

import "context"

// Connection defines anydb database behavior.
type Connection interface {
	Persist(ctx context.Context, record map[string]interface{}) error
}

// InsertPersonCommand defines an insert a person command.
type InsertPersonCommand struct {
	ID   string
	Name string
}

func (i *InsertPersonCommand) toAnyRecord() map[string]interface{} {
	return map[string]interface{}{
		"id":   i.ID,
		"name": i.Name,
	}
}
