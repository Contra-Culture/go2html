package go2html_test

import (
	. "github.com/Contra-Culture/go2html"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2html", func() {
	Describe("RepNode", func() {
		Describe("Rep()", func() {
			It("returns repeatable node", func() {
				spec := Spec("test", Text("ok"))
				t, _ := spec.Precompile()
				n := Rep("manySomething", "repTest", t.Node(NO_ALIAS, NO_INJECTION_SCOPE))
				Expect(n).NotTo(BeNil())
			})
		})
		Describe(".Title()", func() {
			It("returns title", func() {
				spec := Spec("test", Text("ok"))
				t, _ := spec.Precompile()
				n := Rep("manySomething", "repTest", t.Node(NO_ALIAS, NO_INJECTION_SCOPE))
				Expect(n.Title()).To(Equal("{manySomething}*"))
			})
		})
		Describe(".Commit()", func() {
			It("commits", func() {
				repSpec := Spec("repeatable", Text("ok"))
				t, _ := repSpec.Precompile()
				spec := Spec("testTemplate", Rep("repTest", "repScope", t.Node(NO_ALIAS, NO_INJECTION_SCOPE)))
				t, rn := spec.Precompile()
				Expect(t).NotTo(BeNil())
				Expect(rn.Title).To(Equal("TEMPLATE(testTemplate)"))
				Expect(rn.Messages).To(BeEmpty())
				Expect(rn.Children).To(HaveLen(1))
				ch := rn.Children[0]
				Expect(ch.Title).To(Equal("{repTest}*"))
				Expect(ch.Messages).To(Equal([]string{
					"ok",
				}))
				Expect(ch.Children).To(BeEmpty())
			})
		})
		Describe(".Populate()", func() {
			spec := Spec("test", Injection("injtext"))
			t, _ := spec.Precompile()
			n := Rep("manySomething", "repTest", t.Node(NO_ALIAS, NO_INJECTION_SCOPE))

			Context("when without injections", func() {
				It("returns string", func() {
					s := n.Populate(nil)
					Expect(s).To(Equal(""))
				})
			})
			Context("when with injections", func() {
				It("returns string", func() {
					Expect(func() {
						n.Populate([]interface{}{
							map[string]interface{}{"injtext": "test1"},
							map[string]interface{}{"injtext": "test2"},
							map[string]interface{}{"wronginj": "test3"},
						})
					}).To(
						PanicWith("replacement for \"injtext\" key is not provied"),
					)
					s := n.Populate([]interface{}{
						map[string]interface{}{"injtext": "test1"},
						map[string]interface{}{"injtext": "test2"},
						map[string]interface{}{"injtext": "test3"},
					})
					Expect(s).To(Equal("test1test2test3"))
				})
			})
		})
		Describe(".Scope()", func() {
			It("returns scope", func() {
				spec := Spec("test", Text("ok"))
				t, _ := spec.Precompile()
				n := Rep("manySomething", "repTest", t.Node(NO_ALIAS, NO_INJECTION_SCOPE))
				scope, ok := n.Scope()
				Expect(scope).To(Equal("repTest"))
				Expect(ok).To(BeTrue())
			})
		})
	})
})
