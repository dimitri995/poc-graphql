package graph

import (
	"poc-graphql/graph/model"
	"poc-graphql/pkg/config"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users  []*model.User
	config config.Config

	UserClient *UserClient
}
