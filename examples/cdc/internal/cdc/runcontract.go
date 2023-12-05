package cdc

import (
	"cdcexample/internal/authorization"
	"cdcexample/internal/movie"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	cdcutil "cdcexample/internal/cdc/util"
)

func RunContracts(
	t *testing.T,
	cdcPath string,
	eval func(
		r *http.Request,
		w http.ResponseWriter,
		movieService movie.MovieService,
		authorizer authorization.Authorizer,
	),
) {
	contract := MustLoadCDC(cdcPath)

	t.Logf("Running contract test case: %s", contract.Name)

	movieServiceMock := &mockMovieService{movies: contract.Database.Movies}
	authorizerMock := &mockAuthorizer{validTokens: contract.Database.Authorizer}

	for _, requestAndResponse := range contract.CallChain {
		var bodyReader io.Reader

		if requestAndResponse.Request.Body != nil {
			bodyBytes, err := json.Marshal(cdcutil.YamlToJson(*requestAndResponse.Request.Body))
			if err != nil {
				panic(err)
			}

			bodyReader = strings.NewReader(string(bodyBytes))
		}

		request := httptest.NewRequest(
			requestAndResponse.Request.Method,
			requestAndResponse.Request.Uri,
			bodyReader,
		)

		for _, keyValue := range requestAndResponse.Request.Headers {
			request.Header.Set(keyValue.Name, keyValue.Value)
		}

		responseWriter := httptest.NewRecorder()

		eval(request, responseWriter, movieServiceMock, authorizerMock)

		response := responseWriter.Result()
		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		assert.Equal(t, requestAndResponse.Response.StatusCode, response.StatusCode)
		if requestAndResponse.Response.Type == "success" && requestAndResponse.Response.Body != nil {
			cdcutil.AssertEqualInterface(t, *requestAndResponse.Response.Body, body)
		} else if requestAndResponse.Response.Type == "error" {
			assert.Equal(t, requestAndResponse.Response.ErrorBody, string(body), "error body does not match")
		}
	}
}

type mockMovieService struct {
	movies []movie.Movie
}

func (m *mockMovieService) ListAll() ([]movie.Movie, error) {
	return m.movies, nil
}

func (m *mockMovieService) Store(movie movie.Movie) error {
	m.movies = append(m.movies, movie)
	return nil
}

type mockAuthorizer struct {
	validTokens []string
}

func (m *mockAuthorizer) Authorize(bearerToken string) bool {
	for _, validToken := range m.validTokens {
		if bearerToken == validToken {
			return true
		}
	}

	return false
}
