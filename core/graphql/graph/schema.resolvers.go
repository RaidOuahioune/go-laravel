package graph

import (
	"context"
	"errors"

	"demo.com/hello/core/graphql/graph/model"
	"github.com/google/uuid"
)

// In-memory store for the todos
var todos = []*model.Todo{}

// CreateTodo is the resolver for the createTodo field.

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	newTodo := &model.Todo{
		ID:   uuid.New().String(),
		Text: input.Text,
		Done: false,
		User: &model.User{ID: input.UserID, Name: "user1"},
	}
	todos = append(todos, newTodo)
	return newTodo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	if len(todos) == 0 {
		return nil, errors.New("no todos found")
	}
	return todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
