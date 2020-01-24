package gqlgen_kmakeapi

import (
	"context"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/bythepowerof/kmake-controller/gql"
	// v11 "k8s.io/api/core/v1"
	// 	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *kmakeRunResolver) Status(ctx context.Context, obj *v1.KmakeRun) (string, error) {
	return obj.Status.Status, nil
}

// func (r *kmakeRunResolver) Operation(ctx context.Context, obj *v1.KmakeRun) (*v1.KmakeRunOperation, error) {
// 	return &obj.Spec.KmakeRunOperation, nil
// }

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
func (r *kmakeRunResolver) Kmakename(ctx context.Context, obj *v1.KmakeRun) (*string, error) {
	kmakename := obj.GetKmakeName()
	return &kmakename, nil
}
func (r *kmakeRunJobResolver) Args(ctx context.Context, obj *v1.KmakeRunJob) ([]*string, error) {
	ret := make([]*string, 0)

	for i := 0; i < len(obj.Template.Spec.Containers[0].Args); i++ {
		ret = append(ret, &obj.Template.Spec.Containers[0].Args[i])
	}
	return ret, nil
}

func (r *queryResolver) Kmakeruns(ctx context.Context, namespace string, kmake *string, jobtype *controller.JobType, kmakerun *string) ([]*v1.KmakeRun, error) {
	return r.KmakeController.Kmakeruns(ctx, &namespace, kmake, jobtype, kmakerun)
}

func (r *kmakeRunResolver) Schedulerun(ctx context.Context, obj *v1.KmakeRun, kmakescheduler *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {
	kmake := obj.GetKmakeName()
	kmakerun := obj.GetName()
	return r.KmakeController.Kmakescheduleruns(ctx, obj.GetNamespace(), &kmake, &kmakerun, kmakescheduler, name, runtype)
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
