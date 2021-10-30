package go2html

import (
	"fmt"
)

type (
	TemplateConfiguringProxy struct {
		template *Template
	}
)

func (tcp *TemplateConfiguringProxy) appendFragment(rawNewFragment interface{}) FragmentPosition {
	t := tcp.template
	newFragment, ok := rawNewFragment.(string)
	if !ok {
		t.fragments = append(t.fragments, rawNewFragment)
		return FragmentPosition{
			FragmentIndex: len(t.fragments) - 1,
			RangeBegin:    0,
			RangeEnd:      0,
		}
	}
	lastFragmentIdx := len(t.fragments) - 1
	if len(t.fragments) == 0 {
		t.fragments = append(t.fragments, rawNewFragment)
		return FragmentPosition{
			FragmentIndex: 0,
			RangeBegin:    0,
			RangeEnd:      len(newFragment) - 1,
		}
	}
	rawLastFragment := t.fragments[lastFragmentIdx]
	lastFragment, ok := rawLastFragment.(string)
	if !ok {
		t.fragments = append(t.fragments, rawNewFragment)
		return FragmentPosition{
			FragmentIndex: len(t.fragments) - 1,
			RangeBegin:    0,
			RangeEnd:      0,
		}
	}
	t.fragments[lastFragmentIdx] = lastFragment + newFragment
	return FragmentPosition{
		FragmentIndex: len(t.fragments) - 1,
		RangeBegin:    len(lastFragment),
		RangeEnd:      len(lastFragment) + len(newFragment) - 1,
	}
}
func (tcp *TemplateConfiguringProxy) Elem(
	name string,
	configureSelf func(*ElemConfiguringProxy),
	configureNested func(*NestedNodesConfiguringProxy),
) {

	posBegin := tcp.appendFragment(fmt.Sprintf("<%s", name))
	node := &Node{
		PosBegin: posBegin,
		Kind:     ELEM_NODE_KIND,
		Title:    name,
		Children: []*Node{},
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	configureSelf(&ElemConfiguringProxy{
		tcp:  tcp,
		node: node,
	})
	typ := elemTyp(name)
	if typ == VOID_ELEM_TYPE {
		node.PosEnd = tcp.appendFragment("/>")
		return
	}
	tcp.appendFragment(">")
	configureNested(&NestedNodesConfiguringProxy{
		tcp:    tcp,
		parent: node,
	})
	node.PosEnd = tcp.appendFragment(fmt.Sprintf("</%s>", name))
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
	pos := tcp.appendFragment(fmt.Sprintf("<!-- %s -->", text))
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     COMMENT_NODE_KIND,
		Title:    COMMENT_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
}
func (tcp *TemplateConfiguringProxy) Doctype() {
	pos := tcp.appendFragment("<!DOCTYPE html>")
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     DOCTYPE_NODE_KIND,
		Title:    DOCTYPE_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
}
func (tcp *TemplateConfiguringProxy) TextInjection(key string) {
	pos := tcp.appendFragment(injection{
		key: key,
		modifiers: []func(string) string{
			HTMLEscape,
		},
	})
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEXT_INJECTION_NODE_KIND,
		Title:    key,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
}
func (tcp *TemplateConfiguringProxy) UnsafeTextInjection(key string) {
	pos := tcp.appendFragment(injection{
		key: key,
	})
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEXT_INJECTION_NODE_KIND,
		Title:    key,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
}
func (tcp *TemplateConfiguringProxy) Text(text string) {
	pos := tcp.appendFragment(safeTextReplacer.Replace(text))
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEXT_NODE_KIND,
		Title:    TEXT_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
}
func (tcp *TemplateConfiguringProxy) UnsafeText(text string) {
	pos := tcp.appendFragment(text)
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEXT_NODE_KIND,
		Title:    TEXT_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
}
func (tcp *TemplateConfiguringProxy) Repeat(key string, t *Template) {
	pos := tcp.appendFragment(repetition{
		key:      key,
		template: t,
	})
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     REPEAT_NODE_KIND,
		Title:    key,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
}
