package gqlgen_todos

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
	v11 "k8s.io/api/core/v1"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *Resolver) Dummy() DummyResolver {
	return &dummyResolver{r}
}
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

type dummyResolver struct{ *Resolver }

func (r *dummyResolver) _(ctx context.Context, obj *Dummy) (string, error) {
	panic("not implemented")
}

type kmakeResolver struct{ *Resolver }

func (r *kmakeResolver) Status(ctx context.Context, obj *v1.Kmake) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

type namespaceResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Namespaces(ctx context.Context, name *string) ([]*v11.Namespace, error) {
	panic("not implemented")
}
