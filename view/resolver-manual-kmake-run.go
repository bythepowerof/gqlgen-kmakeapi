package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/bythepowerof/kmake-controller/gql"
)

func (r *kmakeRunResolver) Kmakename(ctx context.Context, obj *v1.KmakeRun) (*string, error) {
	kmakename := obj.GetKmakeName()
	return &kmakename, nil
}

func (r *kmakeRunResolver) Schedulerun(ctx context.Context, obj *v1.KmakeRun, kmakescheduler *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {
	kmake := obj.GetKmakeName()
	kmakerun := obj.GetName()
	return r.KmakeController.Kmakescheduleruns(obj.GetNamespace(), &kmake, &kmakerun, kmakescheduler, name, runtype)
}

func (r *kmakeRunResolver) Operation(ctx context.Context, obj *v1.KmakeRun) (gql.KmakeRunOperation, error) {

	if obj.Spec.KmakeRunOperation.Job != nil {
		return obj.Spec.KmakeRunOperation.Job, nil
	}
	if obj.Spec.KmakeRunOperation.Dummy != nil {
		return obj.Spec.KmakeRunOperation.Dummy, nil
	}
	if obj.Spec.KmakeRunOperation.FileWait != nil {
		return obj.Spec.KmakeRunOperation.FileWait, nil
	}
	return nil, nil
}

func (r *kmakeRunResolver) UID(ctx context.Context, obj *v1.KmakeRun) (*string, error) {
	ret := string(obj.GetUID())
	return &ret, nil
}
