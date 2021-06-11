package anydb

// PersonEntity defines a person entity for any db.
type PersonEntity struct {
	ID   string
	Name string
}

func toPersonEntity(insertPersonCommand InsertPersonCommand) PersonEntity {
	return PersonEntity{
		ID:   insertPersonCommand.ID(),
		Name: insertPersonCommand.Name(),
	}
}
