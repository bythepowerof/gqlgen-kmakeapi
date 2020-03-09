package controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/gqlgen-kmakeapi/k8s"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Controller", func() {
	var k k8sclient.Client
	var kmc *KubernetesController
	var fo *k8s.FakeObjects

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		kmc = &KubernetesController{Client: k}
	})

	Describe("KubernetesController", func() {
		Context("implements interface", func() {
			It("KmakeController", func() {
				v := KubernetesController{}
				var i interface{} = v
				_, ok := i.(KmakeController)
				Expect(ok).To(Equal(false))

				var p interface{} = &v
				_, ok = p.(KmakeController)
				Expect(ok).To(Equal(true))
			})
		})
		Context("fetches", func() {
			It("namespace", func() {
				ns := "ns1"
				namespaces, err := kmc.Namespaces(&ns)
				Expect(err).To(BeNil())
				Expect(namespaces[0].GetName()).To(Equal(ns))
				Expect(len(namespaces)).To(Equal(1))
			})
			It("kmake", func() {
				ns := "ns1"
				n := "test-kmake"
				kmakes, err := kmc.Kmakes(&ns, &n)
				Expect(err).To(BeNil())
				Expect(kmakes[0].GetName()).To(Equal(n))
				Expect(kmakes[0].GetNamespace()).To(Equal(ns))
				Expect(len(kmakes)).To(Equal(1))
			})
			It("kmakerun", func() {
				ns := "ns1"
				n := "test-kmake-run"
				km := "test-kmake"
				jt := JobTypeJob
				kmakeruns, err := kmc.Kmakeruns(&ns, &km, &jt, &n)

				Expect(err).To(BeNil())
				Expect(kmakeruns[0].GetName()).To(Equal(n))
				Expect(kmakeruns[0].GetNamespace()).To(Equal(ns))
				Expect(len(kmakeruns)).To(Equal(1))
			})
			It("now scheduler", func() {
				ns := "ns1"
				n := "test-now-scheduler"
				mon := "now"
				schedulers, err := kmc.Kmakenowschedulers(ns, &n, &mon)
				Expect(err).To(BeNil())
				Expect(schedulers[0].GetName()).To(Equal(n))
				Expect(schedulers[0].GetNamespace()).To(Equal(ns))
				Expect(len(schedulers)).To(Equal(1))
			})
			It("kmakeschedulerun", func() {
				ns := "ns1"
				n := "test-kmakeschedulerun"
				km := "test-kmake"
				kmr := "test-kmake-run"
				sched := "test-now-scheduler"
				rt := RunTypeStart

				kmsr, err := kmc.Kmakescheduleruns(ns, &km, &kmr, &sched, &n, &rt)

				Expect(err).To(BeNil())
				Expect(kmsr[0].GetName()).To(Equal(n))
				Expect(kmsr[0].GetNamespace()).To(Equal(ns))
				Expect(len(kmsr)).To(Equal(1))
			})
		})
	})

	// these have to separate as GenerateName does not work for fake clients
	Describe("resetting a schedule", func() {
		Context("resetting a scheduler", func() {
			It("create schedule run", func() {
				ns := "ns1"
				sched := "test-now-scheduler"
				rt := RunTypeReset

				By("resetting a scheduler")
				kmsr, err := kmc.CreateScheduleRun(ns, nil, nil, &sched, &rt, map[string]string{"full": "true"})
				Expect(err).To(BeNil())
				Expect(kmsr.GetNamespace()).To(Equal(ns))
			})
		})
	})
	Describe("stopping a run", func() {
		Context("stopping a run", func() {
			It("stopping a run", func() {
				ns := "ns1"
				sched := "test-now-scheduler"
				rt := RunTypeStop
				kmr := "test-kmake-run"

				By("stopping a run")
				kmsr2, err := kmc.CreateScheduleRun(ns, nil, &kmr, &sched, &rt, nil)
				Expect(err).To(BeNil())
				Expect(kmsr2.GetNamespace()).To(Equal(ns))
			})
		})
	})
	Describe("restarting a run", func() {
		Context("restarting a run", func() {
			It("restarting a run", func() {
				ns := "ns1"
				sched := "test-now-scheduler"
				rt := RunTypeRestart
				kmr := "test-kmake-run"

				By("stopping a run")
				kmsr2, err := kmc.CreateScheduleRun(ns, nil, &kmr, &sched, &rt, nil)
				Expect(err).To(BeNil())
				Expect(kmsr2.GetNamespace()).To(Equal(ns))
			})
		})
	})
})
