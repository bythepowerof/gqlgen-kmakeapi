package gqlgen_kmakeapi

import (
	// "encoding/json"
	"context"

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
	var r *kmakeResolver

	BeforeEach(func() {

		var err error
		fo = &k8s.FakeObjects{}

		k, err = fo.FakeK8sClient()
		Expect(err).To(BeNil())

		c = FakeHTTPServer(k)
		r = &kmakeResolver{}
	})

	Context("with default scheme.Scheme", func() {
		It("should be able to get", func() {
			By("Kmake")
			var resp struct {
				Kmakes []struct{ Name string }
			}
			c.MustPost(`{ kmakes(namespace: "ns1") { name } }`, &resp)
			Expect(resp.Kmakes[0].Name).To(Equal("test-kmake"))

			By("Variables")
			vars, err := r.Variables(context.Background(), fo.FakeKmake())
			Expect(err).To(BeNil())
			Expect(len(vars)).To(Equal(2))
		})
	})
})

