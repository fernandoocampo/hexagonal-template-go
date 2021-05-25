package domain

// Configuration contains data related to application configuration parameters.
type Configuration struct {
	ApplicationPort string `env:"APPLICATION_PORT" envDefault:":8080"`
}
