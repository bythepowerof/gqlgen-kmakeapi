package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
	// v11 "k8s.io/api/core/v1"
	// 	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *kmakeRunResolver) Runstatus(ctx context.Context, obj *v1.KmakeRun) (*v1.KmakeRunStatus, error) {
	return &obj.Status, nil
}

func (r *kmakeRunResolver) Operation(ctx context.Context, obj *v1.KmakeRun) (*v1.KmakeRunOperation, error) {
	return &obj.Spec.KmakeRunOperation, nil
}

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

func (r *kmakeRunDummyResolver) Dummy(ctx context.Context, obj *v1.KmakeRunDummy) (string, error) {
	return "1", nil
}
