package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("Template", func() {
		Describe("Spec()", func() {
			It("returns template", func() {
				spec := Spec("testTemplate")
				t, _ := spec.Precompile()
				Expect(t).NotTo(BeNil())
				spec = Spec("testTemplate",
					Elem("p",
						[]Node{},
						[]Node{
							Text("Some text."),
						},
					),
				)
				t, _ = spec.Precompile()
				Expect(t).NotTo(BeNil())
			})
		})
		Describe(".Precompile()", func() {
			It("returns template and report", func() {
				nestedSpec := Spec("nested",
					Elem("span",
						[]Node{
							AttrValueInjection("class", "class"),
						},
						[]Node{
							Injection("text"),
						},
					),
				)
				t, _ := nestedSpec.Precompile()
				spec := Spec("testTemplate",
					Elem("p",
						[]Node{},
						[]Node{
							Text("Some text."),
						},
					),
					t.Node(NO_ALIAS, NO_INJECTION_SCOPE),
				)
				_, rn := spec.Precompile()
				Expect(rn.Title).To(Equal("TEMPLATE(testTemplate)"))
				Expect(rn.Messages).To(BeEmpty())
				Expect(rn.Children).To(HaveLen(2))
				ch1 := rn.Children[0]
				Expect(ch1.Title).To(Equal("<p>"))
				Expect(ch1.Messages).To(Equal([]string{
					"ok: opening",
					"ok: closing",
				}))
				Expect(ch1.Children).To(HaveLen(1))
				ch1_1 := ch1.Children[0]
				Expect(ch1_1.Title).To(Equal("\"text\""))
				Expect(ch1_1.Messages).To(Equal([]string{
					"ok",
				}))
				Expect(ch1_1.Children).To(BeEmpty())
				ch2 := rn.Children[1]
				Expect(ch2.Title).To(Equal("TEMPLATE(nested)"))
				Expect(ch2.Messages).To(Equal([]string{
					"ok",
				}))
				Expect(ch2.Children).To(BeEmpty())
			})
		})
		Describe(".Populate()", func() {
			Context("when params not needed", func() {
				It("populates template", func() {
					spec := Spec("testTemplate",
						Elem("p",
							[]Node{},
							[]Node{
								Text("Some text."),
							},
						),
					)
					t, _ := spec.Precompile()
					s := t.Populate(nil)
					Expect(s).To(Equal("<p>Some text.</p>"))
				})
			})
			Context("when params provided", func() {
				It("populates template", func() {
					nestedSpec := Spec("nested",
						Elem("span",
							[]Node{
								AttrValueInjection("class", "class"),
							},
							[]Node{
								Injection("text"),
							},
						),
					)
					t, _ := nestedSpec.Precompile()
					spec := Spec("testTemplate",
						Elem("p",
							[]Node{},
							[]Node{
								Text("Some text."),
							},
						),
						t.Node(NO_ALIAS, NO_INJECTION_SCOPE),
					)
					t, _ = spec.Precompile()
					s := t.Populate(map[string]interface{}{
						"class": "testClass",
						"text":  "Test text.",
					})
					Expect(s).To(Equal("<p>Some text.</p><span class=\"testClass\">Test text.</span>"))
				})
			})
			Context("when params not provided", func() {
				It("populates template", func() {
					nestedSpec := Spec("nested",
						Elem("span",
							[]Node{
								AttrValueInjection("class", "class"),
							},
							[]Node{
								Injection("text"),
							},
						),
					)
					t, _ := nestedSpec.Precompile()
					spec := Spec("testTemplate",
						Elem("p",
							[]Node{},
							[]Node{
								Text("Some text."),
							},
						),
						t.Node(NO_ALIAS, NO_INJECTION_SCOPE),
					)
					t, _ = spec.Precompile()
					Expect(func() {
						t.Populate(nil)
					}).To(
						SatisfyAny(
							PanicWith("replacement for \"text\" key is not provied"),
							PanicWith("replacement for \"class\" key is not provied"),
						),
					)
				})
			})
		})
	})
	Describe("TemplateNode", func() {
		Describe(".Node()", func() {
			It("returns template node", func() {
				spec := Spec("testTemplate", Elem("div", []Node{}, []Node{}))
				t, _ := spec.Precompile()
				n := t.Node(NO_ALIAS, NO_INJECTION_SCOPE)
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			Context("when with alias", func() {
				It("returns title", func() {
					spec := Spec("testTemplate", Elem("div", []Node{}, []Node{}))
					t, _ := spec.Precompile()
					n := t.Node("some/alias/testTemplate", NO_INJECTION_SCOPE)
					Expect(n.Title()).To(Equal("TEMPLATE(some/alias/testTemplate)"))
				})
			})
			Context("when without alias", func() {
				It("returns title", func() {
					spec := Spec("testTemplate", Elem("div", []Node{}, []Node{}))
					t, _ := spec.Precompile()
					n := t.Node(NO_ALIAS, NO_INJECTION_SCOPE)
					Expect(n.Title()).To(Equal("TEMPLATE(testTemplate)"))
				})
			})
		})
		Describe(".Commit()", func() {
			Context("when with alias", func() {
				Context("when with injection scope", func() {
					It("writes", func() {
						nestedSpec := Spec("testTemplate",
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
						t, _ := nestedSpec.Precompile()
						spec := Spec("mainTemplate",
							Elem("div",
								[]Node{
									Attr("id", "wrapper"),
								},
								[]Node{
									t.Node("test/alias/template", "testScope"),
								},
							),
						)
						t, nr := spec.Precompile()
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
						}))
						Expect(ch2.Children).To(BeEmpty())
						Expect(func() {
							t.Populate(nil)
						}).To(
							PanicWith("replacement for \"testScope\" key is not provied"),
						)
						Expect(func() {
							t.Populate(map[string]interface{}{
								"testScope.divClass": "main-content",
							})
						}).To(
							PanicWith("replacement for \"testScope\" key is not provied"),
						)
						s := t.Populate(map[string]interface{}{
							"testScope": map[string]interface{}{
								"divClass":  "main-content",
								"paragraph": "Test text.",
							},
						})
						Expect(s).To(
							Equal("<div id=\"wrapper\"><div class=\"main-content\"><p>Test text.</p></div></div>"),
						)
					})
				})
				Context("when without injection scope", func() {
					It("writes", func() {
						nestedSpec := Spec("testTemplate",
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
						t, _ := nestedSpec.Precompile()
						spec := Spec("mainTemplate",
							Elem("div",
								[]Node{
									Attr("id", "wrapper"),
								},
								[]Node{
									t.Node("test/alias/template", NO_INJECTION_SCOPE),
								},
							),
						)
						t, nr := spec.Precompile()
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
						}))
						Expect(ch2.Children).To(BeEmpty())
						Expect(func() {
							t.Populate(nil)
						}).To(
							SatisfyAny(
								PanicWith("replacement for \"divClass\" key is not provied"),
								PanicWith("replacement for \"paragraph\" key is not provied"),
							),
						)
						Expect(func() {
							t.Populate(map[string]interface{}{
								"divClass": "main-content",
							})
						}).To(
							PanicWith("replacement for \"paragraph\" key is not provied"),
						)
						s := t.Populate(map[string]interface{}{
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
						nestedSpec := Spec("testTemplate",
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
						t, _ := nestedSpec.Precompile()
						spec := Spec("mainTemplate",
							Elem("div",
								[]Node{
									Attr("id", "wrapper"),
								},
								[]Node{
									t.Node(NO_ALIAS, "testScope"),
								},
							),
						)
						t, nr := spec.Precompile()
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
						}))
						Expect(ch2.Children).To(BeEmpty())
						Expect(func() {
							t.Populate(nil)
						}).To(
							PanicWith("replacement for \"testScope\" key is not provied"),
						)
						Expect(func() {
							t.Populate(map[string]interface{}{
								"testScope": map[string]interface{}{
									"divClass": "main-content",
								},
							})
						}).To(
							PanicWith("replacement for \"paragraph\" key is not provied"),
						)
						s := t.Populate(map[string]interface{}{
							"testScope": map[string]interface{}{
								"divClass":  "main-content",
								"paragraph": "Test text.",
							},
						})
						Expect(s).To(
							Equal("<div id=\"wrapper\"><div class=\"main-content\"><p>Test text.</p></div></div>"),
						)
					})
				})
				Context("when without injection scope", func() {
					It("writes", func() {
						nestedSpec := Spec("testTemplate",
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
						t, _ := nestedSpec.Precompile()
						spec := Spec("mainTemplate",
							Elem("div",
								[]Node{
									Attr("id", "wrapper"),
								},
								[]Node{
									t.Node(NO_ALIAS, NO_INJECTION_SCOPE),
								},
							),
						)
						t, nr := spec.Precompile()
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
						}))
						Expect(ch2.Children).To(BeEmpty())
						Expect(func() {
							t.Populate(nil)
						}).To(
							SatisfyAny(
								PanicWith("replacement for \"divClass\" key is not provied"),
								PanicWith("replacement for \"paragraph\" key is not provied"),
							),
						)
						Expect(func() {
							t.Populate(map[string]interface{}{
								"divClass": "main-content",
							})
						}).To(
							PanicWith("replacement for \"paragraph\" key is not provied"),
						)
						s := t.Populate(map[string]interface{}{
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
