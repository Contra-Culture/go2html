package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
	"github.com/Contra-Culture/go2html/node"
)

type (
	ElemCfgr struct {
		tcp     *TemplateCfgr
		node    *node.Node
		context *fragments.Context
	}
)

func (ecp *ElemCfgr) AttrInjection(key string) {
	fragment := injection{
		key: key,
	}
	ecp.context.InContext(
		func(c *fragments.Context) {
			c.Append(" ")
			c.Append(fragment)
		})
	ecp.node.AddChild(node.ATTRIBUTE_INJECTION_NODE_KIND, []string{key})
}
func (ecp *ElemCfgr) AttrValueInjection(name string, key string) {
	fragment := injection{
		key: key,
	}
	ecp.context.InContext(
		func(c *fragments.Context) {
			c.Append(fmt.Sprintf(" %s=\"", name))
			c.Append(fragment)
			c.Append("\"")
		})
	ecp.node.AddChild(node.ATTRIBUTE_VALUE_INJECTION_NODE_KIND, []string{name, key})
}
func (ecp *ElemCfgr) Attr(name string, value string) {
	ecp.context.InContext(
		func(c *fragments.Context) {
			c.Append(fmt.Sprintf(" %s=\"%s\"", name, value))
		})
	ecp.node.AddChild(node.ATTRIBUTE_NODE_KIND, []string{name})
}
