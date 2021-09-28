package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("Template", func() {
		Describe("Tmplt()", func() {
			It("returns template", func() {
				t := Tmplt("test", func(t *TemplateConfiguringProxy) {
					t.Comment("comment text")
				})
				Expect(t).NotTo(BeNil())
			})
		})
		Describe(".Populate()", func() {
			Context("when all replacements provided", func() {
				It("returns string", func() {
					t := Tmplt("test", func(t *TemplateConfiguringProxy) {
						t.Comment("comment text")
					})
					Expect(t.Populate(map[string]interface{}{})).To(Equal("<!-- comment text -->"))

					t = Tmplt("test", func(t *TemplateConfiguringProxy) {
						t.Doctype()
						t.Comment("comment text")
						t.Text("Some text here.")
						t.TextInjection("text1")
						t.Elem(
							"p",
							func(e *ElemConfiguringProxy) {
								e.Attr("class", "paragraph")
							},
							func(n *NestedNodesConfiguringProxy) {
								n.RawTextInjection("paragraph1")
							})
						t.Elem(
							"div",
							func(e *ElemConfiguringProxy) {
								e.AttrInjection("div1-attr")
								e.AttrValueInjection("data-ok", "div1-data-ok")
								e.AttrValueInjection("data-confirm", "div1-data-confirm")
							},
							func(n *NestedNodesConfiguringProxy) {
								n.Elem(
									"h1",
									func(e *ElemConfiguringProxy) {
										e.Attr("class", "div-header")
									},
									func(n *NestedNodesConfiguringProxy) {
										n.Text("Header1")
										n.Elem(
											"span",
											func(e *ElemConfiguringProxy) {},
											func(n *NestedNodesConfiguringProxy) {
												n.TextInjection("span1-text")
											},
										)
									})
								n.Template(
									"",
									Tmplt(
										"nestedTemplate1",
										func(t *TemplateConfiguringProxy) {
											t.Elem(
												"h2",
												func(e *ElemConfiguringProxy) {
													e.AttrValueInjection("class", "header2-class")
												},
												func(n *NestedNodesConfiguringProxy) {},
											)
											t.Elem(
												"p",
												func(*ElemConfiguringProxy) {},
												func(n *NestedNodesConfiguringProxy) {
													n.TextInjection("paragraph2-text")
												},
											)
										}))
								n.Repeat(
									"paragraphs",
									Tmplt(
										"nestedTemplate1",
										func(t *TemplateConfiguringProxy) {
											t.Elem(
												"p",
												func(e *ElemConfiguringProxy) {
													e.Attr("class", "repeatable-paragraph")
												},
												func(n *NestedNodesConfiguringProxy) {
													n.TextInjection("paragraph-text")
												},
											)
										},
									))
							})
					})
					Expect(
						t.Populate(
							map[string]interface{}{
								"text1":             "Inserted text here.",
								"paragraph1":        "Inserted <b>paragraph1</b> text.",
								"div1-attr":         "title=\"Some title\"",
								"div1-data-ok":      "1",
								"div1-data-confirm": "1",
								"span1-text":        "Some text here.",
								"nestedTemplate1": map[string]interface{}{
									"header2-class":   "subheader",
									"paragraph2-text": "Second <i>paragraph</i>.",
								},
								"paragraphs": []map[string]interface{}{
									{"paragraph-text": "Injected paragraph text 1."},
									{"paragraph-text": "Injected paragraph text 2."},
									{"paragraph-text": "Injected paragraph text 3."},
								},
							},
						),
					).To(Equal("<!DOCTYPE html><!-- comment text -->Some text here.Inserted text here.<p class=\"paragraph\">Inserted <b>paragraph1</b> text.</p><div title=\"Some title\" data-ok=\"1\" data-confirm=\"1\"><h1 class=\"div-header\">Header1<span>Some text here.</span></h1><h2 class=\"subheader\"></h2><p>Second &lt;i&gt;paragraph&lt;/i&gt;.</p><p class=\"repeatable-paragraph\">Injected paragraph text 1.</p><p class=\"repeatable-paragraph\">Injected paragraph text 2.</p><p class=\"repeatable-paragraph\">Injected paragraph text 3.</p></div>"))
				})
			})
		})
	})
})
