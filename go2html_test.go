package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("WrongNode", func() {
		var n Node
		BeforeEach(func() {
			n = Wrong("<p>", []string{"some error"})
		})
		Describe("Wrong()", func() {
			It("returns wrong node", func() {
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(n.Title()).To(Equal("WRONG(<p>)"))
			})
		})
		Describe(".Commit()", func() {
			It("writes", func() {
				spec := Spec("testTemplate", n)
				t, nr := spec.Precompile()
				Expect(nr.Title).To(Equal("TEMPLATE(testTemplate)"))
				Expect(nr.Messages).To(BeEmpty())
				Expect(nr.Children).To(HaveLen(1))
				ch := nr.Children[0]
				Expect(ch.Title).To(Equal("WRONG(<p>)"))
				Expect(ch.Messages).To(Equal([]string{
					"error: some error",
				}))
				Expect(ch.Children).To(BeEmpty())
				s := t.Populate(map[string]interface{}{})
				Expect(s).To(Equal("<!-- WRONG(<p>) -->"))
			})
		})
	})
})
