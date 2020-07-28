package gqlgen_kmakeapi

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/gqlgen-kmakeapi/k8s"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Fake client", func() {
	var k k8sclient.Client
	var fo *k8s.FakeObjects
	var r KmakeScheduleRunResolver

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: controller.NewKubernetesController(k, "all"),
		}
		r = res.KmakeScheduleRun()
	})

	Describe("with KmakeScheduleRun method", func() {
		Context("should be able to get", func() {

			It("Kmakename", func() {
				kmakename, err := r.Kmakename(context.Background(), fo.FakeKmakeScheduleRun())
				Expect(err).To(BeNil())
				Expect(*kmakename).To(Equal("test-kmake"))
			})

			It("Kmakerunname", func() {
				kmakerunname, err := r.Kmakerunname(context.Background(), fo.FakeKmakeScheduleRun())
				Expect(err).To(BeNil())
				Expect(*kmakerunname).To(Equal("test-kmake-run"))
			})

			It("Kmakeschedulename", func() {
				kmakeschedulename, err := r.Kmakeschedulename(context.Background(), fo.FakeKmakeScheduleRun())
				Expect(err).To(BeNil())
				Expect(*kmakeschedulename).To(Equal("test-now-scheduler"))
			})

			It("Operation", func() {
				operation, err := r.Operation(context.Background(), fo.FakeKmakeScheduleRun())
				Expect(err).To(BeNil())
				Expect(operation).NotTo(BeNil())
			})

			It("Uid", func() {
				uid, err := r.UID(context.Background(), fo.FakeKmakeScheduleRun())
				Expect(err).To(BeNil())
				Expect(*uid).To(Equal(""))
			})
		})
	})
})
