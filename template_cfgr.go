package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
	"github.com/Contra-Culture/go2html/node"
)

type (
	TemplateCfgr struct {
		template      *Template
		parentContext *fragments.Context
	}
)

func (tcp *TemplateCfgr) Elem(
	name string,
	configureSelf func(*ElemCfgr),
	configureNested func(*NestedNodesCfgr),
) {
	node := node.New(node.ELEM_NODE_KIND, []string{name})
	tcp.template.nodes = append(tcp.template.nodes, node)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(fmt.Sprintf("<%s", name))
		configureSelf(
			&ElemCfgr{
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
			&NestedNodesCfgr{
				tcp:     tcp,
				parent:  node,
				context: c,
			})
		c.Append(fmt.Sprintf("</%s>", name))
	})
}
func (tcp *TemplateCfgr) Template(key string, t *Template) {
	if len(key) == 0 {
		key = t.key
	}
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.TEMPLATE_NODE_KIND, []string{key}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			&Template{
				key:       key,
				nodes:     t.nodes,
				fragments: t.fragments,
			})
	})
}
func (tcp *TemplateCfgr) TemplateInjection(key string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.TEMPLATE_INJECTION_NODE_KIND, []string{key}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			templateInjection{
				key: key,
			})
	})
}
func (tcp *TemplateCfgr) Comment(text string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.COMMENT_NODE_KIND, []string{}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(fmt.Sprintf("<!-- %s -->", text))
	})
}
func (tcp *TemplateCfgr) Doctype() {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.DOCTYPE_NODE_KIND, []string{}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append("<!DOCTYPE html>")
	})
}
func (tcp *TemplateCfgr) TextInjection(key string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.TEXT_INJECTION_NODE_KIND, []string{key}),
	)
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
func (tcp *TemplateCfgr) UnsafeTextInjection(key string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.UNSAFE_TEXT_INJECTION_NODE_KIND, []string{key}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			injection{
				key: key,
			})
	})
}
func (tcp *TemplateCfgr) Text(text string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.TEXT_NODE_KIND, []string{}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(safeTextReplacer.Replace(text))
	})
}
func (tcp *TemplateCfgr) UnsafeText(text string) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.UNSAFE_TEXT_NODE_KIND, []string{}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(text)
	})
}
func (tcp *TemplateCfgr) Repeat(key string, t *Template) {
	tcp.template.nodes = append(
		tcp.template.nodes,
		node.New(node.REPEAT_NODE_KIND, []string{key}),
	)
	tcp.template.fragments.InContext(func(c *fragments.Context) {
		c.Append(
			repetition{
				key:      key,
				template: t,
			})
	})
}
