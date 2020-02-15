package gqlgen_kmakeapi

import (
	// "encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/99designs/gqlgen/client"
	"github.com/bythepowerof/gqlgen-kmakeapi/k8s"
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
			By("Kmake")
			var resp struct {
				Kmakes []struct{ Name string }
			}
			c.MustPost(`{ kmakes(namespace: "ns1") { name } }`, &resp)

			Expect(resp.Kmakes[0].Name).To(Equal("test-kmake"))
		})
	})

	// Context("with given scheme", func() {
	// 	BeforeEach(func(done Done) {
	// 		scheme := runtime.NewScheme()
	// 		Expect(corev1.AddToScheme(scheme)).To(Succeed())
	// 		Expect(appsv1.AddToScheme(scheme)).To(Succeed())
	// 		cl = fake.NewFakeClientWithScheme(scheme, dep, dep2, cm)
	// 		close(done)
	// 	})
	// 	AssertClientBehavior()
	// })
})
