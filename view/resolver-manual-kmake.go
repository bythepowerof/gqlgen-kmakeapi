package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
)

func (r *kmakeResolver) Variables(ctx context.Context, obj *v1.Kmake) ([]*v1.KV, error) {

	ret := make([]*v1.KV, 0)

	for k, v := range obj.Spec.Variables {
		ret = append(ret, &v1.KV{Key: k, Value: v})
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

func (r *kmakeResolver) Runs(ctx context.Context, obj *v1.Kmake, jobtype *controller.JobType, name *string) ([]*v1.KmakeRun, error) {
	namespace := obj.GetNamespace()
	kmakename := obj.GetName()

	return r.KmakeController.Kmakeruns(&namespace, &kmakename, jobtype, name)
}

func (r *kmakeResolver) UID(ctx context.Context, obj *v1.Kmake) (*string, error) {
	ret := string(obj.GetUID())
	return &ret, nil
}
