package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("template registry", func() {
	Describe("TemplateRegistry", func() {
		Describe("Reg()", func() {
			It("returns registry", func() {
				r := Reg("test")
				Expect(r).NotTo(BeNil())
			})
		})
		Describe(".Mkdir()", func() {
			Context("when does not exist", func() {
				It("returns dir", func() {
					r := Reg("test")
					dir, err := r.Mkdir("1")
					Expect(err).To(BeNil())
					dir, err = r.Mkdir("1", "1-1")
					Expect(err).NotTo(HaveOccurred())
					dir, err = r.Mkdir("2", "2-1", "2-1-1")
					Expect(err).NotTo(HaveOccurred())
					dir, err = r.Mkdir("2", "2-1", "2-1-2")
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
				})
			})
			Context("when exists", func() {
				It("fails and returns error", func() {
					r := Reg("test")
					dir, err := r.Mkdir("1")
					Expect(err).To(BeNil())
					dir, err = r.Mkdir("1")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("directory \"1\" already exists"))
					Expect(dir).To(BeNil())
					dir, err = r.Mkdir("1", "1-1")
					Expect(err).NotTo(HaveOccurred())
					dir, err = r.Mkdir("1", "1-1")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("directory \"1-1\" already exists"))
					Expect(dir).To(BeNil())
					dir, err = r.Mkdir("2", "2-1", "2-1-1")
					Expect(err).To(BeNil())
					dir, err = r.Mkdir("2", "2-1", "2-1-1")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("directory \"2-1-1\" already exists"))
					Expect(dir).To(BeNil())
					dir, err = r.Mkdir("2", "2-1", "2-1-2")
					Expect(err).To(BeNil())
					dir, err = r.Mkdir("2", "2-1", "2-1-2")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("directory \"2-1-2\" already exists"))
					Expect(dir).To(BeNil())
				})
			})
		})
		Describe(".Dir()", func() {
			Context("when exists", func() {
				It("returns dir", func() {
					r := Reg("test")
					r.Mkdir("1")
					r.Mkdir("1", "1-1")
					r.Mkdir("1", "1-2")
					r.Mkdir("1", "1-1", "1-1-1")
					dir, err := r.Dir("1")
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					dir, err = r.Dir("1", "1-1")
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					dir, err = r.Dir("1", "1-2")
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					dir, err = r.Dir("1", "1-1", "1-1-1")
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
				})
			})
			Context("when does not exist", func() {
				It("fails and returns error", func() {
					r := Reg("test")
					r.Mkdir("1")
					r.Mkdir("1", "1-1")
					dir, err := r.Dir("2")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("wrong path, dir \"2\" is not found"))
					Expect(dir).To(BeNil())
					dir, err = r.Dir("1", "1-1", "1-1-1")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("wrong path, dir \"1-1-1\" is not found"))
					Expect(dir).To(BeNil())
				})
			})
		})
		Describe(".Add()", func() {
			Context("when exists", func() {
				It("fails and returns error", func() {
					r := Reg("test")
					r.Mkdir("1", "1-1", "1-1-1")
					err := r.Add(&Template{}, "1", "1-1", "1-1-1", "test-template")
					Expect(err).NotTo(HaveOccurred())
					err = r.Add(&Template{}, "1", "1-1", "1-1-1", "test-template")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("template \"1/1-1/1-1-1/test-template\" already exists"))
				})
			})
			Context("when does not exist", func() {
				It("adds template", func() {
					r := Reg("test")
					r.Mkdir("1", "1-1", "1-1-1")
					err := r.Add(&Template{}, "1", "1-1", "1-1-1", "test-template")
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
		Describe(".Get()", func() {
			Context("when exists", func() {
				It("returns template", func() {
					r := Reg("test")
					r.Mkdir("1", "1-1", "1-1-1")
					err := r.Add(&Template{}, "1", "1-1", "1-1-1", "test-template")
					Expect(err).NotTo(HaveOccurred())
					t, err := r.Get("1", "1-1", "1-1-1", "test-template")
					Expect(err).NotTo(HaveOccurred())
					Expect(t).NotTo(BeNil())
				})
			})
			Context("when does not exist", func() {
				It("fails and returns error", func() {
					r := Reg("test")
					r.Mkdir("1", "1-1", "1-1-1")
					t, err := r.Get("1", "1-1", "1-1-1", "test-template")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("template \"1/1-1/1-1-1/test-template\" does not exist"))
					Expect(t).To(BeNil())
				})
			})
		})
		Describe(".TemplateNode()", func() {
			Context("when template exists", func() {
				It("returns template node", func() {
					reg := Reg("testReg")
					reg.Mkdir("some")
					reg.Mkdir("some", "path")
					reg.Mkdir("some", "path", "here")
					err := reg.Add(Tmplt("test",
						Elem("div",
							[]Node{
								Attr("id", "myID"),
								AttrValueInjection("class", "test-class"),
							},
							[]Node{
								Elem("p", []Node{}, []Node{
									Elem("span", []Node{}, []Node{
										Injection("nested-template-text"),
									}),
								}),
								Injection("text"),
							})),
						"some", "path", "here", "testTemplate",
					)
					Expect(err).NotTo(HaveOccurred())
					elem := Elem("div", []Node{}, []Node{reg.TemplateNode("some", "path", "here", "testTemplate")})
					tmain := Tmplt("test", elem)
					r := tmain.Precompile()
					Expect(r.Title).To(Equal("TEMPLATE(test)"))
					Expect(r.Messages).To(BeEmpty())
					Expect(r.Children).To(HaveLen(1))
					Expect(r.Children).To(HaveLen(1))
					ch1 := r.Children[0]
					Expect(ch1.Title).To(Equal("<div>"))
					Expect(ch1.Messages).To(Equal([]string{
						"ok: opening",
						"ok: closing",
					}))
					Expect(ch1.Children).To(HaveLen(1))
					ch1_1 := ch1.Children[0]
					Expect(ch1_1.Title).To(Equal("TEMPLATE(some/path/here/testTemplate)"))
					Expect(ch1_1.Messages).To(Equal([]string{
						"ok",
						"ok: injection (test.test-class)",
						"ok",
						"ok: injection (test.nested-template-text)",
						"ok",
						"ok: injection (test.text)",
						"ok",
					}))
					Expect(ch1_1.Children).To(BeEmpty())
					str := tmain.Populate(map[string]interface{}{
						"test.nested-template-text": "NESTED TEMPLATE",
						"test.test-class":           "testClass",
						"test.text":                 "testText",
					})
					Expect(str).To(Equal("<div><div id=\"myID\" class=\"testClass\"><p><span>NESTED TEMPLATE</span></p>testText</div></div>"))
				})
			})
			Context("when template does not exist", func() {

			})
		})
		Describe("TemplateRegistryDir", func() {
			Describe(".Mkdir()", func() {
				Context("when does not exist", func() {
					It("returns dir", func() {
						r := Reg("test")
						dir, err := r.Mkdir("1")
						Expect(err).To(BeNil())
						dir, err = dir.Mkdir("1", "1-1")
						Expect(err).NotTo(HaveOccurred())
						dir, err = dir.Mkdir("2", "2-1", "2-1-1")
						Expect(err).NotTo(HaveOccurred())
						dir, err = dir.Mkdir("2", "2-1", "2-1-2")
						Expect(err).NotTo(HaveOccurred())
						Expect(dir).NotTo(BeNil())
					})
				})
				Context("when exists", func() {
					It("fails and returns error", func() {
						r := Reg("test")
						dir1, err := r.Mkdir("1")
						Expect(err).To(BeNil())
						dir1_1, err := dir1.Mkdir("1-1")
						Expect(err).NotTo(HaveOccurred())
						dir1_1, err = dir1.Mkdir("1-1")
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("directory \"1-1\" already exists"))
						Expect(dir1_1).To(BeNil())
						dir1_2, err := dir1.Mkdir("1-2")
						Expect(err).To(BeNil())
						dir1_2, err = dir1.Mkdir("1-2")
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("directory \"1-2\" already exists"))
						Expect(dir1_2).To(BeNil())
					})
				})
			})
			Describe(".Dir()", func() {
				Context("when exists", func() {
					It("returns dir", func() {
						r := Reg("test")
						r.Mkdir("1")
						r.Mkdir("1", "1-1")
						r.Mkdir("1", "1-2")
						r.Mkdir("1", "1-1", "1-1-1")
						dir, err := r.Dir("1")
						Expect(err).NotTo(HaveOccurred())
						_, err = dir.Dir("1-1")
						Expect(err).NotTo(HaveOccurred())
						_, err = dir.Dir("1-2")
						Expect(err).NotTo(HaveOccurred())
						dir1_1_1, err := dir.Dir("1-1", "1-1-1")
						Expect(err).NotTo(HaveOccurred())
						Expect(dir1_1_1).NotTo(BeNil())
					})
				})
				Context("when does not exist", func() {
					It("fails and returns error", func() {
						r := Reg("test")
						dir, _ := r.Mkdir("1")
						r.Mkdir("1", "1-1")
						dir1_2, err := dir.Dir("2")
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("wrong path, dir \"2\" is not found"))
						Expect(dir1_2).To(BeNil())
						dir1_1_1, err := dir.Dir("1-1", "1-1-1")
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("wrong path, dir \"1-1-1\" is not found"))
						Expect(dir1_1_1).To(BeNil())
					})
				})
			})
			Describe(".Add()", func() {
				Context("when exists", func() {
					It("fails and returns error", func() {
						r := Reg("test")
						dir, _ := r.Mkdir("1", "1-1", "1-1-1")
						err := dir.Add(&Template{}, "test-template")
						Expect(err).NotTo(HaveOccurred())
						err = dir.Add(&Template{}, "test-template")
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("template \"test-template\" already exists"))
					})
				})
				Context("when does not exist", func() {
					It("adds template", func() {
						r := Reg("test")
						dir, _ := r.Mkdir("1", "1-1", "1-1-1")
						err := dir.Add(&Template{}, "test-template")
						Expect(err).NotTo(HaveOccurred())
					})
				})
			})
			Describe(".Get()", func() {
				Context("when exists", func() {
					It("returns template", func() {
						r := Reg("test")
						dir, _ := r.Mkdir("1", "1-1", "1-1-1")
						err := r.Add(&Template{}, "1", "1-1", "1-1-1", "test-template")
						Expect(err).NotTo(HaveOccurred())
						t, err := dir.Get("test-template")
						Expect(err).NotTo(HaveOccurred())
						Expect(t).NotTo(BeNil())
					})
				})
				Context("when does not exist", func() {
					It("fails and returns error", func() {
						r := Reg("test")
						dir, _ := r.Mkdir("1", "1-1", "1-1-1")
						t, err := dir.Get("test-template")
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("template \"test-template\" does not exist"))
						Expect(t).To(BeNil())
					})
				})
			})
		})
	})
})
