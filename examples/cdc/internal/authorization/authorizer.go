package authorization

import "cdcexample/internal/inmemorydb"

type Authorizer interface {
	Authorize(bearerToken string) bool
}

type ProductionAuthorizer struct {
	db inmemorydb.InMemoryDB[string]
}

func NewProductionAuthorizer(validTokens []string) *ProductionAuthorizer {
	// should be something like AWS Cognito but for the sake of simplicity we use
	// an in-memory DB
	db := inmemorydb.NewInMemoryDB(validTokens, func(token string) string { return token })
	return &ProductionAuthorizer{db: db}
}
