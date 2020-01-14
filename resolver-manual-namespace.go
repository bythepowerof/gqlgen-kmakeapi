package gqlgen_todos

// //go:generate go run github.com/99designs/gqlgen

import (
	context "context"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *queryResolver) Namespaces(ctx context.Context, name *string) ([]*corev1.Namespace, error) {

	ret := make([]*corev1.Namespace,0)

	if name != nil {
		ns := &corev1.Namespace{}
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
	nsList := &corev1.NamespaceList{}

	err := r.Client.List(context.Background(), nsList)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nsList.Items); i++ {
		ret = append(ret, &nsList.Items[i])
	}
	return ret, nil
}
