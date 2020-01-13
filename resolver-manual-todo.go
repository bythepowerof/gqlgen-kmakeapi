package gqlgen_todos

// //go:generate go run github.com/99designs/gqlgen

import (
	context "context"
	"fmt"
	"math/rand"

	// "k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/rest"
)

type Resolver struct {
	todos []*Todo
	Clientset *kubernetes.Clientset
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

func (r *queryResolver) Todos(ctx context.Context, id *string) ([]*Todo, error) {
	if id == nil {
		return r.todos, nil
	} 
	b := r.todos[:0]
	for _, x := range r.todos {
		if x.UserID == *id {
			b = append(b, x)
		}
	}
	return b, nil
}

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

