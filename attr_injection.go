package go2html

import "fmt"

type AttrInjectionNode struct {
	injname string
}

const ATTR_INJECTION_NODE_TITLE_TEMPLATE = "{{%s}}"

func AttrInjection(injname string) *AttrInjectionNode {
	return &AttrInjectionNode{
		injname,
	}
}
func (n *AttrInjectionNode) title() string {
	return fmt.Sprintf(ATTR_INJECTION_NODE_TITLE_TEMPLATE, n.injname)
}
func (n *AttrInjectionNode) writeTo(btc *breakthroughContext) {
	btc.writeFragment(" ")
	btc.markInjection(n.injname)
	btc.report("ok")
}
