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
	var r QueryResolver

	BeforeEach(func() {
		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		res := &Resolver{
			KmakeController: controller.NewKubernetesController(k, "all"),
			// KmakeController: controller.NewKubernetesController(k, nil, "all"),
		}

		r = res.Query()
	})

	Describe("with Query method", func() {
		Context("should be able to get", func() {

			It("Namespaces", func() {
				ns := "ns1"
				namespaces, err := r.Namespaces(context.Background(), &ns)
				Expect(err).To(BeNil())
				Expect(len(namespaces)).To(Equal(1))
			})

			It("KmakeObjects", func() {
				ns := "ns1"

				kmakeobjects, err := r.KmakeObjects(context.Background(), ns, nil)
				Expect(err).To(BeNil())
				Expect(len(kmakeobjects)).To(Equal(4))
			})

			It("Kmakeschedulers", func() {
				ns := "ns1"

				kmakeschedulers, err := r.Kmakeschedulers(context.Background(), ns, nil, nil)
				Expect(err).To(BeNil())
				Expect(len(kmakeschedulers)).To(Equal(1))
			})

			It("Kmakes", func() {
				ns := "ns1"

				kmakes, err := r.Kmakes(context.Background(), ns, nil)
				Expect(err).To(BeNil())
				Expect(len(kmakes)).To(Equal(1))
			})

			It("Kmakeruns", func() {
				ns := "ns1"

				kmakeruns, err := r.Kmakeruns(context.Background(), ns, nil, nil, nil)
				Expect(err).To(BeNil())
				Expect(len(kmakeruns)).To(Equal(1))
			})

			It("Kmakescheduleruns", func() {
				ns := "ns1"

				kmakescheduleruns, err := r.Kmakescheduleruns(context.Background(), ns, nil, nil, nil, nil, nil)
				Expect(err).To(BeNil())
				Expect(len(kmakescheduleruns)).To(Equal(1))
			})

			//+ Methods Here
		})
	})
})
