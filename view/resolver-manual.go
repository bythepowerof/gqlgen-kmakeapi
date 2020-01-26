package gqlgen_kmakeapi

//go:generate go run github.com/99designs/gqlgen
import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/gql"
)

type Resolver struct {
	KmakeController controller.KmakeController
}

func (r *queryResolver) KmakeObjects(ctx context.Context, namespace string, name *string) ([]gql.KmakeObject, error) {
	ret := []gql.KmakeObject{}

	kms, _ := r.KmakeController.Kmakenowschedulers(ctx, namespace, name, nil)
	for _, v := range kms {
		ret = append(ret, v)
	}

	km, _ := r.KmakeController.Kmakes(ctx, &namespace, name)
	for _, v := range km {
		ret = append(ret, v)
	}

	kmr, _ := r.KmakeController.Kmakeruns(ctx, &namespace, nil, nil, name)
	for _, v := range kmr {
		ret = append(ret, v)
	}

	kmsr, _ := r.KmakeController.Kmakescheduleruns(ctx, namespace, nil, nil, nil, name, nil)

	for _, v := range kmsr {
		ret = append(ret, v)
	}

	return ret, nil
}
