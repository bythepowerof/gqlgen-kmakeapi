package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/bythepowerof/kmake-controller/gql"
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
	var ret string
	if scheduler, ok := obj.GetLabels()["bythepowerof.github.io/schedule-instance"]; ok {
		ret = scheduler
	} else {
		ret = ""
	}
	return &ret, nil
}
func (r *kmakeScheduleRunResolver) Status(ctx context.Context, obj *v1.KmakeScheduleRun) (string, error) {
	return obj.Status.Status, nil
}

func (r *kmakeScheduleRunResolver) Operation(ctx context.Context, obj *v1.KmakeScheduleRun) (gql.KmakeScheduleRunOperation, error) {
	if obj.Spec.KmakeScheduleRunOperation.Start != nil {
		return obj.Spec.KmakeScheduleRunOperation.Start, nil
	}
	if obj.Spec.KmakeScheduleRunOperation.Restart != nil {
		return obj.Spec.KmakeScheduleRunOperation.Restart, nil
	}
	if obj.Spec.KmakeScheduleRunOperation.Stop != nil {
		return obj.Spec.KmakeScheduleRunOperation.Stop, nil
	}
	if obj.Spec.KmakeScheduleRunOperation.Delete != nil {
		return obj.Spec.KmakeScheduleRunOperation.Delete, nil
	}
	if obj.Spec.KmakeScheduleRunOperation.Create != nil {
		return obj.Spec.KmakeScheduleRunOperation.Create, nil
	}
	if obj.Spec.KmakeScheduleRunOperation.Reset != nil {
		return obj.Spec.KmakeScheduleRunOperation.Reset, nil
	}
	if obj.Spec.KmakeScheduleRunOperation.Force != nil {
		return obj.Spec.KmakeScheduleRunOperation.Force, nil
	}
	return nil, nil
}
