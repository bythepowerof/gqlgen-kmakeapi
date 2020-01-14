package gqlgen_todos

import (
	"context"

	"k8s.io/api/core/v1"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Namespace() NamespaceResolver {
	return &namespaceResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
// 	panic("not implemented")
// }

type namespaceResolver struct{ *Resolver }

func (r *namespaceResolver) Kmakes(ctx context.Context, obj *v1.Namespace, name *string) ([]*Kmake, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

// func (r *queryResolver) Todos(ctx context.Context, id *string) ([]*Todo, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) Namespaces(ctx context.Context, name *string) ([]*v1.Namespace, error) {
// 	panic("not implemented")
// }

type todoResolver struct{ *Resolver }

// func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
// 	panic("not implemented")
// }
