package fragments_test

import (
	"github.com/Contra-Culture/go2html/fragments"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("fragments", func() {
	Describe("*Fragments", func() {
		Describe("New()", func() {
			It("returns new fragments", func() {
				Expect(fragments.New([]interface{}{})).NotTo(BeNil())
			})
			Describe(".InContext()", func() {
				It("returns context", func() {
					fs := fragments.New([]interface{}{})
					Expect(
						func() {
							fs.InContext(func(c *fragments.Context) {})
						}).NotTo(Panic())
				})
			})
			Describe(".Fragments()", func() {
				It("returns fragments", func() {
					fs := fragments.New([]interface{}{})
					Expect(fs.Fragments()).NotTo(BeNil())
					Expect(fs.Fragments()).To(BeEmpty())
					fs.InContext(func(c *fragments.Context) {
						c.Append("test1-1")
						c.Append("test1-2")
						c.Append("test1-3")
					})
					fs.InContext(func(c *fragments.Context) {
						c.Append(struct{}{})
						c.Append("test2-1")
						c.Append("test2-2")
						c.Append("test2-3")
						c.InContext(func(c *fragments.Context) {
							c.Append("test2-3-1")
							c.Append("test2-3-2")
							c.Append(struct{}{})
						})
					})
					Expect(fs.Fragments()).NotTo(BeNil())
					Expect(fs.Fragments()).To(HaveLen(4))
					Expect(fs.Fragments()[0].(string)).To(Equal("test1-1test1-2test1-3"))
					Expect(fs.Fragments()[1].(struct{})).To(Equal(struct{}{}))
					Expect(fs.Fragments()[2].(string)).To(Equal("test2-1test2-2test2-3test2-3-1test2-3-2"))
					Expect(fs.Fragments()[3].(struct{})).To(Equal(struct{}{}))
				})
			})
			Describe(".Range()", func() {
				Context("when empty", func() {
					It("returns nil", func() {
						fs := fragments.New([]interface{}{})
						Expect(fs.Range([]int{})).To(BeNil())
						Expect(fs.Range([]int{1, 2, 1})).To(BeNil())
					})
				})
				Context("when not empty", func() {
					It("returns range", func() {
						fs := fragments.New([]interface{}{})
						fs.InContext(func(c *fragments.Context) {
							c.Append("test1-1")
							c.Append("test1-2")
							c.Append("test1-3")
						})
						fs.InContext(func(c *fragments.Context) {
							c.Append(nil)
							c.Append("test2-1")
							c.Append("test2-2")
							c.Append("test2-3")
							c.InContext(func(c *fragments.Context) {
								c.Append("test2-3-1")
								c.Append("test2-3-2")
								c.Append(nil)
							})
						})
						Expect(fs.Range([]int{})).NotTo(BeNil())
						r := fs.Range([]int{})
						Expect(r.BeginFragment).To(Equal(0))
						Expect(r.BeginFragmentPosition).To(Equal(0))
						Expect(r.EndFragment).To(Equal(3))
						Expect(r.EndFragmentPosition).To(Equal(-1))
					})
				})
			})
		})
	})
})
