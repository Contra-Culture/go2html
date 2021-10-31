package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
	"github.com/Contra-Culture/go2html/node"
)

type (
	NestedNodesConfiguringProxy struct {
		tcp     *TemplateConfiguringProxy
		parent  *node.Node
		context *fragments.Context
	}
)

func (nncp *NestedNodesConfiguringProxy) Elem(
	name string,
	configureSelf func(*ElemConfiguringProxy),
	configureNested func(*NestedNodesConfiguringProxy),
) {
	node := nncp.parent.AddChild(node.ELEM_NODE_KIND, []string{name})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(fmt.Sprintf("<%s", name))
		configureSelf(
			&ElemConfiguringProxy{
				tcp:     nncp.tcp,
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
				tcp:     nncp.tcp,
				parent:  node,
				context: c,
			})
		c.Append(fmt.Sprintf("</%s>", name))
	})
}
func (nncp *NestedNodesConfiguringProxy) Template(key string, t *Template) {
	if len(key) == 0 {
		key = t.key
	}
	nncp.parent.AddChild(node.TEMPLATE_NODE_KIND, []string{key})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(
			&Template{
				key:       key,
				nodes:     t.nodes,
				fragments: t.fragments,
			})
	})
}
func (nncp *NestedNodesConfiguringProxy) Comment(text string) {
	nncp.parent.AddChild(node.COMMENT_NODE_KIND, []string{})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(fmt.Sprintf("<!-- %s -->", text))
	})
}
func (nncp *NestedNodesConfiguringProxy) Doctype() {
	nncp.parent.AddChild(node.DOCTYPE_NODE_KIND, []string{})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append("<!DOCTYPE html>")
	})
}
func (nncp *NestedNodesConfiguringProxy) TextInjection(key string) {
	nncp.parent.AddChild(node.TEXT_INJECTION_NODE_KIND, []string{key})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(
			injection{
				key: key,
				modifiers: []func(string) string{
					HTMLEscape,
				},
			})
	})
}
func (nncp *NestedNodesConfiguringProxy) UnsafeTextInjection(key string) {
	nncp.parent.AddChild(node.UNSAFE_TEXT_INJECTION_NODE_KIND, []string{key})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(
			injection{
				key: key,
			})
	})
}
func (nncp *NestedNodesConfiguringProxy) Text(text string) {
	nncp.parent.AddChild(node.TEXT_NODE_KIND, []string{})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(safeTextReplacer.Replace(text))
	})
}
func (nncp *NestedNodesConfiguringProxy) UnsafeText(text string) {
	nncp.parent.AddChild(node.UNSAFE_TEXT_NODE_KIND, []string{})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(text)
	})
}
func (nncp *NestedNodesConfiguringProxy) Repeat(key string, t *Template) {
	nncp.parent.AddChild(node.REPEAT_NODE_KIND, []string{key})
	nncp.context.InContext(func(c *fragments.Context) {
		c.Append(
			repetition{
				key:      key,
				template: t,
			})
	})
}
