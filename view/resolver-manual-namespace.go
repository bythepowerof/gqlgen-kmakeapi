package gqlgen_kmakeapi

import (
	context "context"
	"github.com/bythepowerof/kmake-controller/api/v1"
	v11 "k8s.io/api/core/v1"
)

func (r *namespaceResolver) Kmakes(ctx context.Context, obj *v11.Namespace, name *string) ([]*v1.Kmake, error) {
	namespace := obj.GetName()
	return r.KmakeController.Kmakes(&namespace, name)
}
