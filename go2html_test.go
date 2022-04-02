package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("Template", func() {
		Describe("NewTemplate()", func() {
			It("returns template", func() {
				t := NewTemplate("test", func(t *TemplateCfgr) {
					t.Comment("comment text")
				})
				Expect(t).NotTo(BeNil())
			})
		})
		Describe(".Populate()", func() {
			Context("when all replacements provided", func() {
				It("returns string", func() {
					t := NewTemplate("test", func(t *TemplateCfgr) {
						t.Comment("comment text")
					})
					Expect(t.Populate(map[string]interface{}{})).To(Equal("<!-- comment text -->"))
					t = NewTemplate("test", func(t *TemplateCfgr) {
						t.Doctype()
						t.Comment("comment text")
						t.Text("Some text here.")
						t.TextInjection("text1")
						t.Elem(
							"p",
							func(e *ElemCfgr) {
								e.Class("p-normal")
								e.Attr("class", "paragraph")
							},
							func(n *NestedNodesCfgr) {
								n.UnsafeTextInjection("paragraph1")
							})
						t.Elem(
							"div",
							func(e *ElemCfgr) {
								e.ID("wrapper")
								e.AttrsInjection("div1-attr")
								e.AttrValueInjection("data-ok", "div1-data-ok")
								e.AttrValueInjection("data-confirm", "div1-data-confirm")
							},
							func(n *NestedNodesCfgr) {
								n.Variants(
									map[string]*Template{
										"div_nested_1": NewTemplate(
											"",
											func(cfg *TemplateCfgr) {
												cfg.TextInjection("text")
											},
										),
										"div_nested_2": NewTemplate(
											"",
											func(cfg *TemplateCfgr) {
												cfg.TextInjection("anotherText")
											},
										),
									},
									NewTemplate(
										"",
										func(cfg *TemplateCfgr) {
											cfg.Text("no variant provided")
										},
									))
								n.TemplateInjection("templateINJ")
								n.Elem(
									"h1",
									func(e *ElemCfgr) {
										e.Attr("class", "div-header")
									},
									func(n *NestedNodesCfgr) {
										n.Text("Header1")
										n.Elem(
											"span",
											func(e *ElemCfgr) {},
											func(n *NestedNodesCfgr) {
												n.TextInjection("span1-text")
											})
									})
								n.Template(
									"",
									NewTemplate(
										"nestedTemplate1",
										func(t *TemplateCfgr) {
											t.Variants(
												map[string]*Template{
													"nested_nested_1": NewTemplate(
														"",
														func(cfg *TemplateCfgr) {
															cfg.TextInjection("text")
														}),
													"nested_nested_2": NewTemplate(
														"",
														func(cfg *TemplateCfgr) {
															cfg.TextInjection("anotherText")
														}),
												},
												NewTemplate(
													"",
													func(cfg *TemplateCfgr) {
														cfg.Text("no variant provided")
													},
												))
											t.Elem(
												"h2",
												func(e *ElemCfgr) {
													e.AttrValueInjection("class", "header2-class")
												},
												func(n *NestedNodesCfgr) {},
											)
											t.Elem(
												"p",
												func(*ElemCfgr) {},
												func(n *NestedNodesCfgr) {
													n.TextInjection("paragraph2-text")
												},
											)
										}))
								n.Repeat(
									"paragraphs",
									NewTemplate(
										"nestedTemplate2",
										func(t *TemplateCfgr) {
											t.Elem(
												"p",
												func(e *ElemCfgr) {
													e.Attr("class", "repeatable-paragraph")
												},
												func(n *NestedNodesCfgr) {
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
								"div1-attr":         map[string]string{"title": "Some title"},
								"div1-data-ok":      "1",
								"div1-data-confirm": "1",
								"div_nested_1": map[string]interface{}{
									"text": "variant div_nested_1 text",
								},
								"templateINJ": map[string]interface{}{
									"template": NewTemplate("templateINJ", func(c *TemplateCfgr) {
										c.Elem(
											"p",
											func(c *ElemCfgr) {
												c.AttrValueInjection("class", "value1")
											},
											func(c *NestedNodesCfgr) {
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
					).To(Equal("<!DOCTYPE html><!-- comment text -->Some text here.Inserted text here.<p class=\"p-normal\" class=\"paragraph\">Inserted <b>paragraph1</b> text.</p><div id=\"wrapper\" title=\"Some title\" data-ok=\"1\" data-confirm=\"1\">variant div_nested_1 text<p class=\"Value-1-class\">Value-2-text</p><h1 class=\"div-header\">Header1<span>Some text here.</span></h1>no variant provided<h2 class=\"subheader\"></h2><p>Second &lt;i&gt;paragraph&lt;/i&gt;.</p><p class=\"repeatable-paragraph\">Injected paragraph text 1.</p><p class=\"repeatable-paragraph\">Injected paragraph text 2.</p><p class=\"repeatable-paragraph\">Injected paragraph text 3.</p></div>"))
				})
			})
		})
	})
})
