package gqlgen_kmakeapi

import (
	// "encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/99designs/gqlgen/client"
	"github.com/bythepowerof/gqlgen-kmakeapi/k8s"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Fake client", func() {
	var k k8sclient.Client
	var c *client.Client
	var fo *k8s.FakeObjects

	BeforeEach(func() {

		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		c = FakeHTTPServer(k)
	})

	Context("with default scheme.Scheme", func() {
		It("should be able to get", func() {
			By("kmake objects")
			var resp struct {
				KmakeObjects []struct{ Name string }
			}
			c.MustPost(`{ kmakeObjects(namespace: "ns1") { name } }`, &resp)

			Expect(len(resp.KmakeObjects)).To(Equal(4))
		})

	})
})