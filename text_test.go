package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("TextNode", func() {
		Describe("Text()", func() {
			It("returns text node", func() {
				n := Text("<html>\"text\"</html>")
				Expect(n).NotTo(BeNil())
			})
		})
		Describe("RawText()", func() {
			It("returns text node", func() {
				n := RawText("<html>\"text\"</html>")
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			Context("when safe text node", func() {
				It("returns title", func() {
					n := Text("<html>\"text\"</html>")
					Expect(n.Title()).To(Equal("\"text\""))
				})
			})
			Context("when raw text node", func() {
				It("returns title", func() {
					n := RawText("<html>\"text\"</html>")
					Expect(n.Title()).To(Equal("!\"text\""))
				})
			})
		})
		Describe(".WriteTo()", func() {
			Context("when safe text node", func() {
				It("writes", func() {
					t := Tmplt("testTemplate", Text("<html>\"text\"</html>"))
					nr := t.Precompile()
					Expect(nr.Title).To(Equal("TEMPLATE(testTemplate)"))
					Expect(nr.Messages).To(BeEmpty())
					Expect(nr.Children).To(HaveLen(1))
					ch := nr.Children[0]
					Expect(ch.Title).To(Equal("\"text\""))
					Expect(ch.Messages).To(Equal([]string{
						"ok",
					}))
					Expect(ch.Children).To(BeEmpty())
					s := t.Populate(nil)
					Expect(s).To(Equal("&lt;html&gt;&quottext&quot&lt;/html&gt;"))
				})
			})
			Context("when raw text node", func() {
				It("writes", func() {
					t := Tmplt("testTemplate", RawText("<html>\"text\"</html>"))
					nr := t.Precompile()
					Expect(nr.Title).To(Equal("TEMPLATE(testTemplate)"))
					Expect(nr.Messages).To(BeEmpty())
					Expect(nr.Children).To(HaveLen(1))
					ch := nr.Children[0]
					Expect(ch.Title).To(Equal("!\"text\""))
					Expect(ch.Messages).To(Equal([]string{
						"ok",
					}))
					Expect(ch.Children).To(BeEmpty())
					s := t.Populate(nil)
					Expect(s).To(Equal("<html>\"text\"</html>"))
				})
			})
		})
	})
})
