package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	// v11 "k8s.io/api/core/v1"
	// 	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *queryResolver) Kmakescheduleruns(ctx context.Context, namespace string, kmake *string, kmakerun *string, kmakescheduler *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {
	return r.KmakeController.Kmakescheduleruns(ctx, namespace, kmake, kmakerun, kmakescheduler, name, runtype)
}

func (r *kmakeScheduleRunResolver) Kmakename(ctx context.Context, obj *v1.KmakeScheduleRun) (*string, error) {
	ret := obj.GetKmakeName()
	return &ret, nil
}
func (r *kmakeScheduleRunResolver) Kmakerunname(ctx context.Context, obj *v1.KmakeScheduleRun) (*string, error) {
	ret := obj.GetKmakeRunName()
	return &ret, nil
}
func (r *kmakeScheduleRunResolver) Kmakeschedulename(ctx context.Context, obj *v1.KmakeScheduleRun) (*string, error) {
	ret := obj.GetKmakeScheduleName()
	return &ret, nil
}
func (r *kmakeScheduleRunResolver) Status(ctx context.Context, obj *v1.KmakeScheduleRun) (string, error) {
	return obj.Status.Status, nil
}
func (r *kmakeScheduleRunResolver) Operation(ctx context.Context, obj *v1.KmakeScheduleRun) (*v1.KmakeScheduleRunOperation, error) {
	ret := obj.Spec.KmakeScheduleRunOperation

	return &ret, nil
}

func (r *kmakeScheduleCreateResolver) Dummy(ctx context.Context, obj *v1.KmakeScheduleCreate) (string, error) {
	return "1", nil
}

func (r *kmakeScheduleDeleteResolver) Dummy(ctx context.Context, obj *v1.KmakeScheduleDelete) (string, error) {
	return "1", nil
}

func (r *kmakeScheduleRunStartResolver) Dummy(ctx context.Context, obj *v1.KmakeScheduleRunStart) (string, error) {
	return "1", nil
}
