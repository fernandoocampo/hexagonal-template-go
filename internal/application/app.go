package application

// Instance defines an application.
type Instance struct {
	Name string
}

// New creates a new application.
func New(args []string) *Instance {
	return &Instance{}
}

// Start initialize and start the Instance
func (i *Instance) Start() error {
	return nil
}
