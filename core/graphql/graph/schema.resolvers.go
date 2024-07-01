package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"fmt"

	"demo.com/hello/core/graphql/graph/model"
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

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input *model.UpdateTodo) (*models.Todo, error) {
	var db = (&db.Database{}).GetInstance()

	var columnsMap = map[string]interface{}{}
	if input.Text != nil {
		columnsMap["text"] = *input.Text

	}
	if input.Done != nil {
		columnsMap["done"] = *input.Done

	}

	var todo models.Todo
	db.First(&todo, input.ID)
	if todo.ID == 0 {
		return nil, errors.New("todo not found")
	}

	db.Model(&todo).Updates(columnsMap)

	return &todo, nil

}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*models.Todo, error) {
	var db = (&db.Database{}).GetInstance()
	var todo models.Todo
	db.First(&todo, id)
	if todo.ID == 0 {
		return nil, errors.New("todo not found")
	}
	db.Delete(&todo, id)
	return &todo, nil
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

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	var db = (&db.Database{}).GetInstance()
	var todo models.Todo

	db.First(&todo, id)
	if todo.ID == 0 {
		return nil, errors.New("todo not found")
	}
	return &todo, nil
}

// ID implements TodoResolver.
func (r *todoResolver) ID(ctx context.Context, obj *models.Todo) (string, error) {
	return fmt.Sprint(obj.ID), nil
}

// ID implements UserResolver.
func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return fmt.Sprint(obj.ID), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *newTodoResolver) UserID(ctx context.Context, obj *models.NewTodo, data string) error {
	var user models.User
	var db = (&db.Database{}).GetInstance()
	db.First(&user, obj.UserID)
	if user.ID == 0 {
		return fmt.Errorf("user not found with ID: %d", obj.Text)
	}
	return nil
}

type newTodoResolver struct{ *Resolver }

func (r *Resolver) NewTodo() *newTodoResolver { return &newTodoResolver{r} }
