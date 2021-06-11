package web

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Printf("encoding response %+v", response)
	return json.NewEncoder(w).Encode(response)
}
