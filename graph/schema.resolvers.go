package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/irdaislakhuafa/learn-graphql-go/graph/generated"
	"github.com/irdaislakhuafa/learn-graphql-go/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, newTodo model.NewTodo) (*model.Todo, error) {

	defer func() {
		err := recover()
		if err != nil {
			log.Println("Error while saving new todo :", err)
		}
	}()
	todo := &model.Todo{
		Text: newTodo.Text,
		Done: false,
		User: r.UserRepository.FindById(&newTodo.UserID),
	}

	_, err := r.TodoRepository.Save(todo)

	if err != nil {
		return nil, err
	} else {
		return todo, nil
	}
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("Error while finad all todos :", err)
		}
	}()

	todoDto, err := r.TodoRepository.FindAll()

	todos := make([]*model.Todo, 0)
	for _, value := range todoDto {
		todo := r.TodoRepository.MapToEntity(value)
		todos = append(todos, &todo)
	}
	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
