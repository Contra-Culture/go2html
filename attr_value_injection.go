package go2html

import "fmt"

type AttrValueInjectionNode struct {
	name    string
	injname string
}

const ATTR_VALUE_INJECTION_NODE_TITLE_TEMPLATE = "attr={{%s}}"

func AttrValueInjection(name string, injname string) *AttrValueInjectionNode {
	return &AttrValueInjectionNode{
		name,
		injname,
	}
}
func (n *AttrValueInjectionNode) title() string {
	return fmt.Sprintf(ATTR_VALUE_INJECTION_NODE_TITLE_TEMPLATE, n.injname)
}
func (n *AttrValueInjectionNode) writeTo(btc *breakthroughContext) {
	btc.writeFragment(fmt.Sprintf(" %s=\"", n.name))
	btc.markInjection(n.injname)
	btc.writeFragment("\"")
	btc.report("ok")
}
