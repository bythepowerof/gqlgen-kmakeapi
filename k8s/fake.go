package k8s

import (
	bythepowerofv1 "github.com/bythepowerof/kmake-controller/api/v1"
	v11 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	k8sfakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

type FakeObjects struct{}

func (*FakeObjects) FakeNs() *v11.Namespace {
	return &v11.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ns1",
		},
	}
}

func (*FakeObjects) FakeKmake() *bythepowerofv1.Kmake {
	return &bythepowerofv1.Kmake{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-kmake",
			Namespace: "ns1",
		},
		Spec: bythepowerofv1.KmakeSpec{
			Variables: map[string]string{
				"VAR1": "Value1",
				"VAR2": "Value2",
			},
			Rules: []bythepowerofv1.KmakeRule{
				bythepowerofv1.KmakeRule{
					Targets:  []string{"Rule1"},
					Commands: []string{"@echo $@"},
				},
				bythepowerofv1.KmakeRule{
					Targets:  []string{"Rule2"},
					Commands: []string{"@echo $@"},
				},
			},
		},
	}
}

func (*FakeObjects) FakeKmakeRun() *bythepowerofv1.KmakeRun {
	return &bythepowerofv1.KmakeRun{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-kmake-run",
			Namespace: "ns1",
			Labels: map[string]string{
				"bythepowerof.github.io/kmake":     "test-kmake",
				"bythepowerof.github.io/scheduler": "now",
				"bythepowerof.github.io/workload":  "yes",
			},
		},
		Spec: bythepowerofv1.KmakeRunSpec{
			KmakeRunOperation: bythepowerofv1.KmakeRunOperation{
				Job: &bythepowerofv1.KmakeRunJob{
					Template: v11.PodTemplateSpec{
						Spec: v11.PodSpec{
							Containers: []v11.Container{
								v11.Container{
									Command: []string{"command text"},
									Image:   "image:latest",
									Args:    []string{"arg1", "arg2"},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (*FakeObjects) FakeKmakeNowScheduler() *bythepowerofv1.KmakeNowScheduler {
	return &bythepowerofv1.KmakeNowScheduler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-now-scheduler",
			Namespace: "ns1",
		},
		Spec: bythepowerofv1.KmakeNowSchedulerSpec{
			Monitor: []string{"now"},
		},
	}
}

func (*FakeObjects) FakeKmakeScheduleRun() *bythepowerofv1.KmakeScheduleRun {
	return &bythepowerofv1.KmakeScheduleRun{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-kmakeschedulerun",
			Namespace: "ns1",
			Labels: map[string]string{
				"bythepowerof.github.io/kmake":             "test-kmake",
				"bythepowerof.github.io/run":               "test-kmake-run",
				"bythepowerof.github.io/schedule-instance": "test-now-scheduler",
			},
		},
		Spec: bythepowerofv1.KmakeScheduleRunSpec{
			KmakeScheduleRunOperation: bythepowerofv1.KmakeScheduleRunOperation{
				Start: &bythepowerofv1.KmakeScheduleRunStart{},
			},
		},
	}
}

func (fo *FakeObjects) FakeK8sClient() (k8sclient.Client, error) {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = bythepowerofv1.AddToScheme(scheme)

	return k8sfakeclient.NewFakeClientWithScheme(scheme,
		fo.FakeNs(), fo.FakeKmake(), fo.FakeKmakeRun(), fo.FakeKmakeNowScheduler(), fo.FakeKmakeScheduleRun(),
	), nil
}
