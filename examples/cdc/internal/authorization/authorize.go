package authorization

func (authorizer *ProductionAuthorizer) Authorize(bearerToken string) bool {
	_, ok := authorizer.db.Get(bearerToken)
	return ok
}
