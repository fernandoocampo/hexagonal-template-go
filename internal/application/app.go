package application

// Unity defines an application.
type Unity struct {
	Name string
}

// New creates a new application.
func New(args []string) *Unity {
	return &Unity{}
}

// Start initialize and start the unity
func (i *Unity) Start() error {
	return nil
}
