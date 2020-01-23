package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/bythepowerof/kmake-controller/gql"
	// v11 "k8s.io/api/core/v1"
	// 	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *queryResolver) Kmakeschedulers(ctx context.Context, namespace string, name *string, monitor *string) ([]gql.KmakeScheduler, error) {
	kms, _ := r.KmakeController.Kmakenowschedulers(ctx, namespace, name, monitor)
	ret := []gql.KmakeScheduler{}

	for _, v := range kms {
		ret = append(ret, v)
	}
	return ret, nil
}

func (r *kmakeNowSchedulerResolver) Scheduleruns(ctx context.Context, obj *v1.KmakeNowScheduler, kmake *string, kmakerun *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {
	schedulename := obj.GetName()
	return r.KmakeController.Kmakescheduleruns(ctx, obj.GetNamespace(), kmake, kmakerun, &schedulename, name, runtype)
}
