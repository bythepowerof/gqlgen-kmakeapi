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
	var r KmakeNowSchedulerResolver

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: controller.NewKubernetesController(k, nil, "all"),
		}
		r = res.KmakeNowScheduler()
	})

	Describe("with KmakeNowScheduler method", func() {
		Context("should be able to get", func() {

			It("Scheduleruns", func() {
				scheduleruns, err := r.Scheduleruns(context.Background(), fo.FakeKmakeNowScheduler(), nil, nil, nil, nil)
				// Scheduleruns(ctx context.Context, obj *v1.KmakeNowScheduler, kmake *string, kmakerun *string, name *string, runtype *controller.RunType) ([]*v1.KmakeScheduleRun, error) {

				Expect(err).To(BeNil())
				Expect(len(scheduleruns)).To(Equal(1))
			})

			//+ Methods Here
		})
	})
})
