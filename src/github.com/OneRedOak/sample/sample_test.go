package sample_test

import (
	. "github.com/OneRedOak/sample"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
    Describe("Calculating square root", func() {
        Context("of the number 4", func() {
            It("should result 2", func() {
                Setup()
                Expect(GetResult()).To(Equal(2.0))
            })
        })
    })
})
