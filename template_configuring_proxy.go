package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
)

type (
	TemplateConfiguringProxy struct {
		template *Template
	}
)

func (tcp *TemplateConfiguringProxy) Elem(
	name string,
	configureSelf func(*ElemConfiguringProxy),
	configureNested func(*NestedNodesConfiguringProxy),
) {
	node := &Node{
		Kind:                     ELEM_NODE_KIND,
		Title:                    name,
		Attributes:               map[string]string{},
		AttributeInjections:      map[string]injection{},
		AttributeValueInjections: map[string]injection{},
		Children:                 []*Node{},
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(fmt.Sprintf("<%s", name), path)

	configureSelf(&ElemConfiguringProxy{
		tcp:  tcp,
		node: node,
		path: path,
	})
	typ := elemTyp(name)
	if typ == VOID_ELEM_TYPE {
		tcp.template.fragments.Append("/>", path)
		return
	}
	tcp.template.fragments.Append(">", path)
	configureNested(&NestedNodesConfiguringProxy{
		tcp:    tcp,
		parent: node,
		path:   path,
	})
	tcp.template.fragments.Append(fmt.Sprintf("</%s>", name), path)
}
func (tcp *TemplateConfiguringProxy) Template(key string, t *Template) {
	if len(key) == 0 {
		key = t.key
	}
	node := &Node{
		Kind:  TEMPLATE_NODE_KIND,
		Title: key,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		&Template{
			key:       key,
			nodes:     t.nodes,
			fragments: t.fragments,
		},
		path,
	)
}
func (tcp *TemplateConfiguringProxy) Comment(text string) {
	node := &Node{
		Kind:  COMMENT_NODE_KIND,
		Title: COMMENT_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		fmt.Sprintf("<!-- %s -->", text),
		path,
	)
}
func (tcp *TemplateConfiguringProxy) Doctype() {
	node := &Node{
		Kind:  DOCTYPE_NODE_KIND,
		Title: DOCTYPE_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		"<!DOCTYPE html>",
		path,
	)
}
func (tcp *TemplateConfiguringProxy) TextInjection(key string) {
	node := &Node{
		Kind:  TEXT_INJECTION_NODE_KIND,
		Title: key,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		injection{
			key: key,
			modifiers: []func(string) string{
				HTMLEscape,
			},
		},
		path,
	)
}
func (tcp *TemplateConfiguringProxy) UnsafeTextInjection(key string) {
	node := &Node{
		Kind:  TEXT_INJECTION_NODE_KIND,
		Title: key,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		injection{
			key: key,
		},
		path,
	)
}
func (tcp *TemplateConfiguringProxy) Text(text string) {
	node := &Node{
		Kind:  TEXT_NODE_KIND,
		Title: TEXT_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		safeTextReplacer.Replace(text),
		path,
	)
}
func (tcp *TemplateConfiguringProxy) UnsafeText(text string) {
	node := &Node{
		Kind:  TEXT_NODE_KIND,
		Title: TEXT_NODE_TITLE,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		text,
		path,
	)
}
func (tcp *TemplateConfiguringProxy) Repeat(key string, t *Template) {
	node := &Node{
		Kind:  REPEAT_NODE_KIND,
		Title: key,
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	path := fragments.NodePath([]int{len(tcp.template.nodes) - 1})
	tcp.template.fragments.Append(
		repetition{
			key:      key,
			template: t,
		},
		path,
	)
}
