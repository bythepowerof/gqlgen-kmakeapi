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

func FakeK8sClient() (k8sclient.Client, error) {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = bythepowerofv1.AddToScheme(scheme)

	ns := &v11.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ns1",
		},
	}

	kmake := &bythepowerofv1.Kmake{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-kmake",
			Namespace: "ns1",
		},
	}

	kmakerun := &bythepowerofv1.KmakeRun{
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
				Job: &bythepowerofv1.KmakeRunJob{},
			},
		},
	}

	kmakescheduler := &bythepowerofv1.KmakeNowScheduler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-now-scheduler",
			Namespace: "ns1",
		},
		Spec: bythepowerofv1.KmakeNowSchedulerSpec{
			Monitor: []string{"now"},
		},
	}

	kmakeschedulerun := &bythepowerofv1.KmakeScheduleRun{
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

	return k8sfakeclient.NewFakeClientWithScheme(scheme,
		ns, kmake, kmakerun, kmakescheduler, kmakeschedulerun,
	), nil
}
