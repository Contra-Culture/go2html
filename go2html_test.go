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
			t := Tmplt("test", textNode)
			rn := t.Precompile()
			Expect(rn.Title).To(Equal("TEMPLATE(test) ROOT"))
			Expect(rn.Messages).To(HaveLen(0))
			Expect(rn.Children).To(HaveLen(1))
			chrn := rn.Children[0]
			Expect(chrn.Title).To(Equal("\"text\""))
			Expect(chrn.Messages).To(HaveLen(1))
			Expect(chrn.Messages[0]).To(Equal("ok"))
			Expect(chrn.Children).To(HaveLen(0))
			str := t.Populate(nil)
			Expect(str).To(Equal("&lt;html&gt;&quottext&quot&lt;/html&gt;"))
		})
	})
	Describe("RawText()", func() {
		It("returns text node", func() {
			textNode := RawText("<html>\"text\"</html>")
			Expect(textNode).NotTo(BeNil())
			t := Tmplt("test", textNode)
			rn := t.Precompile()
			Expect(rn.Title).To(Equal("TEMPLATE(test) ROOT"))
			Expect(rn.Messages).To(HaveLen(0))
			Expect(rn.Children).To(HaveLen(1))
			chrn := rn.Children[0]
			Expect(chrn.Title).To(Equal("\"text\""))
			Expect(chrn.Messages).To(HaveLen(1))
			Expect(chrn.Messages[0]).To(Equal("ok"))
			Expect(chrn.Children).To(HaveLen(0))
			str := t.Populate(nil)
			Expect(str).To(Equal("<html>\"text\"</html>"))
		})
	})
	Describe("Comment()", func() {
		It("returns comment node", func() {
			commentNode := Comment("comment")
			Expect(commentNode).NotTo(BeNil())
			t := Tmplt("test", commentNode)
			rn := t.Precompile()
			Expect(rn.Title).To(Equal("TEMPLATE(test) ROOT"))
			Expect(rn.Messages).To(HaveLen(0))
			Expect(rn.Children).To(HaveLen(1))
			chrn := rn.Children[0]
			Expect(chrn.Title).To(Equal("<!---->"))
			Expect(chrn.Messages).To(HaveLen(1))
			Expect(chrn.Messages[0]).To(Equal("ok"))
			Expect(chrn.Children).To(HaveLen(0))
			str := t.Populate(nil)
			Expect(str).To(Equal("<!-- comment -->"))
		})
	})
	Describe("Injection()", func() {
		It("returns injection node", func() {
			injectionNode := Injection("comment")
			Expect(injectionNode).NotTo(BeNil())
			Expect(func() {
				t := Tmplt("test", injectionNode)
				t.Precompile()
				t.Populate(nil)
			}).To(PanicWith("replacement for \"comment\" key is not provied"))
			t := Tmplt("test", injectionNode)
			rn := t.Precompile()
			Expect(rn.Title).To(Equal("TEMPLATE(test) ROOT"))
			Expect(rn.Messages).To(HaveLen(0))
			Expect(rn.Children).To(HaveLen(1))
			chrn := rn.Children[0]
			Expect(chrn.Title).To(Equal("{{comment}}"))
			Expect(chrn.Messages).To(HaveLen(1))
			Expect(chrn.Messages[0]).To(Equal("ok"))
			Expect(chrn.Children).To(HaveLen(0))
			str := t.Populate(map[string]interface{}{
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
				t := Tmplt("test", elemNode)
				rn := t.Precompile()
				Expect(rn.Title).To(Equal("TEMPLATE(test) ROOT"))
				Expect(rn.Messages).To(HaveLen(0))
				Expect(rn.Children).To(HaveLen(1))
				chrn1 := rn.Children[0]
				Expect(chrn1.Title).To(Equal("<div>"))
				Expect(chrn1.Messages).To(HaveLen(2))
				Expect(chrn1.Messages[0]).To(Equal("ok: opening"))
				Expect(chrn1.Messages[1]).To(Equal("ok: closing"))
				Expect(chrn1.Children).To(HaveLen(2))
				chrn1_1 := chrn1.Children[0]
				Expect(chrn1_1.Title).To(Equal("<p>"))
				Expect(chrn1_1.Messages).To(HaveLen(2))
				Expect(chrn1_1.Messages[0]).To(Equal("ok: opening"))
				Expect(chrn1_1.Messages[1]).To(Equal("ok: closing"))
				Expect(chrn1_1.Children).To(HaveLen(1))
				chrn1_1_1 := chrn1_1.Children[0]
				Expect(chrn1_1_1.Title).To(Equal("\"text\""))
				Expect(chrn1_1_1.Messages).To(HaveLen(1))
				Expect(chrn1_1_1.Messages[0]).To(Equal("ok"))
				Expect(chrn1_1_1.Children).To(BeEmpty())
				chrn1_2 := chrn1.Children[1]
				Expect(chrn1_2.Title).To(Equal("{{text}}"))
				Expect(chrn1_2.Messages).To(HaveLen(1))
				Expect(chrn1_2.Messages[0]).To(Equal("ok"))
				Expect(chrn1_2.Children).To(HaveLen(0))
				str := t.Populate(map[string]interface{}{
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
					t := Tmplt("test", elemNode)
					rn := t.Precompile()
					Expect(rn.Title).To(Equal("TEMPLATE(test) ROOT"))
					Expect(rn.Messages).To(HaveLen(0))
					Expect(rn.Children).To(HaveLen(1))
					chrn1 := rn.Children[0]
					Expect(chrn1.Title).To(Equal("<br>"))
					Expect(chrn1.Messages).To(HaveLen(1))
					Expect(chrn1.Messages[0]).To(Equal("ok: self-closing"))
					Expect(chrn1.Children).To(HaveLen(0))
					str := t.Populate(nil)
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
					t := Tmplt("test", elemNode)
					rn := t.Precompile()
					Expect(rn.Title).To(Equal("TEMPLATE(test) ROOT"))
					Expect(rn.Messages).To(HaveLen(0))
					Expect(rn.Children).To(HaveLen(1))
					chrn1 := rn.Children[0]
					Expect(chrn1.Title).To(Equal("<br>"))
					Expect(chrn1.Messages).To(HaveLen(1))
					Expect(chrn1.Messages[0]).To(Equal("error: void element can't have children (children ignored)"))
					Expect(chrn1.Children).To(HaveLen(0))
					str := t.Populate(nil)
					Expect(str).To(Equal("<br id=\"myID\" class=\"content\"/>"))
				})
			})
		})
	})
})
