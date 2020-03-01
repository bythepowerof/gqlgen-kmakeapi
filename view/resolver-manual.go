package gqlgen_kmakeapi

//go:generate go run github.com/99designs/gqlgen
import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/gql"
	v11 "k8s.io/api/core/v1"
	"github.com/bythepowerof/kmake-controller/api/v1"
)

type Resolver struct {
	KmakeController controller.KmakeController
}

func (r *queryResolver) Namespaces(ctx context.Context, name *string) ([]*v11.Namespace, error) {
	return r.KmakeController.Namespaces(ctx, name)
}

func (r *queryResolver) KmakeObjects(ctx context.Context, namespace string, name *string) ([]gql.KmakeObject, error) {
	ret := []gql.KmakeObject{}

	kms, _ := r.KmakeController.Kmakenowschedulers(ctx, namespace, name, nil)
	for _, v := range kms {
		ret = append(ret, v)
	}

	km, _ := r.KmakeController.Kmakes(ctx, &namespace, name)
	for _, v := range km {
		ret = append(ret, v)
	}

	kmr, _ := r.KmakeController.Kmakeruns(ctx, &namespace, nil, nil, name)
	for _, v := range kmr {
		ret = append(ret, v)
	}

	kmsr, _ := r.KmakeController.Kmakescheduleruns(ctx, namespace, nil, nil, nil, name, nil)

	for _, v := range kmsr {
		ret = append(ret, v)
	}

	return ret, nil
}

func (r *queryResolver) Kmakeruns(ctx context.Context, namespace string, kmake *string, jobtype *controller.JobType, kmakerun *string) ([]*v1.KmakeRun, error) {
	return r.KmakeController.Kmakeruns(ctx, &namespace, kmake, jobtype, kmakerun)
}

func (r *queryResolver) Kmakescheduleruns(ctx context.Context, namespace string, kmake *string, kmakerun *string, kmakescheduler *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {
	return r.KmakeController.Kmakescheduleruns(ctx, namespace, kmake, kmakerun, kmakescheduler, name, runtype)
}

func (r *queryResolver) Kmakeschedulers(ctx context.Context, namespace string, name *string, monitor *string) ([]gql.KmakeScheduler, error) {
	kms, _ := r.KmakeController.Kmakenowschedulers(ctx, namespace, name, monitor)
	ret := []gql.KmakeScheduler{}

	for _, v := range kms {
		ret = append(ret, v)
	}
	return ret, nil
}

func (r *queryResolver) Kmakes(ctx context.Context, namespace string, kmake *string) ([]*v1.Kmake, error) {
	return r.KmakeController.Kmakes(ctx, &namespace, kmake)
}
