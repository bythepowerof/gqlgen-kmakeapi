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
			By("Namespace")
			var resp1 struct {
				Namespaces []struct{ Name string }
			}
			c.MustPost(`{ namespaces(name: "ns1") { name } }`, &resp1)

			Expect(resp1.Namespaces[0].Name).To(Equal("ns1"))

			By("Kmake")
			var resp struct {
				Kmakes []struct{ Name string }
			}
			c.MustPost(`{ kmakes(namespace: "ns1") { name } }`, &resp)

			Expect(resp.Kmakes[0].Name).To(Equal("test-kmake"))

			By("Kmakerun")
			var resp3 struct {
				Kmakeruns []struct{ Name string }
			}
			c.MustPost(`{ kmakeruns(namespace: "ns1") { name } }`, &resp3)

			Expect(resp3.Kmakeruns[0].Name).To(Equal("test-kmake-run"))

			By("scheduler")
			var resp2 struct {
				KmakeSchedulers []struct{ Name string }
			}
			c.MustPost(`{ kmakeschedulers(namespace: "ns1") { name } }`, &resp2)

			Expect(resp2.KmakeSchedulers[0].Name).To(Equal("test-now-scheduler"))
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
