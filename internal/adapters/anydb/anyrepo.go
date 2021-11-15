package anydb

import (
	"context"
	"errors"
	"log"
)

// Client defines logic for Any repository client.
type Client struct {
	anyClient Connection
}

// NewClient creates a new any client
func NewClient(anyClient Connection) *Client {
	return &Client{
		anyClient: anyClient,
	}
}

// CreatePerson persist the given person into the any database.
func (c *Client) CreatePerson(ctx context.Context, newPerson InsertPersonCommand) error {
	person := newPerson.toAnyRecord()
	err := c.anyClient.Persist(ctx, person)
	if err != nil {
		log.Println("msg", "person could not be created", "method", "any.Client.CreatePerson")
		return errors.New("person could not be created, database is not available")
	}
	return nil
}

// AnyClient simulates a hypotetical external library.
type AnyClient struct {
}

// Persist hypotetical persist method.
func (a *AnyClient) Persist(ctx context.Context, data map[string]interface{}) error {
	return errors.New("server unreachable")
}
