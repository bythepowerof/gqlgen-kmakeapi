package gqlgen_kmakeapi

import (
	context "context"
	v11 "k8s.io/api/core/v1"
)

func (r *queryResolver) Namespaces(ctx context.Context, name *string) ([]*v11.Namespace, error) {

	return r.KmakeController.Namespaces(ctx, name)
}
