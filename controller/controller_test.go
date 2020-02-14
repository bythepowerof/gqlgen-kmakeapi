package controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/bythepowerof/gqlgen-kmakeapi/controller"
)

var _ = Describe("Controller", func() {
    Describe("KubernetesController", func() {
        Context("implements interface", func() {
            It("KmakeController", func() {

                v := KubernetesController{}
                var i interface{} = v
                _, ok := i.(KmakeController)
                Expect(ok).To(Equal(false))

                // c.Assert(ok, Equals, false)
            
                var p interface{} = &v
                _, ok = p.(KmakeController)
                // c.Assert(ok, Equals, true)

                Expect(ok).To(Equal(true))
            })
        })
    })
})
