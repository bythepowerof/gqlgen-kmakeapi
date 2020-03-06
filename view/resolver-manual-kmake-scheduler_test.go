
package gqlgen_kmakeapi

import (
	// "context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bythepowerof/gqlgen-kmakeapi/k8s"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
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
			KmakeController: &controller.KubernetesController{
				Client: k,
			},
		}
		r = res.KmakeNowScheduler()
	})

	Describe("with KmakeNowScheduler method", func() {
		Context("should be able to get", func() {
        //+ Methods Here

		})
	})
})
