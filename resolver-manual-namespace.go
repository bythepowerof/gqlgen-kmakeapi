package gqlgen_todos

// //go:generate go run github.com/99designs/gqlgen

import (
	context "context"
	v11 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *queryResolver) Namespaces(ctx context.Context, name *string) ([]*v11.Namespace, error) {

	ret := make([]*v11.Namespace, 0)

	if name != nil {
		ns := &v11.Namespace{}
		err := r.Client.Get(context.Background(), client.ObjectKey{
			Namespace: "",
			Name:      *name,
		}, ns)

		if err != nil {
			return nil, err
		}
		ret = append(ret, ns)
		return ret, nil
	}
	nsList := &v11.NamespaceList{}

	err := r.Client.List(context.Background(), nsList)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nsList.Items); i++ {
		ret = append(ret, &nsList.Items[i])
	}
	return ret, nil
}
