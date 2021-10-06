package go2html

import "fmt"

type (
	ElemConfiguringProxy struct {
		tcp  *TemplateConfiguringProxy
		node *Node
	}
)

func (ecp *ElemConfiguringProxy) AttrInjection(key string) {
	ecp.tcp.appendFragment(" ")
	fragment := injection{
		key: key,
	}
	ecp.tcp.appendFragment(fragment)
	ecp.node.AttributeInjections[key] = fragment
}
func (ecp *ElemConfiguringProxy) AttrValueInjection(name string, key string) {
	ecp.tcp.appendFragment(fmt.Sprintf(" %s=\"", name))
	fragment := injection{
		key: key,
	}
	ecp.tcp.appendFragment(fragment)
	ecp.tcp.appendFragment("\"")
	ecp.node.AttributeValueInjections[key] = fragment
}
func (ecp *ElemConfiguringProxy) Attr(name string, value string) {
	ecp.tcp.appendFragment(fmt.Sprintf(" %s=\"%s\"", name, value))
	ecp.node.Attributes[name] = value
}
