package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("Template", func() {
		Describe("TemplateNode", func() {
			Describe(".Node()", func() {
				It("returns template node", func() {
					t := Tmplt("testTemplate", Elem("div", []Node{}, []Node{}))
					n := t.Node(NO_ALIAS, NO_INJECTION_SCOPE)
					Expect(n).NotTo(BeNil())
				})
			})
			Describe(".Title()", func() {
				Context("when with alias", func() {
					It("returns title", func() {
						t := Tmplt("testTemplate", Elem("div", []Node{}, []Node{}))
						n := t.Node("some/alias/testTemplate", NO_INJECTION_SCOPE)
						Expect(n.Title()).To(Equal("TEMPLATE(some/alias/testTemplate)"))
					})
				})
				Context("when without alias", func() {
					It("returns title", func() {
						t := Tmplt("testTemplate", Elem("div", []Node{}, []Node{}))
						n := t.Node(NO_ALIAS, NO_INJECTION_SCOPE)
						Expect(n.Title()).To(Equal("TEMPLATE(testTemplate)"))
					})
				})
			})
			Describe(".WriteTo()", func() {
				Context("when with alias", func() {
					Context("when with injection scope", func() {
						It("writes", func() {
							t := Tmplt("testTemplate",
								Elem("div",
									[]Node{
										AttrValueInjection("class", "divClass"),
									},
									[]Node{
										Elem("p",
											[]Node{},
											[]Node{
												Injection("paragraph"),
											},
										),
									},
								),
							)
							tmain := Tmplt("mainTemplate",
								Elem("div",
									[]Node{
										Attr("id", "wrapper"),
									},
									[]Node{
										t.Node("test/alias/template", "testScope"),
									},
								),
							)
							nr := tmain.Precompile()
							Expect(nr.Title).To(Equal("TEMPLATE(mainTemplate)"))
							Expect(nr.Messages).To(BeEmpty())
							Expect(nr.Children).To(HaveLen(1))
							ch := nr.Children[0]
							Expect(ch.Title).To(Equal("<div>"))
							Expect(ch.Messages).To(Equal([]string{
								"ok: opening",
								"ok: closing",
							}))
							Expect(ch.Children).To(HaveLen(2))
							ch1 := ch.Children[0]
							Expect(ch1.Title).To(Equal("[]attrs"))
							Expect(ch1.Messages).To(Equal([]string{
								"ok",
							}))
							Expect(ch1.Children).To(BeEmpty())
							ch2 := ch.Children[1]
							Expect(ch2.Title).To(Equal("TEMPLATE(test/alias/template)"))
							Expect(ch2.Messages).To(Equal([]string{
								"ok",
								"ok: injection (testScope.divClass)",
								"ok",
								"ok: injection (testScope.paragraph)",
								"ok",
							}))
							Expect(ch2.Children).To(BeEmpty())
							Expect(func() {
								tmain.Populate(nil)
							}).To(
								SatisfyAny(
									PanicWith("replacement for \"testScope.divClass\" key is not provied"),
									PanicWith("replacement for \"testScope.paragraph\" key is not provied"),
								),
							)
							Expect(func() {
								tmain.Populate(map[string]interface{}{
									"testScope.divClass": "main-content",
								})
							}).To(
								PanicWith("replacement for \"testScope.paragraph\" key is not provied"),
							)
							s := tmain.Populate(map[string]interface{}{
								"testScope.divClass":  "main-content",
								"testScope.paragraph": "Test text.",
							})
							Expect(s).To(
								Equal("<div id=\"wrapper\"><div class=\"main-content\"><p>Test text.</p></div></div>"),
							)
						})
					})
					Context("when without injection scope", func() {
						It("writes", func() {
							t := Tmplt("testTemplate",
								Elem("div",
									[]Node{
										AttrValueInjection("class", "divClass"),
									},
									[]Node{
										Elem("p",
											[]Node{},
											[]Node{
												Injection("paragraph"),
											},
										),
									},
								),
							)
							tmain := Tmplt("mainTemplate",
								Elem("div",
									[]Node{
										Attr("id", "wrapper"),
									},
									[]Node{
										t.Node("test/alias/template", NO_INJECTION_SCOPE),
									},
								),
							)
							nr := tmain.Precompile()
							Expect(nr.Title).To(Equal("TEMPLATE(mainTemplate)"))
							Expect(nr.Messages).To(BeEmpty())
							Expect(nr.Children).To(HaveLen(1))
							ch := nr.Children[0]
							Expect(ch.Title).To(Equal("<div>"))
							Expect(ch.Messages).To(Equal([]string{
								"ok: opening",
								"ok: closing",
							}))
							Expect(ch.Children).To(HaveLen(2))
							ch1 := ch.Children[0]
							Expect(ch1.Title).To(Equal("[]attrs"))
							Expect(ch1.Messages).To(Equal([]string{
								"ok",
							}))
							Expect(ch1.Children).To(BeEmpty())
							ch2 := ch.Children[1]
							Expect(ch2.Title).To(Equal("TEMPLATE(test/alias/template)"))
							Expect(ch2.Messages).To(Equal([]string{
								"ok",
								"ok: injection (divClass)",
								"ok",
								"ok: injection (paragraph)",
								"ok",
							}))
							Expect(ch2.Children).To(BeEmpty())
							Expect(func() {
								tmain.Populate(nil)
							}).To(
								SatisfyAny(
									PanicWith("replacement for \"divClass\" key is not provied"),
									PanicWith("replacement for \"paragraph\" key is not provied"),
								),
							)
							Expect(func() {
								tmain.Populate(map[string]interface{}{
									"divClass": "main-content",
								})
							}).To(
								PanicWith("replacement for \"paragraph\" key is not provied"),
							)
							s := tmain.Populate(map[string]interface{}{
								"divClass":  "main-content",
								"paragraph": "Test text.",
							})
							Expect(s).To(
								Equal("<div id=\"wrapper\"><div class=\"main-content\"><p>Test text.</p></div></div>"),
							)
						})
					})
				})
				Context("when without alias", func() {
					Context("when with injection scope", func() {
						It("writes", func() {
							t := Tmplt("testTemplate",
								Elem("div",
									[]Node{
										AttrValueInjection("class", "divClass"),
									},
									[]Node{
										Elem("p",
											[]Node{},
											[]Node{
												Injection("paragraph"),
											},
										),
									},
								),
							)
							tmain := Tmplt("mainTemplate",
								Elem("div",
									[]Node{
										Attr("id", "wrapper"),
									},
									[]Node{
										t.Node(NO_ALIAS, "testScope"),
									},
								),
							)
							nr := tmain.Precompile()
							Expect(nr.Title).To(Equal("TEMPLATE(mainTemplate)"))
							Expect(nr.Messages).To(BeEmpty())
							Expect(nr.Children).To(HaveLen(1))
							ch := nr.Children[0]
							Expect(ch.Title).To(Equal("<div>"))
							Expect(ch.Messages).To(Equal([]string{
								"ok: opening",
								"ok: closing",
							}))
							Expect(ch.Children).To(HaveLen(2))
							ch1 := ch.Children[0]
							Expect(ch1.Title).To(Equal("[]attrs"))
							Expect(ch1.Messages).To(Equal([]string{
								"ok",
							}))
							Expect(ch1.Children).To(BeEmpty())
							ch2 := ch.Children[1]
							Expect(ch2.Title).To(Equal("TEMPLATE(testTemplate)"))
							Expect(ch2.Messages).To(Equal([]string{
								"ok",
								"ok: injection (testScope.divClass)",
								"ok",
								"ok: injection (testScope.paragraph)",
								"ok",
							}))
							Expect(ch2.Children).To(BeEmpty())
							Expect(func() {
								tmain.Populate(nil)
							}).To(
								SatisfyAny(
									PanicWith("replacement for \"testScope.divClass\" key is not provied"),
									PanicWith("replacement for \"testScope.paragraph\" key is not provied"),
								),
							)
							Expect(func() {
								tmain.Populate(map[string]interface{}{
									"testScope.divClass": "main-content",
								})
							}).To(
								PanicWith("replacement for \"testScope.paragraph\" key is not provied"),
							)
							s := tmain.Populate(map[string]interface{}{
								"testScope.divClass":  "main-content",
								"testScope.paragraph": "Test text.",
							})
							Expect(s).To(
								Equal("<div id=\"wrapper\"><div class=\"main-content\"><p>Test text.</p></div></div>"),
							)
						})
					})
					Context("when without injection scope", func() {
						It("writes", func() {
							t := Tmplt("testTemplate",
								Elem("div",
									[]Node{
										AttrValueInjection("class", "divClass"),
									},
									[]Node{
										Elem("p",
											[]Node{},
											[]Node{
												Injection("paragraph"),
											},
										),
									},
								),
							)
							tmain := Tmplt("mainTemplate",
								Elem("div",
									[]Node{
										Attr("id", "wrapper"),
									},
									[]Node{
										t.Node(NO_ALIAS, NO_INJECTION_SCOPE),
									},
								),
							)
							nr := tmain.Precompile()
							Expect(nr.Title).To(Equal("TEMPLATE(mainTemplate)"))
							Expect(nr.Messages).To(BeEmpty())
							Expect(nr.Children).To(HaveLen(1))
							ch := nr.Children[0]
							Expect(ch.Title).To(Equal("<div>"))
							Expect(ch.Messages).To(Equal([]string{
								"ok: opening",
								"ok: closing",
							}))
							Expect(ch.Children).To(HaveLen(2))
							ch1 := ch.Children[0]
							Expect(ch1.Title).To(Equal("[]attrs"))
							Expect(ch1.Messages).To(Equal([]string{
								"ok",
							}))
							Expect(ch1.Children).To(BeEmpty())
							ch2 := ch.Children[1]
							Expect(ch2.Title).To(Equal("TEMPLATE(testTemplate)"))
							Expect(ch2.Messages).To(Equal([]string{
								"ok",
								"ok: injection (divClass)",
								"ok",
								"ok: injection (paragraph)",
								"ok",
							}))
							Expect(ch2.Children).To(BeEmpty())
							Expect(func() {
								tmain.Populate(nil)
							}).To(
								SatisfyAny(
									PanicWith("replacement for \"divClass\" key is not provied"),
									PanicWith("replacement for \"paragraph\" key is not provied"),
								),
							)
							Expect(func() {
								tmain.Populate(map[string]interface{}{
									"divClass": "main-content",
								})
							}).To(
								PanicWith("replacement for \"paragraph\" key is not provied"),
							)
							s := tmain.Populate(map[string]interface{}{
								"divClass":  "main-content",
								"paragraph": "Test text.",
							})
							Expect(s).To(
								Equal("<div id=\"wrapper\"><div class=\"main-content\"><p>Test text.</p></div></div>"),
							)
						})
					})
				})
			})
		})
	})
})
