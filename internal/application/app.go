package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fernandoocampo/hexagonal-template-go/internal/adapters/anydb"
	"github.com/fernandoocampo/hexagonal-template-go/internal/adapters/web"
	"github.com/fernandoocampo/hexagonal-template-go/internal/configurations"
	"github.com/fernandoocampo/hexagonal-template-go/internal/people"
)

// Instance defines an application.
type Instance struct {
	Name          string
	configuration configurations.Application
	dbConn        *anyDBConnection
}

// Event contains an application event.
type Event struct {
	Message string
	Error   error
}

// New creates a new application.
func New(args []string) *Instance {
	return &Instance{}
}

// Start initialize and start the Instance
func (i *Instance) Start() error {
	log.Println("level", "INFO", "msg", "starting application")

	confError := i.loadConfiguration()
	if confError != nil {
		panic(confError)
	}
	log.Println("level", "DEBUG", "msg", "application configuration", "parameters", i.configuration)

	log.Println("level", "INFO", "msg", "starting database connection")
	err := i.openDBConnection()
	if err != nil {
		log.Println("level", "ERROR", "msg", "database connection could not be stablished")
		return err
	}

	peopleRepository := i.createPeopleRepository()
	peopleService := people.NewService(peopleRepository)
	peopleEndpoints := people.NewEndpoints(peopleService)

	eventStream := make(chan Event)
	i.listenToOSSignal(eventStream)
	i.startWebServer(peopleEndpoints, eventStream)

	eventMessage := <-eventStream
	fmt.Println(
		"level", "INFO",
		"msg", "ending server",
		"event", eventMessage.Message,
	)

	if eventMessage.Error != nil {
		fmt.Println(
			"level", "ERROR",
			"msg", "ending server with error",
			"error", eventMessage.Error,
		)
		return eventMessage.Error
	}
	return nil
}

// Stop stop application, take advantage of this to clean resources
func (i *Instance) Stop() {
	log.Println("level", "INFO", "msg", "stopping the application")
	if i.dbConn != nil {
		i.dbConn.Close()
	}
}

func (i *Instance) listenToOSSignal(eventStream chan<- Event) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		osSignal := fmt.Sprintf("%s", <-c)
		event := Event{
			Message: osSignal,
		}
		eventStream <- event
	}()
}

// startWebServer starts the web server.
func (i *Instance) startWebServer(endpoints *people.Endpoints, eventStream chan<- Event) {
	go func() {
		log.Println("msg", "starting http server", "http:", i.configuration.ApplicationPort)
		handler := web.NewHTTPServer(endpoints)
		err := http.ListenAndServe(i.configuration.ApplicationPort, handler)
		if err != nil {
			eventStream <- Event{
				Message: "web server was ended with error",
				Error:   err,
			}
			return
		}
		eventStream <- Event{
			Message: "web server was ended",
		}
	}()
}

func (i *Instance) loadConfiguration() error {
	applicationSetUp, err := configurations.Load()
	if err != nil {
		log.Println("level", "ERROR", "msg", "application setup could not be loaded", "error", err)
		return errors.New("application setup could not be loaded")
	}
	i.configuration = applicationSetUp
	return nil
}

func (i *Instance) createPeopleRepository() *anydb.Client {
	newPeopleRepository := anydb.NewClient(i.dbConn)
	return newPeopleRepository
}

func (i *Instance) openDBConnection() error {
	newAnyDBConnection := anyDBConnection{
		data: make(map[string]interface{}),
	}
	i.dbConn = &newAnyDBConnection
	return nil
}

// anyDBConnection simulates a hypotetical external library.
type anyDBConnection struct {
	data map[string]interface{}
}

// Persist hypotetical persist method.
func (a *anyDBConnection) Persist(ctx context.Context, data map[string]interface{}) error {
	log.Println("level", "DEBUG", "msg", "storing new record", data)
	a.data = data
	return nil
}

// Close close any db connection.
func (a *anyDBConnection) Close() error {
	return nil
}
