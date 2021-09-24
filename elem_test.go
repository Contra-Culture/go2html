package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("ElementNode", func() {
		Describe("Elem()", func() {
			It("returns element node", func() {
				n := Elem("p", []Node{}, []Node{})
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				Expect(Elem("p", []Node{}, []Node{}).Title()).To(Equal("<p>"))
				Expect(Elem("div", []Node{}, []Node{}).Title()).To(Equal("<div>"))
				Expect(Elem("head", []Node{}, []Node{}).Title()).To(Equal("<head>"))
				Expect(Elem("meta", []Node{}, []Node{}).Title()).To(Equal("<meta>"))

			})
		})
		Describe(".WriteTo()", func() {
			Context("when normal element", func() {
				Context("when no attributes and children", func() {
					It("writes", func() {
						t := Tmplt("testTemplate", Elem("div", []Node{}, []Node{}))
						rn := t.Precompile()
						Expect(rn.Title).To(Equal("TEMPLATE(testTemplate)"))
						Expect(rn.Messages).To(BeEmpty())
						Expect(rn.Children).To(HaveLen(1))
						ch := rn.Children[0]
						Expect(ch.Title).To(Equal("<div>"))
						Expect(ch.Messages).To(Equal([]string{
							"ok: opening",
							"ok: closing",
						}))
						Expect(ch.Children).To(BeEmpty())
						s := t.Populate(map[string]interface{}{})
						Expect(s).To(Equal("<div></div>"))
					})
				})
				Context("when with attributes", func() {
					It("writes", func() {
						t := Tmplt("testTemplate", Elem("div",
							[]Node{
								Attr("test1", "value1"),
								AttrInjection("injattr"),
								AttrValueInjection("test2", "injval"),
							},
							[]Node{},
						))
						rn := t.Precompile()
						Expect(rn.Title).To(Equal("TEMPLATE(testTemplate)"))
						Expect(rn.Messages).To(BeEmpty())
						Expect(rn.Children).To(HaveLen(1))
						ch := rn.Children[0]
						Expect(ch.Title).To(Equal("<div>"))
						Expect(ch.Messages).To(Equal([]string{
							"ok: opening",
							"ok: closing",
						}))
						Expect(ch.Children).To(HaveLen(1))
						ch1_1 := ch.Children[0]
						Expect(ch1_1.Title).To(Equal("[]attrs"))
						Expect(ch1_1.Messages).To(Equal([]string{
							"ok",
							"ok",
							"ok",
						}))
						Expect(ch1_1.Children).To(BeEmpty())
						Expect(func() {
							t.Populate(nil)
						}).To(
							SatisfyAny(
								PanicWith("replacement for \"injval\" key is not provied"),
								PanicWith("replacement for \"injattr\" key is not provied"),
							),
						)
						s := t.Populate(map[string]interface{}{
							"injattr": "testValue=\"testValue\"",
							"injval":  "testValue2",
						})
						Expect(s).To(Equal("<div test1=\"value1\" testValue=\"testValue\" test2=\"testValue2\"></div>"))
					})
				})
				Context("when with attributes and children", func() {
					It("writes", func() {
						t := Tmplt("testTemplate",
							Elem("div",
								[]Node{
									Attr("class", "my-class"),
								},
								[]Node{
									Elem("p",
										[]Node{},
										[]Node{
											Text("test text"),
										}),
								},
							),
						)
						rn := t.Precompile()
						Expect(rn.Title).To(Equal("TEMPLATE(testTemplate)"))
						Expect(rn.Messages).To(BeEmpty())
						Expect(rn.Children).To(HaveLen(1))
						ch1 := rn.Children[0]
						Expect(ch1.Title).To(Equal("<div>"))
						Expect(ch1.Messages).To(Equal([]string{
							"ok: opening",
							"ok: closing",
						}))
						Expect(ch1.Children).To(HaveLen(2))
						ch1_1 := ch1.Children[0]
						Expect(ch1_1.Title).To(Equal("[]attrs"))
						Expect(ch1_1.Messages).To(Equal([]string{
							"ok",
						}))
						Expect(ch1_1.Children).To(BeEmpty())
						ch1_2 := ch1.Children[1]
						Expect(ch1_2.Title).To(Equal("<p>"))
						Expect(ch1_2.Messages).To(Equal([]string{
							"ok: opening",
							"ok: closing",
						}))
						Expect(ch1_2.Children).To(HaveLen(1))
						ch1_2_1 := ch1_2.Children[0]
						Expect(ch1_2_1.Title).To(Equal("\"text\""))
						Expect(ch1_2_1.Messages).To(Equal([]string{
							"ok",
						}))
						Expect(ch1_2_1.Children).To(BeEmpty())
						s := t.Populate(nil)
						Expect(s).To(Equal("<div class=\"my-class\"><p>test text</p></div>"))
					})
				})
				Context("when void element", func() {
					Context("when has no children", func() {
						It("writes", func() {
							t := Tmplt("testTemplate",
								Elem("br", []Node{
									Attr("id", "myID"),
									Attr("class", "content"),
									AttrInjection("test-attr"),
								},
									[]Node{},
								))
							rn := t.Precompile()
							Expect(rn.Title).To(Equal("TEMPLATE(testTemplate)"))
							Expect(rn.Messages).To(BeEmpty())
							Expect(rn.Children).To(HaveLen(1))
							ch1 := rn.Children[0]
							Expect(ch1.Title).To(Equal("<br>"))
							Expect(ch1.Messages).To(Equal([]string{
								"ok: self-closing",
							}))
							Expect(ch1.Children).To(HaveLen(1))
							ch1_1 := ch1.Children[0]
							Expect(ch1_1.Title).To(Equal("[]attrs"))
							Expect(ch1_1.Messages).To(Equal([]string{
								"ok",
								"ok",
								"ok",
							}))
							Expect(ch1_1.Children).To(BeEmpty())
							Expect(func() {
								t.Populate(nil)
							}).To(PanicWith("replacement for \"test-attr\" key is not provied"))
							s := t.Populate(map[string]interface{}{
								"test-attr": "testAttr=\"test value\"",
							})
							Expect(s).To(Equal("<br id=\"myID\" class=\"content\" testAttr=\"test value\"/>"))
						})
					})
					Context("when has children", func() {
						It("fails", func() {
							t := Tmplt("testTemplate",
								Elem("br", []Node{
									Attr("id", "myID"),
									AttrValueInjection("class", "test-class"),
								}, []Node{
									Elem("p", []Node{}, []Node{
										Text("<span>text</span>"),
									}),
									Injection("text"),
								},
								),
							)
							rn := t.Precompile()
							Expect(rn.Title).To(Equal("TEMPLATE(testTemplate)"))
							Expect(rn.Messages).To(HaveLen(0))
							Expect(rn.Children).To(HaveLen(1))
							ch1 := rn.Children[0]
							Expect(ch1.Title).To(Equal("<br>"))
							Expect(ch1.Messages).To(Equal([]string{
								"error: void element can't have children (children ignored)",
							}))
							Expect(ch1.Children).To(HaveLen(1))
							ch1_1 := ch1.Children[0]
							Expect(ch1_1.Title).To(Equal("[]attrs"))
							Expect(ch1_1.Messages).To(Equal([]string{
								"ok",
								"ok",
							}))
							Expect(ch1_1.Children).To(BeEmpty())
							str := t.Populate(map[string]interface{}{
								"test-class": "testClass",
							})
							Expect(str).To(Equal("<br id=\"myID\" class=\"testClass\"/>"))
						})
					})
				})
			})
		})
	})
})
