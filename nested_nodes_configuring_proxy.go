package go2html

import "fmt"

type (
	NestedNodesConfiguringProxy struct {
		tcp *TemplateConfiguringProxy
	}
)

func (nncp *NestedNodesConfiguringProxy) Elem(
	name string,
	elemConfig func(*ElemConfiguringProxy),
	nestedNodesConfig func(*NestedNodesConfiguringProxy),
) {
	nncp.tcp.appendFragment(fmt.Sprintf("<%s", name))
	elemConfig(&ElemConfiguringProxy{
		tcp: nncp.tcp,
	})
	typ := elemTyp(name)
	if typ == VOID_ELEM_TYPE {
		nncp.tcp.appendFragment("/>")
		return
	}
	nncp.tcp.appendFragment(">")
	nestedNodesConfig(&NestedNodesConfiguringProxy{
		tcp: nncp.tcp,
	})
	nncp.tcp.appendFragment(fmt.Sprintf("</%s>", name))
}
func (nncp *NestedNodesConfiguringProxy) Template(key string, t *Template) {
	if len(key) == 0 {
		key = t.key
	}
	nncp.tcp.appendFragment(&Template{
		key:       key,
		fragments: t.fragments,
	})
}
func (nncp *NestedNodesConfiguringProxy) Comment(text string) {
	nncp.tcp.appendFragment(fmt.Sprintf("<!-- %s -->", text))
}
func (nncp *NestedNodesConfiguringProxy) Doctype() {
	nncp.tcp.appendFragment("<!DOCTYPE html>")
}
func (nncp *NestedNodesConfiguringProxy) TextInjection(key string) {
	nncp.tcp.appendFragment(injection{
		key: key,
		modifiers: []func(string) string{
			HTMLEscape,
		},
	})
}
func (nncp *NestedNodesConfiguringProxy) RawTextInjection(key string) {
	nncp.tcp.appendFragment(injection{
		key: key,
	})
}
func (nncp *NestedNodesConfiguringProxy) Text(text string) {
	nncp.tcp.appendFragment(safeTextReplacer.Replace(text))
}
func (nncp *NestedNodesConfiguringProxy) RawText(text string) {
	nncp.tcp.appendFragment(text)
}
func (nncp *NestedNodesConfiguringProxy) Repeat(key string, t *Template) {
	nncp.tcp.appendFragment(repetition{
		key:      key,
		template: t,
	})
}
