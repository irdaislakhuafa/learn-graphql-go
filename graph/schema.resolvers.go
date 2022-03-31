package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/irdaislakhuafa/learn-graphql-go/graph/generated"
	"github.com/irdaislakhuafa/learn-graphql-go/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, in model.NewTodo) (*model.Todo, error) {
	defer func() {
		err := recover()
		if err != nil {
			panic(fmt.Errorf("not implemented"))
		}
	}()

	// mapping model.NewTodo to model.Todo
	todo := &model.Todo{
		ID:   strconv.Itoa(rand.Int()),
		Text: in.Text,
		User: &model.User{
			ID:   in.UserID,
			Name: "User " + in.UserID,
		},
	}

	// add new model.Todo to slice `todos` (InMemoryDatabase)
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	defer func() {
		err := recover()
		if err != nil {
			panic(fmt.Errorf("not implemented"))
		}
	}()

	// return list of todos from InMemoryDatabase
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
