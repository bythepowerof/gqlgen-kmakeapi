package gqlgen_todos

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
	// v11 "k8s.io/api/core/v1"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// type Resolver struct{}

func (r *Resolver) Kmake() KmakeResolver {
	return &kmakeResolver{r}
}
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

type kmakeResolver struct{ *Resolver }

// func (r *kmakeResolver) Variables(ctx context.Context, obj *v1.Kmake) ([]*Variable, error) {
// 	panic("not implemented")
// }
func (r *kmakeResolver) Rules(ctx context.Context, obj *v1.Kmake) ([]*Rule, error) {
	panic("not implemented")
}
func (r *kmakeResolver) Status(ctx context.Context, obj *v1.Kmake) (*string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
// 	panic("not implemented")
// }

type namespaceResolver struct{ *Resolver }

// func (r *namespaceResolver) Kmakes(ctx context.Context, obj *v11.Namespace, name *string) ([]*v1.Kmake, error) {
// 	panic("not implemented")
// }

type queryResolver struct{ *Resolver }

// func (r *queryResolver) Todos(ctx context.Context, id *string) ([]*Todo, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) Namespaces(ctx context.Context, name *string) ([]*v11.Namespace, error) {
// 	panic("not implemented")
// }

type todoResolver struct{ *Resolver }

// func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
// 	panic("not implemented")
// }
