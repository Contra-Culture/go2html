package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
)

type (
	NestedNodesConfiguringProxy struct {
		tcp    *TemplateConfiguringProxy
		parent *Node
		path   fragments.NodePath
	}
)

func (nncp *NestedNodesConfiguringProxy) Elem(
	name string,
	configureSelf func(*ElemConfiguringProxy),
	configureNested func(*NestedNodesConfiguringProxy),
) {

	path := append(nncp.path, len(nncp.parent.Children))
	node := &Node{
		Kind:                     ELEM_NODE_KIND,
		Title:                    name,
		Attributes:               map[string]string{},
		AttributeInjections:      map[string]injection{},
		AttributeValueInjections: map[string]injection{},
		Children:                 []*Node{},
	}
	nncp.parent.Children = append(nncp.parent.Children, node)
	nncp.tcp.template.fragments.Append(
		fmt.Sprintf("<%s", name),
		path,
	)
	configureSelf(&ElemConfiguringProxy{
		tcp:  nncp.tcp,
		node: node,
		path: path,
	})
	typ := elemTyp(name)
	if typ == VOID_ELEM_TYPE {
		nncp.tcp.template.fragments.Append(
			"/>",
			path,
		)
		return
	}
	nncp.tcp.template.fragments.Append(
		">",
		path,
	)
	configureNested(&NestedNodesConfiguringProxy{
		tcp:    nncp.tcp,
		parent: node,
		path:   path,
	})
	nncp.tcp.template.fragments.Append(
		fmt.Sprintf("</%s>", name),
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) Template(key string, t *Template) {
	if len(key) == 0 {
		key = t.key
	}
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  TEMPLATE_NODE_KIND,
			Title: key,
		},
	)
	nncp.tcp.template.fragments.Append(
		&Template{
			key:       key,
			nodes:     t.nodes,
			fragments: t.fragments,
		},
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) Comment(text string) {
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  COMMENT_NODE_KIND,
			Title: COMMENT_NODE_TITLE,
		},
	)
	nncp.tcp.template.fragments.Append(
		fmt.Sprintf("<!-- %s -->", text),
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) Doctype() {
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  DOCTYPE_NODE_KIND,
			Title: DOCTYPE_NODE_TITLE,
		},
	)
	nncp.tcp.template.fragments.Append(
		"<!DOCTYPE html>",
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) TextInjection(key string) {
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  TEXT_INJECTION_NODE_KIND,
			Title: key,
		},
	)
	nncp.tcp.template.fragments.Append(
		injection{
			key: key,
			modifiers: []func(string) string{
				HTMLEscape,
			},
		},
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) UnsafeTextInjection(key string) {
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  TEXT_INJECTION_NODE_KIND,
			Title: key,
		},
	)
	nncp.tcp.template.fragments.Append(
		injection{
			key: key,
		},
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) Text(text string) {
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  TEXT_NODE_KIND,
			Title: TEXT_NODE_TITLE,
		},
	)
	nncp.tcp.template.fragments.Append(
		safeTextReplacer.Replace(text),
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) UnsafeText(text string) {
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  TEXT_NODE_KIND,
			Title: TEXT_NODE_TITLE,
		},
	)
	nncp.tcp.template.fragments.Append(
		text,
		path,
	)
}
func (nncp *NestedNodesConfiguringProxy) Repeat(key string, t *Template) {
	path := append(nncp.path, len(nncp.parent.Children))
	nncp.parent.Children = append(
		nncp.parent.Children,
		&Node{
			Kind:  REPEAT_NODE_KIND,
			Title: key,
		},
	)
	nncp.tcp.template.fragments.Append(
		repetition{
			key:      key,
			template: t,
		},
		path,
	)
}
