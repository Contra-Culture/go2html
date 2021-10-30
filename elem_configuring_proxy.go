package go2html

import "fmt"

type (
	ElemConfiguringProxy struct {
		tcp  *TemplateConfiguringProxy
		node *Node
	}
)

func (ecp *ElemConfiguringProxy) AttrInjection(key string) {
	posBegin := ecp.tcp.appendFragment(" ")
	fragment := injection{
		key: key,
	}
	posEnd := ecp.tcp.appendFragment(fragment)
	ecp.node.Children = append(
		ecp.node.Children,
		&Node{
			PosBegin: posBegin,
			PosEnd:   posEnd,
			Kind:     ATTRIBUTE_INJECTION_NODE_KIND,
			Title:    key,
		},
	)
}
func (ecp *ElemConfiguringProxy) AttrValueInjection(name string, key string) {
	posBegin := ecp.tcp.appendFragment(fmt.Sprintf(" %s=\"", name))
	fragment := injection{
		key: key,
	}
	ecp.tcp.appendFragment(fragment)
	posEnd := ecp.tcp.appendFragment("\"")
	ecp.node.Children = append(
		ecp.node.Children,
		&Node{
			PosBegin: posBegin,
			PosEnd:   posEnd,
			Kind:     ATTRIBUTE_VALUE_INJECTION_NODE_KIND,
			Title:    key,
		},
	)
}
func (ecp *ElemConfiguringProxy) Attr(name string, value string) {
	pos := ecp.tcp.appendFragment(fmt.Sprintf(" %s=\"%s\"", name, value))
	ecp.node.Children = append(ecp.node.Children, &Node{
		PosBegin: pos,
		PosEnd:   pos,
		Kind:     ATTRIBUTE_NODE_KIND,
		Title:    name,
	})
}
