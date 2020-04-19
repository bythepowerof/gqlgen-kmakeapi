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
	var r MutationResolver

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: controller.NewKubernetesController(k, nil, "all"),
		}
		r = res.Mutation()
	})

	Describe("with Mutation method", func() {
		Context("should be able to get", func() {

			It("Reset", func() {
				in := controller.NewReset{
					Namespace:      "ns1",
					Kmakescheduler: "test-now-scheduler",
					Full:           true,
				}

				reset, err := r.Reset(context.Background(), in)
				Expect(err).To(BeNil())
				Expect(reset).NotTo(BeNil())
			})

			It("Stop", func() {
				in := controller.RunLevelIn{
					Namespace:      "ns1",
					Kmakescheduler: "test-now-scheduler",
					Kmakerun:       "test-kmake-run",
				}

				stop, err := r.Stop(context.Background(), in)
				Expect(err).To(BeNil())
				Expect(stop).NotTo(BeNil())
			})

			It("Restart", func() {
				in := controller.RunLevelIn{
					Namespace:      "ns1",
					Kmakescheduler: "test-now-scheduler",
					Kmakerun:       "test-kmake-run",
				}

				restart, err := r.Restart(context.Background(), in)
				Expect(err).To(BeNil())
				Expect(restart).NotTo(BeNil())
			})

			//+ Methods Here
		})
	})
})
