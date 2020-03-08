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
	var r KmakeRunResolver

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: &controller.KubernetesController{
				Client: k,
			},
		}
		r = res.KmakeRun()
	})

	Describe("with KmakeRun method", func() {
		Context("should be able to get", func() {

			It("Kmakename", func() {
				kmakename, err := r.Kmakename(context.Background(), fo.FakeKmakeRun())
				Expect(err).To(BeNil())
				Expect(*kmakename).To(Equal("test-kmake"))
			})

			It("Operation", func() {
				operation, err := r.Operation(context.Background(), fo.FakeKmakeRun())
				Expect(err).To(BeNil())
				Expect(operation).NotTo(BeNil())
			})

			It("Schedulerun", func() {
				n := "test-kmakeschedulerun"
				sched := "test-now-scheduler"
				rt := controller.RunTypeStart

				schedulerun, err := r.Schedulerun(context.Background(), fo.FakeKmakeRun(), &sched, &n, &rt)

				Expect(err).To(BeNil())
				Expect(schedulerun[0].GetName()).To(Equal("test-kmakeschedulerun"))
			})

			// It("Image", func() {
			// 	image, err := r.Image(context.Background(), fo.FakeKmakeRun())
			// 	Expect(err).To(BeNil())
			// 	Expect(image).To(Equal(1))
			// })

			// It("Command", func() {
			// 	command, err := r.Command(context.Background(), fo.FakeKmakeRun())
			// 	Expect(err).To(BeNil())
			// 	Expect(command).To(Equal(1))
			// })

			// It("Args", func() {
			// 	args, err := r.Args(context.Background(), fo.FakeKmakeRun())
			// 	Expect(err).To(BeNil())
			// 	Expect(len(args)).To(Equal(1))
			// })

			//+ Methods Here
		})
	})
})
