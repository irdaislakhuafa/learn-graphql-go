package graph

import (
	"github.com/irdaislakhuafa/learn-graphql-go/graph/model"
	"github.com/irdaislakhuafa/learn-graphql-go/src/repositories"
	"github.com/jmoiron/sqlx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos          []*model.Todo
	UserRepository *repositories.UserRepository
	TodoRepository *repositories.TodoRepository
	Connection     *sqlx.DB
}
