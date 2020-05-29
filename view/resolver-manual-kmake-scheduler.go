package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
)

func (r *kmakeNowSchedulerResolver) Scheduleruns(ctx context.Context, obj *v1.KmakeNowScheduler, kmake *string, kmakerun *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {
	schedulename := obj.GetName()
	return r.KmakeController.Kmakescheduleruns(obj.GetNamespace(), kmake, kmakerun, &schedulename, name, runtype)
}

func (r *kmakeNowSchedulerResolver) UID(ctx context.Context, obj *v1.KmakeNowScheduler) (*string, error) {
	ret := string(obj.GetUID())
	return &ret, nil
}
