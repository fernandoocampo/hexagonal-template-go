package anydb

import (
	"context"

	"github.com/fernandoocampo/hexagonal-template-go/internal/adapters/storage"
)

// Connection defines anydb database behavior.
type Connection interface {
	Persist(ctx context.Context, record map[string]interface{}) error
}

func insertPersonToAnyRecord(icmd *storage.InsertPersonCommand) map[string]interface{} {
	return map[string]interface{}{
		"id":   icmd.ID,
		"name": icmd.Name,
	}
}
