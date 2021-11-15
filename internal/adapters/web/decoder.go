package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fernandoocampo/hexagonal-template-go/internal/people"
)

func decodeCreatePersonRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("level", "DEBUG", "msg", "decoding new person request")
	var req CreatePersonRequest
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("level", "ERROR", "new person request could not be decoded. Request: %q because of: %s", string(body), err.Error())
		return nil, err
	}

	log.Println("level", "DEBUG", "msg", "person request was decoded", "request", req)

	newPerson := people.NewPerson{
		Name: req.JName,
	}

	return newPerson, nil
}
