package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("CommentNode", func() {
		n := Comment("comment")

		Describe("Comment()", func() {
			It("returns comment node", func() {
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(n.Title()).To(Equal("<!---->"))
			})
		})
		Describe(".WriteTo()", func() {
			It("writes", func() {
				t := Tmplt("testTemplate", n)
				nr := t.Precompile()
				Expect(nr.Title).To(Equal("TEMPLATE(testTemplate)"))
				Expect(nr.Messages).To(BeEmpty())
				Expect(nr.Children).To(HaveLen(1))
				ch := nr.Children[0]
				Expect(ch.Title).To(Equal("<!---->"))
				Expect(ch.Messages).To(Equal([]string{
					"ok",
				}))
				Expect(ch.Children).To(BeEmpty())
				s := t.Populate(map[string]interface{}{})
				Expect(s).To(Equal("<!-- comment -->"))
			})
		})
	})
})
