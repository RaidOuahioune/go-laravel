package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"fmt"

	"demo.com/hello/db"
	"demo.com/hello/models"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	//FOr now the user is being fetched from the mutation (but it should be fetched from the context using the JWT token)
	newTodo := &models.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	var db = (&db.Database{}).GetInstance()
	db.Create(newTodo)
	return newTodo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var db = (&db.Database{}).GetInstance()
	var todos []*models.Todo

	// preloading is very crucial in gorm? but not performant at all
	db.Preload("User").Find(&todos)

	if len(todos) == 0 {
		return nil, errors.New("no todos found")
	}
	return todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var db = (&db.Database{}).GetInstance()
	var users []*models.User
	db.Find(&users)

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}
	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var db = (&db.Database{}).GetInstance()
	var user models.User

	// u see u should eager load the todos ? otherwise they will be nulls
	db.Preload("Todos").First(&user, id)
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// ID implements TodoResolver.
func (r *todoResolver) ID(ctx context.Context, obj *models.Todo) (string, error) {
	return fmt.Sprint(obj.ID), nil
}

// ID implements UserResolver.
func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return fmt.Sprint(obj.ID), nil
}

// UserID is the resolver for the userId field.
func (r *newTodoResolver) UserID(ctx context.Context, obj *models.NewTodo, data string) error {
	var user models.User
	var db = (&db.Database{}).GetInstance()
	db.First(&user, obj.UserID)
	if user.ID == 0 {
		return fmt.Errorf("user not found with ID: %d", obj.UserID)
	}
	return nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// NewTodo returns NewTodoResolver implementation.
func (r *Resolver) NewTodo() NewTodoResolver { return &newTodoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type newTodoResolver struct{ *Resolver }
