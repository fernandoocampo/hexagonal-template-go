package anydb

import (
	"context"
	"errors"
	"log"
)

// InsertPersonCommand defines an insert a person command.
type InsertPersonCommand interface {
	ID() string
	Name() string
}

// Client defines logic for Any repository client.
type Client struct {
	anyClient *AnyClient
}

// NewClient creates a new any client
func NewClient() *Client {
	return &Client{
		anyClient: &AnyClient{},
	}
}

// CreatePerson persist the given person into the any database.
func (c *Client) CreatePerson(ctx context.Context, newPerson InsertPersonCommand) error {
	person := toPersonEntity(newPerson)
	err := c.anyClient.persist(ctx, person)
	if err != nil {
		log.Println("msg", "person could not be created", "method", "any.Client.CreatePerson")
		return errors.New("person could not be created, database is not available")
	}
	return nil
}

// AnyClient simulates an any external library.
type AnyClient struct {
}

func (a *AnyClient) persist(ctx context.Context, data interface{}) error {
	return errors.New("server unreachable")
}
