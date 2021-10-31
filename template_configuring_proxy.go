package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
)

type (
	TemplateConfiguringProxy struct {
		template      *Template
		parentContext *fragments.Context
	}
)

func (tcp *TemplateConfiguringProxy) Elem(
	name string,
	configureSelf func(*ElemConfiguringProxy),
	configureNested func(*NestedNodesConfiguringProxy),
) {
	node := &Node{
		Kind:     ELEM_NODE_KIND,
		Title:    name,
		Children: []*Node{},
	}
	tcp.template.nodes = append(tcp.template.nodes, node)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(fmt.Sprintf("<%s", name))
		configureSelf(
			&ElemConfiguringProxy{
				tcp:     tcp,
				node:    node,
				context: c,
			})
		typ := elemTyp(name)
		if typ == VOID_ELEM_TYPE {
			c.Append("/>")
			return
		}
		c.Append(">")
		configureNested(
			&NestedNodesConfiguringProxy{
				tcp:     tcp,
				parent:  node,
				context: c,
			})
		c.Append(fmt.Sprintf("</%s>", name))
	})
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
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			&Template{
				key:       key,
				nodes:     t.nodes,
				fragments: t.fragments,
			})
	})
}
func (tcp *TemplateConfiguringProxy) Comment(text string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		&Node{
			Kind:  COMMENT_NODE_KIND,
			Title: COMMENT_NODE_TITLE,
		})
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(fmt.Sprintf("<!-- %s -->", text))
	})
}
func (tcp *TemplateConfiguringProxy) Doctype() {
	tcp.template.nodes = append(
		tcp.template.nodes,
		&Node{
			Kind:  DOCTYPE_NODE_KIND,
			Title: DOCTYPE_NODE_TITLE,
		})
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append("<!DOCTYPE html>")
	})
}
func (tcp *TemplateConfiguringProxy) TextInjection(key string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		&Node{
			Kind:  TEXT_INJECTION_NODE_KIND,
			Title: key,
		})
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			injection{
				key: key,
				modifiers: []func(string) string{
					HTMLEscape,
				},
			})
	})
}
func (tcp *TemplateConfiguringProxy) UnsafeTextInjection(key string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		&Node{
			Kind:  TEXT_INJECTION_NODE_KIND,
			Title: key,
		})
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			injection{
				key: key,
			})
	})
}
func (tcp *TemplateConfiguringProxy) Text(text string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		&Node{
			Kind:  TEXT_NODE_KIND,
			Title: TEXT_NODE_TITLE,
		})
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(safeTextReplacer.Replace(text))
	})
}
func (tcp *TemplateConfiguringProxy) UnsafeText(text string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		&Node{
			Kind:  TEXT_NODE_KIND,
			Title: TEXT_NODE_TITLE,
		})
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(text)
	})
}
func (tcp *TemplateConfiguringProxy) Repeat(key string, t *Template) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		&Node{
			Kind:  REPEAT_NODE_KIND,
			Title: key,
		})
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			repetition{
				key:      key,
				template: t,
			})
	})
}
