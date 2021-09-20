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
func (n *AttrValueInjectionNode) Title() string {
	return fmt.Sprintf(ATTR_VALUE_INJECTION_NODE_TITLE_TEMPLATE, n.injname)
}
func (n *AttrValueInjectionNode) WriteTo(btc *BreakthroughContext) {
	btc.WriteFragment(fmt.Sprintf(" %s=\"", n.name))
	btc.MarkInjection(n.injname)
	btc.WriteFragment("\"")
	btc.Report("ok")
}
