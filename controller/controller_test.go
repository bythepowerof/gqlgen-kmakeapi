package controller_test

import (
	. "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    
    context "context"

    . "github.com/bythepowerof/gqlgen-kmakeapi/controller"
    k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
    "github.com/bythepowerof/gqlgen-kmakeapi/k8s"

)

var _ = Describe("Controller", func() {
    var k k8sclient.Client
    var kmc *KubernetesController

	BeforeEach(func() {
		var err error
		k, err = k8s.FakeK8sClient()
        Expect(err).To(BeNil())
        
        kmc = &KubernetesController{ Client: k}
	})

    Describe("KubernetesController", func() {
        Context("implements interface", func() {
            It("KmakeController", func() {
                v := KubernetesController{}
                var i interface{} = v
                _, ok := i.(KmakeController)
                Expect(ok).To(Equal(false))
            
                var p interface{} = &v
                _, ok = p.(KmakeController)
                Expect(ok).To(Equal(true))
            })
        })
        Context("fetches", func() {
            It("kmake", func() {
                ns := "ns1"
                kmakes, err := kmc.Kmakes(context.Background(), &ns, nil)
                Expect(err).To(BeNil())
                Expect(kmakes[0].GetName()).To(Equal("test-kmake"))
            })
        })
    })
})
