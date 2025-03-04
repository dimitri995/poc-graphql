package graph

import (
	"poc-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	products []*model.Product
	users    []*model.User
}
