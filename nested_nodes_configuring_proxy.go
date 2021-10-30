package go2html

import "fmt"

type (
	NestedNodesConfiguringProxy struct {
		tcp    *TemplateConfiguringProxy
		parent *Node
	}
)

func (nncp *NestedNodesConfiguringProxy) Elem(
	name string,
	configureSelf func(*ElemConfiguringProxy),
	configureNested func(*NestedNodesConfiguringProxy),
) {
	posBegin := nncp.tcp.appendFragment(fmt.Sprintf("<%s", name))
	node := &Node{
		PosBegin: posBegin,
		Kind:     ELEM_NODE_KIND,
		Title:    name,
		Children: []*Node{},
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
	configureSelf(&ElemConfiguringProxy{
		tcp:  nncp.tcp,
		node: node,
	})
	typ := elemTyp(name)
	if typ == VOID_ELEM_TYPE {
		node.PosEnd = nncp.tcp.appendFragment("/>")
		return
	}
	nncp.tcp.appendFragment(">")
	configureNested(&NestedNodesConfiguringProxy{
		tcp:    nncp.tcp,
		parent: node,
	})
	node.PosEnd = nncp.tcp.appendFragment(fmt.Sprintf("</%s>", name))
}
func (nncp *NestedNodesConfiguringProxy) Template(key string, t *Template) {
	if len(key) == 0 {
		key = t.key
	}
	pos := nncp.tcp.appendFragment(&Template{
		key:       key,
		nodes:     t.nodes,
		fragments: t.fragments,
	})
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEMPLATE_NODE_KIND,
		Title:    key,
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
}
func (nncp *NestedNodesConfiguringProxy) Comment(text string) {
	pos := nncp.tcp.appendFragment(fmt.Sprintf("<!-- %s -->", text))
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     COMMENT_NODE_KIND,
		Title:    COMMENT_NODE_TITLE,
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
}
func (nncp *NestedNodesConfiguringProxy) Doctype() {
	pos := nncp.tcp.appendFragment("<!DOCTYPE html>")
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     DOCTYPE_NODE_KIND,
		Title:    DOCTYPE_NODE_TITLE,
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
}
func (nncp *NestedNodesConfiguringProxy) TextInjection(key string) {
	pos := nncp.tcp.appendFragment(injection{
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
	nncp.parent.Children = append(nncp.parent.Children, node)
}
func (nncp *NestedNodesConfiguringProxy) UnsafeTextInjection(key string) {
	pos := nncp.tcp.appendFragment(injection{
		key: key,
	})
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEXT_INJECTION_NODE_KIND,
		Title:    key,
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
}
func (nncp *NestedNodesConfiguringProxy) Text(text string) {
	pos := nncp.tcp.appendFragment(safeTextReplacer.Replace(text))
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEXT_NODE_KIND,
		Title:    TEXT_NODE_TITLE,
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
}
func (nncp *NestedNodesConfiguringProxy) UnsafeText(text string) {
	pos := nncp.tcp.appendFragment(text)
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     TEXT_NODE_KIND,
		Title:    TEXT_NODE_TITLE,
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
}
func (nncp *NestedNodesConfiguringProxy) Repeat(key string, t *Template) {
	pos := nncp.tcp.appendFragment(repetition{
		key:      key,
		template: t,
	})
	node := &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     REPEAT_NODE_KIND,
		Title:    key,
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
}
