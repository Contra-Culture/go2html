package go2html

import (
	"fmt"

	"github.com/Contra-Culture/go2html/fragments"
)

type (
	ElemConfiguringProxy struct {
		tcp  *TemplateConfiguringProxy
		node *Node
		path fragments.NodePath
	}
)

func (ecp *ElemConfiguringProxy) AttrInjection(key string) {
	ecp.tcp.template.fragments.Append(
		" ",
		ecp.path,
	)
	fragment := injection{
		key: key,
	}
	ecp.tcp.template.fragments.Append(
		fragment,
		ecp.path,
	)
	ecp.node.AttributeInjections[key] = fragment
}
func (ecp *ElemConfiguringProxy) AttrValueInjection(name string, key string) {
	ecp.tcp.template.fragments.Append(
		fmt.Sprintf(" %s=\"", name),
		ecp.path,
	)
	fragment := injection{
		key: key,
	}
	ecp.tcp.template.fragments.Append(
		fragment,
		ecp.path,
	)
	ecp.tcp.template.fragments.Append(
		"\"",
		ecp.path,
	)
	ecp.node.AttributeValueInjections[key] = fragment
}
func (ecp *ElemConfiguringProxy) Attr(name string, value string) {
	ecp.tcp.template.fragments.Append(
		fmt.Sprintf(" %s=\"%s\"", name, value),
		ecp.path,
	)
	ecp.node.Attributes[name] = value
}
