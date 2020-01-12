package gqlgen_todos

// //go:generate go run github.com/99designs/gqlgen

import (
	context "context"
	"fmt"
	"math/rand"
)

type Resolver struct {
	todos []*Todo
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo := &Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

