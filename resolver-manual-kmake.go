package gqlgen_todos

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
	v11 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *namespaceResolver) Kmakes(ctx context.Context, obj *v11.Namespace, name *string) ([]*v1.Kmake, error) {
	ret := make([]*v1.Kmake, 0)

	kmakeList := &v1.KmakeList{}
	o := &client.ListOptions{}
	client.InNamespace(obj.GetName()).ApplyToList(o)

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	err := r.Client.List(context.Background(), kmakeList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(kmakeList.Items); i++ {
		ret = append(ret, &kmakeList.Items[i])
	}
	return ret, nil
}

func (r *kmakeResolver) Variables(ctx context.Context, obj *v1.Kmake) ([]*Variable, error) {

	ret := make([]*Variable, 0)

	for k, v := range obj.Spec.Variables {
		ret = append(ret, &Variable{Name: k, Value: v})
	}
	return ret, nil
}

func (r *kmakeResolver) Rules(ctx context.Context, obj *v1.Kmake) ([]*v1.KmakeRule, error) {
	ret := make([]*v1.KmakeRule, 0)

	for _, v := range obj.Spec.Rules {
		ret = append(ret, &v)
	}

	return ret, nil
}

func (r *kmakeResolver) Status(ctx context.Context, obj *v1.Kmake) (string, error) {
	return obj.Status.Status, nil
}

func (r *kmakeResolver) Runs(ctx context.Context, obj *v1.Kmake, jobtype *JobType, name *string) ([]*v1.KmakeRun, error) {
	ret := make([]*v1.KmakeRun, 0)

	kmakerunList := &v1.KmakeRunList{}

	labels := map[string]string{"bythepowerof.github.io/kmake": obj.GetName()}

	o := &client.ListOptions{}
	client.InNamespace(obj.GetNamespace()).ApplyToList(o)
	client.MatchingLabels(labels).ApplyToList(o)

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	err := r.Client.List(context.Background(), kmakerunList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(kmakerunList.Items); i++ {
		if jobtype != nil {
			if *jobtype == "DUMMY" && kmakerunList.Items[i].Spec.KmakeRunOperation.Dummy == nil {
				continue
			}
			if *jobtype == "JOB" && kmakerunList.Items[i].Spec.KmakeRunOperation.Job == nil {
				continue
			}
			if *jobtype == "FILEWAIT" && kmakerunList.Items[i].Spec.KmakeRunOperation.FileWait == nil {
				continue
			}
		}
		ret = append(ret, &kmakerunList.Items[i])
	}
	return ret, nil
}
