package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/bythepowerof/kmake-controller/gql"
)

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

func (r *kmakeScheduleRunResolver) UID(ctx context.Context, obj *v1.KmakeScheduleRun) (*string, error) {
	ret := string(obj.GetUID())
	return &ret, nil
}
