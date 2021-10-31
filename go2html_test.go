package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("Template", func() {
		Describe("NewTemplate()", func() {
			It("returns template", func() {
				t := NewTemplate("test", func(t *TemplateConfiguringProxy) {
					t.Comment("comment text")
				})
				Expect(t).NotTo(BeNil())
			})
		})
		Describe(".Populate()", func() {
			Context("when all replacements provided", func() {
				It("returns string", func() {
					t := NewTemplate("test", func(t *TemplateConfiguringProxy) {
						t.Comment("comment text")
					})
					Expect(t.Populate(map[string]interface{}{})).To(Equal("<!-- comment text -->"))
					t = NewTemplate("test", func(t *TemplateConfiguringProxy) {
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
								n.UnsafeTextInjection("paragraph1")
							})
						t.Elem(
							"div",
							func(e *ElemConfiguringProxy) {
								e.AttrInjection("div1-attr")
								e.AttrValueInjection("data-ok", "div1-data-ok")
								e.AttrValueInjection("data-confirm", "div1-data-confirm")
							},
							func(n *NestedNodesConfiguringProxy) {
								n.TemplateInjection("templateINJ")
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
									NewTemplate(
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
									NewTemplate(
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
								"templateINJ": map[string]interface{}{
									"template": NewTemplate("templateINJ", func(c *TemplateConfiguringProxy) {
										c.Elem(
											"p",
											func(c *ElemConfiguringProxy) {
												c.AttrValueInjection("class", "value1")
											},
											func(c *NestedNodesConfiguringProxy) {
												c.TextInjection("value2")
											})
									}),
									"values": map[string]interface{}{
										"value1": "Value-1-class",
										"value2": "Value-2-text",
									},
								},
								"span1-text": "Some text here.",
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
					).To(Equal("<!DOCTYPE html><!-- comment text -->Some text here.Inserted text here.<p class=\"paragraph\">Inserted <b>paragraph1</b> text.</p><div title=\"Some title\" data-ok=\"1\" data-confirm=\"1\"><p class=\"Value-1-class\">Value-2-text</p><h1 class=\"div-header\">Header1<span>Some text here.</span></h1><h2 class=\"subheader\"></h2><p>Second &lt;i&gt;paragraph&lt;/i&gt;.</p><p class=\"repeatable-paragraph\">Injected paragraph text 1.</p><p class=\"repeatable-paragraph\">Injected paragraph text 2.</p><p class=\"repeatable-paragraph\">Injected paragraph text 3.</p></div>"))
				})
			})
		})
	})
})
