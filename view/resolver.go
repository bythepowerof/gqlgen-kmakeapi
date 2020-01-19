package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *Resolver) Kmake() KmakeResolver {
	return &kmakeResolver{r}
}
func (r *Resolver) KmakeRun() KmakeRunResolver {
	return &kmakeRunResolver{r}
}
func (r *Resolver) KmakeRunDummy() KmakeRunDummyResolver {
	return &kmakeRunDummyResolver{r}
}
func (r *Resolver) KmakeRunJob() KmakeRunJobResolver {
	return &kmakeRunJobResolver{r}
}
func (r *Resolver) Namespace() NamespaceResolver {
	return &namespaceResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type kmakeResolver struct{ *Resolver }

type kmakeRunResolver struct{ *Resolver }

type kmakeRunDummyResolver struct{ *Resolver }

type kmakeRunJobResolver struct{ *Resolver }

type namespaceResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Kmakes(ctx context.Context, namespace *string) ([]*v1.Kmake, error) {
	panic("not implemented")
}
func (r *queryResolver) Kmakeruns(ctx context.Context, namespace *string, kmake *string) ([]*v1.KmakeRun, error) {
	panic("not implemented")
}
