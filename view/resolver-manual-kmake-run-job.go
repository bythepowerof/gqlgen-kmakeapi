package gqlgen_kmakeapi

import (
	"context"

	// "github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	// "github.com/bythepowerof/kmake-controller/gql"
)

func (r *kmakeRunJobResolver) Image(ctx context.Context, obj *v1.KmakeRunJob) (string, error) {
	return obj.Template.Spec.Containers[0].Image, nil
}

func (r *kmakeRunJobResolver) Command(ctx context.Context, obj *v1.KmakeRunJob) ([]*string, error) {
	ret := make([]*string, 0)

	for i := 0; i < len(obj.Template.Spec.Containers[0].Command); i++ {
		ret = append(ret, &obj.Template.Spec.Containers[0].Command[i])
	}

	return ret, nil
}

func (r *kmakeRunJobResolver) Args(ctx context.Context, obj *v1.KmakeRunJob) ([]*string, error) {
	ret := make([]*string, 0)

	for i := 0; i < len(obj.Template.Spec.Containers[0].Args); i++ {
		ret = append(ret, &obj.Template.Spec.Containers[0].Args[i])
	}
	return ret, nil
}
