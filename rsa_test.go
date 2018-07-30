package rsa

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSecurity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Security Suite")
}

var _ = Describe("security", func() {
	It("rsa/MakeSSHKeyPair", func() {
		pvt, pub, err := MakeSSHKeyPair("123456778", 2048)
		Expect(err).Should(BeNil())
		Expect(pvt).ShouldNot(BeEmpty())
		Expect(pub).ShouldNot(BeEmpty())
		fmt.Println(pub)

		pvt, pub, err = MakeSSHKeyPair("123456778", 2048)
		Expect(err).Should(BeNil())
		Expect(pvt).ShouldNot(BeEmpty())
		Expect(pub).ShouldNot(BeEmpty())
		fmt.Println(pub)

		pvt, pub, err = MakeSSHKeyPair("TPUKQLBWX3WJCPLXM4TLL2SZJ2MCVZXD", 2048)
		Expect(err).Should(BeNil())
		Expect(pvt).ShouldNot(BeEmpty())
		Expect(pub).ShouldNot(BeEmpty())
		fmt.Println(pub)

		pvt, pub, err = MakeSSHKeyPair("", 2048)
		Expect(err).Should(BeNil())
		Expect(pvt).ShouldNot(BeEmpty())
		Expect(pub).ShouldNot(BeEmpty())
		fmt.Println(pub)

		pvt, pub, err = MakeSSHKeyPair("", 2048)
		Expect(err).Should(BeNil())
		Expect(pvt).ShouldNot(BeEmpty())
		Expect(pub).ShouldNot(BeEmpty())
		fmt.Println(pub)
	})
})
