package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("template registry", func() {
	Describe("TemplatesReg", func() {
		Describe(".Mkdir()", func() {
			Context("when does not exist", func() {
				It("returns dir", func() {
					r := Reg()
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
					r := Reg()
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
					r := Reg()
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
				It("fails and returns error", func() {
					r := Reg()
					r.Mkdir([]string{"1"})
					r.Mkdir([]string{"1", "1-1"})
					dir, err := r.Dir([]string{"2"})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("wrong path, dir \"2\" is not found"))
					Expect(dir).To(BeNil())
					dir, err = r.Dir([]string{"1", "1-1", "1-1-1"})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("wrong path, dir \"1-1-1\" is not found"))
					Expect(dir).To(BeNil())
				})
			})
		})
		Describe(".Add()", func() {
			Context("when exists", func() {
				It("fails and returns error", func() {
					r := Reg()
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
					r := Reg()
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					err := r.Add(&Template{}, []string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
		Describe(".Get()", func() {
			Context("when exists", func() {
				It("returns template", func() {
					r := Reg()
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
					r := Reg()
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					t, err := r.Get([]string{"1", "1-1", "1-1-1", "test-template"})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("template \"1/1-1/1-1-1/test-template\" does not exist"))
					Expect(t).To(BeNil())
				})
			})
		})
	})
})
