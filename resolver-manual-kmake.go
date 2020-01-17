package gqlgen_todos

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
	v11 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *namespaceResolver) Kmakes(ctx context.Context, obj *v11.Namespace, name *string) ([]*v1.Kmake, error) {
	ret := make([]*v1.Kmake, 0)

	if name != nil {
		kmake := &v1.Kmake{}
		err := r.Client.Get(context.Background(), client.ObjectKey{
			Namespace: obj.GetName(),
			Name:      *name,
		}, kmake)

		if err != nil {
			return nil, err
		}
		ret = append(ret, kmake)
		return ret, nil
	}
	kmakeList := &v1.KmakeList{}

	err := r.Client.List(context.Background(), kmakeList, client.InNamespace(obj.GetNamespace()))
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(kmakeList.Items); i++ {
		ret = append(ret, &kmakeList.Items[i])
	}
	return ret, nil
}

func (r *kmakeResolver) Variables(ctx context.Context, obj *v1.Kmake) ([]*Variable, error) {

	ret := make([]*Variable, 0)

	for k, v := range obj.Spec.Variables {
		ret = append(ret, &Variable{Name: k, Value: v})
	}
	return ret, nil
}

func (r *kmakeResolver) Rules(ctx context.Context, obj *v1.Kmake) ([]*v1.KmakeRule, error) {
	ret := make([]*v1.KmakeRule, 0)

	for _, v := range obj.Spec.Rules {
		ret = append(ret, &v)
	}

	return ret, nil
}

func (r *kmakeResolver) Status(ctx context.Context, obj *v1.Kmake) (*string, error) {
	return &obj.Status.Status, nil
}
