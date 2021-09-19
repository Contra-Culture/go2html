package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("Text()", func() {
		It("returns text node", func() {
			textNode := Text("<html>\"text\"</html>")
			Expect(textNode).NotTo(BeNil())
			str := textNode.Template().CompileWith(nil)
			Expect(str).To(Equal("&lt;html&gt;&quottext&quot&lt;/html&gt;"))
		})
	})
	Describe("RawText()", func() {
		It("returns text node", func() {
			textNode := RawText("<html>\"text\"</html>")
			Expect(textNode).NotTo(BeNil())
			str := textNode.Template().CompileWith(nil)
			Expect(str).To(Equal("<html>\"text\"</html>"))
		})
	})
	Describe("Comment()", func() {
		It("returns comment node", func() {
			commentNode := Comment("comment")
			Expect(commentNode).NotTo(BeNil())
			str := commentNode.Template().CompileWith(nil)
			Expect(str).To(Equal("<!-- comment  -->"))
		})
	})
	Describe("Injection()", func() {
		It("returns injection node", func() {
			injectionNode := Injection("comment")
			Expect(injectionNode).NotTo(BeNil())
			Expect(func() {
				injectionNode.Template().CompileWith(nil)
			}).To(PanicWith("replacement for \"comment\" key is not provied"))
			str := injectionNode.Template().CompileWith(map[string]interface{}{
				"comment": "some comment",
			})
			Expect(str).To(Equal("some comment"))
		})
	})
	Describe("Elem()", func() {
		Context("when normal elem", func() {
			It("returns element node", func() {
				elemNode := Elem("div", [][2]string{
					[2]string{"id", "myID"},
					[2]string{"class", "content"},
				},
					Elem("p", [][2]string{}, Text("<span>text</span>")),
					Injection("text"),
				)
				Expect(elemNode).NotTo(BeNil())
				str := elemNode.Template().CompileWith(map[string]interface{}{
					"text": "replacement text",
				})
				Expect(str).To(Equal("<div id=\"myID\" class=\"content\"><p>\n&lt;span&gt;text&lt;/span&gt;\n</p>replacement text</div>"))
			})
		})
		Context("when void elem", func() {
			Context("when has no children", func() {
				It("returns element node", func() {
					elemNode := Elem("br", [][2]string{
						[2]string{"id", "myID"},
						[2]string{"class", "content"},
					})
					Expect(elemNode).NotTo(BeNil())
					str := elemNode.Template().CompileWith(nil)
					Expect(str).To(Equal("<br id=\"myID\" class=\"content\"/>"))
				})
			})
			Context("when has children", func() {
				It("returns element node without children", func() {
					elemNode := Elem("br", [][2]string{
						[2]string{"id", "myID"},
						[2]string{"class", "content"},
					},
						Elem("p", [][2]string{}, Text("<span>text</span>")),
						Injection("text"),
					)
					Expect(elemNode).NotTo(BeNil())
					str := elemNode.Template().CompileWith(nil)
					Expect(str).To(Equal("<br id=\"myID\" class=\"content\"/>"))
				})
			})
		})
	})
})
