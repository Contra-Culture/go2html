package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("AttrInjectionNode", func() {
		var n = AttrInjection("testAttr")

		Describe("AttrInjection()", func() {
			It("returns attr injection node", func() {
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(n.Title()).To(Equal("{{testAttr}}"))
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
				Expect(ch.Title).To(Equal("{{testAttr}}"))
				Expect(ch.Messages).To(Equal([]string{
					"ok",
				}))
				Expect(ch.Children).To(BeEmpty())
				Expect(func() {
					t.Populate(map[string]interface{}{})
				}).Should(PanicWith("replacement for \"testAttr\" key is not provied"))
				s := t.Populate(map[string]interface{}{"testAttr": "myAttr=\"value\""})
				Expect(s).To(Equal(" myAttr=\"value\""))
			})
		})
	})
})
