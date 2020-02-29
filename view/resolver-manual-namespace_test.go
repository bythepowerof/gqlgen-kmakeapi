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

	BeforeEach(func() {

		var err error
		k, err = k8s.FakeK8sClient()
		Expect(err).To(BeNil())

		c = FakeHTTPServer(k)
	})

	Context("with default scheme.Scheme", func() {
		It("should be able to get", func() {
			By("namespace")
			var resp struct {
				Namespaces []struct{ Name string }
			}
			c.MustPost(`{ namespaces(name: "ns1") { name } }`, &resp)

			 Expect(resp.Namespaces[0].Name).To(Equal("ns1"))
		})

	})
})
