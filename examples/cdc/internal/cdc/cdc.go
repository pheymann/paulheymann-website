package cdc

import (
	"cdcexample/internal/movie"
)

type ClientDefinedContract struct {
	Name      string                  `yaml:"name"`
	View      string                  `yaml:"view"`
	Database  CDCDatabase             `yaml:"database"`
	CallChain []CDCRequestAndResponse `yaml:"callChain"`
}

type CDCDatabase struct {
	Movies     []movie.Movie `yaml:"movies"`
	Authorizer []string      `yaml:"authorizer"`
}

type CDCRequestAndResponse struct {
	Request  CDCRequest  `yaml:"request"`
	Response CDCResponse `yaml:"response"`
}

type CDCRequest struct {
	Uri          string        `yaml:"uri"`
	PathPameters []CDCKeyValue `yaml:"pathParameters"`
	Method       string        `yaml:"method"`
	Headers      []CDCKeyValue `yaml:"headers"`
	Body         *interface{}  `yaml:"body"`
}

type CDCKeyValue struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func (request CDCRequest) GetMapHeaders() map[string]string {
	headers := make(map[string]string)
	for _, header := range request.Headers {
		headers[header.Name] = header.Value
	}

	return headers
}

type CDCResponse struct {
	Type       string       `yaml:"type"`
	StatusCode int          `yaml:"status"`
	Body       *interface{} `yaml:"body"`
	ErrorBody  string       `yaml:"errorBody"`
}
