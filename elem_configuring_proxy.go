package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
)

type (
	ElemConfiguringProxy struct {
		tcp     *TemplateConfiguringProxy
		node    *Node
		context *fragments.Context
	}
)

func (ecp *ElemConfiguringProxy) AttrInjection(key string) {
	fragment := injection{
		key: key,
	}
	ecp.context.InContext(
		func(c *fragments.Context) {
			c.Append(" ")
			c.Append(fragment)
		})
	ecp.node.Children = append(
		ecp.node.Children,
		&Node{
			Kind:  ATTRIBUTE_INJECTION_NODE_KIND,
			Title: key,
		},
	)
}
func (ecp *ElemConfiguringProxy) AttrValueInjection(name string, key string) {
	fragment := injection{
		key: key,
	}
	ecp.context.InContext(
		func(c *fragments.Context) {
			c.Append(fmt.Sprintf(" %s=\"", name))
			c.Append(fragment)
			c.Append("\"")
		})
	ecp.node.Children = append(
		ecp.node.Children,
		&Node{
			Kind:  ATTRIBUTE_VALUE_INJECTION_NODE_KIND,
			Title: key,
		},
	)
}
func (ecp *ElemConfiguringProxy) Attr(name string, value string) {
	ecp.context.InContext(
		func(c *fragments.Context) {
			c.Append(fmt.Sprintf(" %s=\"%s\"", name, value))
		})
	ecp.node.Children = append(
		ecp.node.Children,
		&Node{
			Kind:  ATTRIBUTE_NODE_KIND,
			Title: name,
		})
}
