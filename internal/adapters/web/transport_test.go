package web_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fernandoocampo/hexagonal-template-go/internal/adapters/web"
	"github.com/go-kit/kit/endpoint"
	"github.com/stretchr/testify/assert"
)

func TestCreatePersonSuccessfuly(t *testing.T) {
	// GIVEN
	givenPersonJSON := `{"name": "Enrique"}`
	expectedResponse := "true"
	mockedEndpoints := &anyEndpoints{}
	httpHandler := web.NewHTTPServer(mockedEndpoints)

	request := httptest.NewRequest(http.MethodPost, "/people", strings.NewReader(givenPersonJSON))
	response := httptest.NewRecorder()

	// WHEN
	httpHandler.ServeHTTP(response, request)

	var createPersonResponse string
	err := json.NewDecoder(response.Body).Decode(&createPersonResponse)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		t.FailNow()
	}

	assert.Equal(t, createPersonResponse, expectedResponse)
	assert.Equal(t, response.Code, http.StatusOK)

}

// anyEndpoints is a hypothetical people endpoints.
type anyEndpoints struct{}

func (a *anyEndpoints) CreatePerson() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "true", nil
	}
}
