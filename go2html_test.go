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
			t := textNode.Template()
			r := t.Report()
			Expect(r.Title).To(Equal("\"text\""))
			Expect(r.Messages).To(HaveLen(1))
			Expect(r.Messages[0]).To(Equal("ok"))
			Expect(r.Children).To(HaveLen(0))
			str := t.CompileWith(nil)
			Expect(str).To(Equal("&lt;html&gt;&quottext&quot&lt;/html&gt;"))
		})
	})
	Describe("RawText()", func() {
		It("returns text node", func() {
			textNode := RawText("<html>\"text\"</html>")
			Expect(textNode).NotTo(BeNil())
			t := textNode.Template()
			r := t.Report()
			Expect(r.Title).To(Equal("\"text\""))
			Expect(r.Messages).To(HaveLen(1))
			Expect(r.Messages[0]).To(Equal("ok"))
			Expect(r.Children).To(HaveLen(0))
			str := t.CompileWith(nil)
			Expect(str).To(Equal("<html>\"text\"</html>"))
		})
	})
	Describe("Comment()", func() {
		It("returns comment node", func() {
			commentNode := Comment("comment")
			Expect(commentNode).NotTo(BeNil())
			t := commentNode.Template()
			r := t.Report()
			Expect(r.Title).To(Equal("<!---->"))
			Expect(r.Messages).To(HaveLen(1))
			Expect(r.Messages[0]).To(Equal("ok"))
			Expect(r.Children).To(HaveLen(0))
			str := t.CompileWith(nil)
			Expect(str).To(Equal("<!-- comment -->"))
		})
	})
	Describe("Injection()", func() {
		It("returns injection node", func() {
			injectionNode := Injection("comment")
			Expect(injectionNode).NotTo(BeNil())
			Expect(func() {
				injectionNode.Template().CompileWith(nil)
			}).To(PanicWith("replacement for \"comment\" key is not provied"))
			t := injectionNode.Template()
			r := t.Report()
			Expect(r.Title).To(Equal("{{comment}}"))
			Expect(r.Messages).To(HaveLen(1))
			Expect(r.Messages[0]).To(Equal("ok"))
			Expect(r.Children).To(HaveLen(0))
			str := t.CompileWith(map[string]interface{}{
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
				t := elemNode.Template()
				r := t.Report()
				Expect(r.Title).To(Equal("<div>"))
				Expect(r.Messages).To(HaveLen(2))
				Expect(r.Messages[0]).To(Equal("ok: opening"))
				Expect(r.Messages[1]).To(Equal("ok: closing"))
				Expect(r.Children).To(HaveLen(2))
				chr1 := r.Children[0]
				Expect(chr1.Title).To(Equal("<p>"))
				Expect(chr1.Messages).To(HaveLen(2))
				Expect(chr1.Messages[0]).To(Equal("ok: opening"))
				Expect(chr1.Messages[1]).To(Equal("ok: closing"))
				Expect(chr1.Children).To(HaveLen(1))
				chr1_1 := chr1.Children[0]
				Expect(chr1_1.Title).To(Equal("\"text\""))
				Expect(chr1_1.Messages).To(HaveLen(1))
				Expect(chr1_1.Messages[0]).To(Equal("ok"))
				Expect(chr1_1.Children).To(BeEmpty())
				chr2 := r.Children[1]
				Expect(chr2.Title).To(Equal("{{text}}"))
				Expect(chr2.Messages).To(HaveLen(1))
				Expect(chr2.Messages[0]).To(Equal("ok"))
				Expect(chr2.Children).To(HaveLen(0))
				str := elemNode.Template().CompileWith(map[string]interface{}{
					"text": "replacement text",
				})
				Expect(str).To(Equal("<div id=\"myID\" class=\"content\"><p>&lt;span&gt;text&lt;/span&gt;</p>replacement text</div>"))
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
					t := elemNode.Template()
					r := t.Report()
					Expect(r.Title).To(Equal("<br>"))
					Expect(r.Messages).To(HaveLen(1))
					Expect(r.Messages[0]).To(Equal("ok: self-closing"))
					Expect(r.Children).To(HaveLen(0))
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
					t := elemNode.Template()
					r := t.Report()
					Expect(r.Title).To(Equal("<br>"))
					Expect(r.Messages).To(HaveLen(1))
					Expect(r.Messages[0]).To(Equal("error: void element can't have children (children ignored)"))
					Expect(r.Children).To(HaveLen(0))

					str := t.CompileWith(nil)
					Expect(str).To(Equal("<br id=\"myID\" class=\"content\"/>"))
				})
			})
		})
	})
})
