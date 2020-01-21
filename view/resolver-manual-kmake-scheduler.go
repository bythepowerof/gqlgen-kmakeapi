package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	// v11 "k8s.io/api/core/v1"
	// 	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *queryResolver) Kmakenowschedulers(ctx context.Context, namespace string, name *string, monitor *string) ([]*v1.KmakeNowScheduler, error) {
	return r.KmakeController.Kmakenowschedulers(ctx, namespace, name, monitor)
}

func (r *kmakeNowSchedulerResolver) Variables(ctx context.Context, obj *v1.KmakeNowScheduler) ([]*controller.KV, error) {
	ret := make([]*controller.KV, 0)

	for k, v := range obj.Spec.Variables {
		ret = append(ret, &controller.KV{Key: k, Value: v})
	}
	return ret, nil
}

func (r *kmakeNowSchedulerResolver) Monitor(ctx context.Context, obj *v1.KmakeNowScheduler) ([]string, error) {
	return obj.Spec.Monitor, nil
}

func (r *kmakeNowSchedulerResolver) Scheduleruns(ctx context.Context, obj *v1.KmakeNowScheduler, kmake *string, kmakerun *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {
	schedulename := obj.GetName()
	return r.KmakeController.Kmakescheduleruns(ctx, obj.GetNamespace(), kmake, kmakerun, &schedulename, name, runtype)
}
