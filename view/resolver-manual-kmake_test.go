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
	var r KmakeResolver

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: controller.NewKubernetesController(k, nil, "all"),
		}
		r = res.Kmake()
	})

	Describe("with kmake method", func() {
		Context("should be able to get", func() {

			It("Variables", func() {
				vars, err := r.Variables(context.Background(), fo.FakeKmake())
				Expect(err).To(BeNil())
				Expect(len(vars)).To(Equal(2))
			})

			It("Rules", func() {
				rules, err := r.Rules(context.Background(), fo.FakeKmake())
				Expect(err).To(BeNil())
				Expect(len(rules)).To(Equal(2))
			})

			It("Runs", func() {
				runs, err := r.Runs(context.Background(), fo.FakeKmake(), nil, nil)
				Expect(err).To(BeNil())
				Expect(len(runs)).To(Equal(1))
			})
		})
	})
})
