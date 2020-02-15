/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gqlgen_kmakeapi

import (
	// "encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"

	myfake "github.com/bythepowerof/gqlgen-kmakeapi/fakek8s"
	// apimain "github.com/bythepowerof/gqlgen-kmakeapi/main"

)

var _ = Describe("Fake client", func() {
	var k k8sclient.Client
	var c *client.Client

	BeforeEach(func() {

		var err error
		k, err = myfake.FakeK8sClient()
		Expect(err).To(BeNil())

		// c = apimain.FakeHTTPServer(c)
		c = client.New(handler.NewDefaultServer(NewExecutableSchema(
			Config{
				Resolvers: &Resolver{
					KmakeController: &controller.KubernetesController{
						Client: k,
					},
				},
			},
		)))
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
