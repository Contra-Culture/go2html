package registry_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/Contra-Culture/go2html/registry"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("template registry", func() {
	Describe("Registry", func() {
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
		Describe(".Mkdirf()", func() {
			Context("when does not exist", func() {
				It("returns dir", func() {
					r := Reg()
					_, err := r.Mkdirf([]string{"1"}, func(dir Registry) {
						dir.Mkdir([]string{"1-1"})
						dir.Mkdir([]string{"1-2"})
					})
					Expect(err).To(BeNil())
					dir, err := r.Mkdirf([]string{"2"}, func(dir Registry) {
						dir.Mkdir([]string{"1-1"})
						dir.Mkdir([]string{"1-2"})
					})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
				})
			})
			Context("when exists", func() {
				It("returns dir", func() {
					r := Reg()
					_, err := r.Mkdirf([]string{"1"}, func(dir Registry) {
						dir.Mkdir([]string{"1-1"})
					})
					Expect(err).NotTo(HaveOccurred())
					dir, err := r.Mkdirf([]string{"1"}, func(dir Registry) {
						dir.Mkdir([]string{"1-1"})
					})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					_, err = r.Mkdirf([]string{"2"}, func(dir Registry) {
						dir.Mkdirf([]string{"2-1"}, func(dir Registry) {
							dir.Mkdir([]string{"2-1-1"})
						})
					})
					Expect(err).To(BeNil())
					dir, err = r.Mkdirf([]string{"2"}, func(dir Registry) {
						dir.Mkdirf([]string{"2-1"}, func(dir Registry) {
							dir.Mkdir([]string{"2-1-1"})
						})
					})
					Expect(err).NotTo(HaveOccurred())
					Expect(dir).NotTo(BeNil())
					_, err = r.Mkdirf([]string{"2"}, func(dir Registry) {
						dir.Mkdirf([]string{"2-1"}, func(dir Registry) {
							dir.Mkdir([]string{"2-1-2"})
						})
					})
					Expect(err).To(BeNil())
					dir, err = r.Mkdirf([]string{"2"}, func(dir Registry) {
						dir.Mkdirf([]string{"2-1"}, func(dir Registry) {
							dir.Mkdir([]string{"2-1-2"})
						})
					})
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
		Describe(".T()", func() {
			Context("when exists", func() {
				It("fails and returns error", func() {
					r := Reg()
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					err := r.T([]string{"1", "1-1", "1-1-1", "test-template"}, "test", func(t *TemplateCfgr) {
						t.Comment("comment text")
					})
					Expect(err).NotTo(HaveOccurred())
					err = r.T([]string{"1", "1-1", "1-1-1", "test-template"}, "test", func(t *TemplateCfgr) {
						t.Comment("comment text")
					})
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("template \"1/1-1/1-1-1/test-template\" already exists"))
				})
			})
			Context("when does not exist", func() {
				It("adds template", func() {
					r := Reg()
					r.Mkdir([]string{"1", "1-1", "1-1-1"})
					err := r.T([]string{"1", "1-1", "1-1-1", "test-template"}, "test", func(t *TemplateCfgr) {
						t.Comment("comment text")
					})
					Expect(err).NotTo(HaveOccurred())
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
