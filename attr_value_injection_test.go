package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("AttrValueInjectionNode", func() {
		var n = AttrValueInjection("myAttr", "myAttr")

		Describe("AttrValueInjection()", func() {
			It("returns attr value injection node", func() {
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(n.Title()).To(Equal("attr={{myAttr}}"))
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
				Expect(ch.Title).To(Equal("attr={{myAttr}}"))
				Expect(ch.Messages).To(Equal([]string{
					"ok",
				}))
				Expect(ch.Children).To(BeEmpty())
				Expect(func() {
					t.Populate(map[string]interface{}{})
				}).Should(PanicWith("replacement for \"myAttr\" key is not provied"))
				s := t.Populate(map[string]interface{}{
					"myAttr": "test value",
				})
				Expect(s).To(Equal(" myAttr=\"test value\""))
			})
		})
	})
})
