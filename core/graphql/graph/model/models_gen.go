// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Query struct {
}

type UpdateTodo struct {
	ID   string  `json:"id"`
	Text *string `json:"text,omitempty"`
	Done *bool   `json:"done,omitempty"`
}