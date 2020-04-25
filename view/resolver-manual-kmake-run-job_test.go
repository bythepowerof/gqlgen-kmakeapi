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
	var r KmakeRunJobResolver

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: controller.NewKubernetesController(k, "all"),
		}
		r = res.KmakeRunJob()
	})

	Describe("with KmakeRun method", func() {
		Context("should be able to get", func() {

			It("Image", func() {
				image, err := r.Image(context.Background(), fo.FakeKmakeRun().Spec.KmakeRunOperation.Job)
				Expect(err).To(BeNil())
				Expect(image).To(Equal("image:latest"))
			})

			It("Command", func() {
				command, err := r.Command(context.Background(), fo.FakeKmakeRun().Spec.KmakeRunOperation.Job)
				Expect(err).To(BeNil())
				Expect(len(command)).To(Equal(1))
			})

			It("Args", func() {
				args, err := r.Args(context.Background(), fo.FakeKmakeRun().Spec.KmakeRunOperation.Job)
				Expect(err).To(BeNil())
				Expect(len(args)).To(Equal(2))
			})

			//+ Methods Here
		})
	})
})
