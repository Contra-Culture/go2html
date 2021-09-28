package go2html

import (
	"fmt"
)

type (
	TemplateConfiguringProxy struct {
		template *Template
	}
)

func (tcp *TemplateConfiguringProxy) templateConfiguringProxy() *TemplateConfiguringProxy {
	return &TemplateConfiguringProxy{
		template: tcp.template,
	}
}
func (tcp *TemplateConfiguringProxy) elemConfiguringProxy() *ElemConfiguringProxy {
	return &ElemConfiguringProxy{
		tcp,
	}
}
func (tcp *TemplateConfiguringProxy) appendFragment(fragment interface{}) {
	t := tcp.template
	t.fragments = append(t.fragments, fragment)
}
func (tcp *TemplateConfiguringProxy) Elem(
	name string,
	elemConfig func(*ElemConfiguringProxy),
	nestedNodesConfig func(*NestedNodesConfiguringProxy),
) {
	tcp.appendFragment(fmt.Sprintf("<%s", name))
	elemConfig(&ElemConfiguringProxy{
		tcp: tcp,
	})
	typ := elemTyp(name)
	if typ == VOID_ELEM_TYPE {
		tcp.appendFragment("/>")
		return
	}
	tcp.appendFragment(">")
	nestedNodesConfig(&NestedNodesConfiguringProxy{
		tcp: tcp,
	})
	tcp.appendFragment(fmt.Sprintf("</%s>", name))
}
func (tcp *TemplateConfiguringProxy) Template(key string, t *Template) {
	if len(key) == 0 {
		key = t.key
	}
	tcp.appendFragment(&Template{
		key:       key,
		fragments: t.fragments,
	})
}
func (tcp *TemplateConfiguringProxy) Comment(text string) {
	tcp.appendFragment(fmt.Sprintf("<!-- %s -->", text))
}
func (tcp *TemplateConfiguringProxy) Doctype() {
	tcp.appendFragment("<!DOCTYPE html>")
}
func (tcp *TemplateConfiguringProxy) TextInjection(key string) {
	tcp.appendFragment(injection{
		key: key,
		modifiers: []func(string) string{
			HTMLEscape,
		},
	})
}
func (tcp *TemplateConfiguringProxy) RawTextInjection(key string) {
	tcp.appendFragment(injection{
		key: key,
	})
}
func (tcp *TemplateConfiguringProxy) Text(text string) {
	tcp.appendFragment(safeTextReplacer.Replace(text))
}
func (tcp *TemplateConfiguringProxy) RawText(text string) {
	tcp.appendFragment(text)
}
func (tcp *TemplateConfiguringProxy) Repeat(key string, t *Template) {
	tcp.appendFragment(repetition{
		key:      key,
		template: t,
	})
}
