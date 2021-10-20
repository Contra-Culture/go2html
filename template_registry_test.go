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
					_, err := r.Mkdir([]string{"1"})
					Expect(err).To(BeNil())
					_, err = r.Mkdir([]string{"1", "1-1"})
					Expect(err).NotTo(HaveOccurred())
					_, err = r.Mkdir([]string{"2", "2-1", "2-1-1"})
					Expect(err).NotTo(HaveOccurred())
					dir, err := r.Mkdir([]string{"2", "2-1", "2-1-2"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
				})
			})
			Context("when exists", func() {
				It("returns dir", func() {
					r := Reg("test")
					_, err := r.Mkdir([]string{"1"})
					Expect(err).To(BeNil())
					dir, err := r.Mkdir([]string{"1"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					_, err = r.Mkdir([]string{"1", "1-1"})
					Expect(err).NotTo(HaveOccurred())
					dir, err = r.Mkdir([]string{"1", "1-1"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					_, err = r.Mkdir([]string{"2", "2-1", "2-1-1"})
					Expect(err).To(BeNil())
					dir, err = r.Mkdir([]string{"2", "2-1", "2-1-1"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					_, err = r.Mkdir([]string{"2", "2-1", "2-1-2"})
					Expect(err).To(BeNil())
					dir, err = r.Mkdir([]string{"2", "2-1", "2-1-2"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
				})
			})
		})
		Describe(".Dir()", func() {
			Context("when exists", func() {
				It("returns dir", func() {
					r := Reg("test")
					r.Mkdir([]string{"1"})
					r.Mkdir([]string{"1", "1-1"})
					r.Mkdir([]string{"1", "1-2"})
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					dir, err := r.Dir([]string{"1"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					dir, err = r.Dir([]string{"1", "1-1"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					dir, err = r.Dir([]string{"1", "1-2"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					dir, err = r.Dir([]string{"1", "1-1", "1-1-1"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
				})
			})
			Context("when does not exist", func() {
				It("returns dir", func() {
					r := Reg("test")
					r.Mkdir([]string{"1"})
					r.Mkdir([]string{"1", "1-1"})
					dir, err := r.Dir([]string{"2"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					dir, err = r.Dir([]string{"1", "1-1", "1-1-1"})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
				})
			})
		})
		Describe(".Add()", func() {
			Context("when exists", func() {
				It("fails and returns error", func() {
					r := Reg("test")
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					err := r.Add(&Template{}, []string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).NotTo(HaveOccurred())
					err = r.Add(&Template{}, []string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("template \"1/1-1/1-1-1/test-template\" already exists"))
				})
			})
			Context("when does not exist", func() {
				It("adds template", func() {
					r := Reg("test")
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					err := r.Add(&Template{}, []string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
		Describe(".Get()", func() {
			Context("when exists", func() {
				It("returns template", func() {
					r := Reg("test")
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					err := r.Add(&Template{}, []string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).NotTo(HaveOccurred())
					t, err := r.Get([]string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).NotTo(HaveOccurred())
					Expect(t).NotTo(BeNil())
				})
			})
			Context("when does not exist", func() {
				It("fails and returns error", func() {
					r := Reg("test")
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					t, err := r.Get([]string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("template \"1/1-1/1-1-1/test-template\" does not exist"))
					Expect(t).To(BeNil())
				})
			})
		})
		Describe("TemplateRegistryDir", func() {
			Describe(".Mkdir()", func() {
				Context("when does not exist", func() {
					It("returns dir", func() {
						r := Reg("test")
						dir, err := r.Mkdir([]string{"1"})
						Expect(err).To(BeNil())
						dir, err = dir.Mkdir([]string{"1", "1-1"})
						Expect(err).NotTo(HaveOccurred())
						dir, err = dir.Mkdir([]string{"2", "2-1", "2-1-1"})
						Expect(err).NotTo(HaveOccurred())
						dir, err = dir.Mkdir([]string{"2", "2-1", "2-1-2"})
						Expect(err).NotTo(HaveOccurred())
						Expect(dir).NotTo(BeNil())
					})
				})
				Context("when exists", func() {
					It("returns dir", func() {
						r := Reg("test")
						dir1, err := r.Mkdir([]string{"1"})
						Expect(err).To(BeNil())
						_, err = dir1.Mkdir([]string{"1-1"})
						Expect(err).NotTo(HaveOccurred())
						dir1_1, err := dir1.Mkdir([]string{"1-1"})
						Expect(err).NotTo(HaveOccurred())
						Expect(dir1_1).NotTo(BeNil())
						_, err = dir1.Mkdir([]string{"1-2"})
						Expect(err).To(BeNil())
						dir1_2, err := dir1.Mkdir([]string{"1-2"})
						Expect(err).NotTo(HaveOccurred())
						Expect(dir1_2).NotTo(BeNil())
					})
				})
			})
			Describe(".Dir()", func() {
				Context("when exists", func() {
					It("returns dir", func() {
						r := Reg("test")
						r.Mkdir([]string{"1"})
						r.Mkdir([]string{"1", "1-1"})
						r.Mkdir([]string{"1", "1-2"})
						r.Mkdir([]string{"1", "1-1", "1-1-1"})
						dir, err := r.Dir([]string{"1"})
						Expect(err).NotTo(HaveOccurred())
						_, err = dir.Dir([]string{"1-1"})
						Expect(err).NotTo(HaveOccurred())
						_, err = dir.Dir([]string{"1-2"})
						Expect(err).NotTo(HaveOccurred())
						dir1_1_1, err := dir.Dir([]string{"1-1", "1-1-1"})
						Expect(err).NotTo(HaveOccurred())
						Expect(dir1_1_1).NotTo(BeNil())
					})
				})
				Context("when does not exist", func() {
					It("fails and returns error", func() {
						r := Reg("test")
						dir, _ := r.Mkdir([]string{"1"})
						r.Mkdir([]string{"1", "1-1"})
						dir1_2, err := dir.Dir([]string{"2"})
						Expect(err).NotTo(HaveOccurred())
						Expect(dir1_2).NotTo(BeNil())
						dir1_1_1, err := dir.Dir([]string{"1-1", "1-1-1"})
						Expect(err).NotTo(HaveOccurred())
						Expect(dir1_1_1).NotTo(BeNil())
					})
				})
			})
			Describe(".Add()", func() {
				Context("when exists", func() {
					It("fails and returns error", func() {
						r := Reg("test")
						dir, _ := r.Mkdir([]string{"1", "1-1", "1-1-1"})
						err := dir.Add(&Template{}, []string{"test-template"})
						Expect(err).NotTo(HaveOccurred())
						err = dir.Add(&Template{}, []string{"test-template"})
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("template \"test-template\" already exists"))
					})
				})
				Context("when does not exist", func() {
					It("adds template", func() {
						r := Reg("test")
						dir, _ := r.Mkdir([]string{"1", "1-1", "1-1-1"})
						err := dir.Add(&Template{}, []string{"test-template"})
						Expect(err).NotTo(HaveOccurred())
					})
				})
			})
			Describe(".Get()", func() {
				Context("when exists", func() {
					It("returns template", func() {
						r := Reg("test")
						dir, _ := r.Mkdir([]string{"1", "1-1", "1-1-1"})
						err := r.Add(&Template{}, []string{"1", "1-1", "1-1-1", "test-template"})
						Expect(err).NotTo(HaveOccurred())
						t, err := dir.Get([]string{"test-template"})
						Expect(err).NotTo(HaveOccurred())
						Expect(t).NotTo(BeNil())
					})
				})
				Context("when does not exist", func() {
					It("fails and returns error", func() {
						r := Reg("test")
						dir, _ := r.Mkdir([]string{"1", "1-1", "1-1-1"})
						t, err := dir.Get([]string{"test-template"})
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("template \"test-template\" does not exist"))
						Expect(t).To(BeNil())
					})
				})
			})
		})
	})
})
