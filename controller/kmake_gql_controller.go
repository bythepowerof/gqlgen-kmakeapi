package controller

import (
	context "context"
	"github.com/bythepowerof/kmake-controller/api/v1"
	v11 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type JobType string

const (
	JobTypeJob      JobType = "JOB"
	JobTypeDummy    JobType = "DUMMY"
	JobTypeFilewait JobType = "FILEWAIT"
)

type KmakeController interface {
	Namespaces(ctx context.Context, name *string) ([]*v11.Namespace, error)
	Kmakes(ctx context.Context, namespace *string, name *string) ([]*v1.Kmake, error)
	KmakeRuns(ctx context.Context, namespace *string, kmakename *string, jobtype *JobType, name *string) ([]*v1.KmakeRun, error)
}

type KubernetesController struct {
	Client client.Client
}

func (r *KubernetesController) Namespaces(ctx context.Context, name *string) ([]*v11.Namespace, error) {

	ret := make([]*v11.Namespace, 0)

	nsList := &v11.NamespaceList{}
	o := &client.ListOptions{}

	if name != nil {
		fields := map[string]string{"metadata.name": *name}
		client.MatchingFields(fields).ApplyToList(o)
	}

	err := r.Client.List(context.Background(), nsList, o)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nsList.Items); i++ {
		ret = append(ret, &nsList.Items[i])
	}
	return ret, nil
}

func (r *KubernetesController) Kmakes(ctx context.Context, namespace *string, name *string) ([]*v1.Kmake, error) {
	ret := make([]*v1.Kmake, 0)

	kmakeList := &v1.KmakeList{}
	o := &client.ListOptions{}
	client.InNamespace(*namespace).ApplyToList(o)

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

func (r *KubernetesController) KmakeRuns(ctx context.Context, namespace *string, kmakename *string, jobtype *JobType, name *string) ([]*v1.KmakeRun, error) {
	ret := make([]*v1.KmakeRun, 0)

	kmakerunList := &v1.KmakeRunList{}

	labels := map[string]string{"bythepowerof.github.io/kmake": *kmakename}

	o := &client.ListOptions{}
	client.InNamespace(*namespace).ApplyToList(o)
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
