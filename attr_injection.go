package go2html

import "fmt"

type AttrInjectionNode struct {
	injname string
}

const ATTR_INJECTION_NODE_TITLE_TEMPLATE = "{{%s}}"

func AttrInjection(injname string) Node {
	return &AttrInjectionNode{
		injname,
	}
}
func (n *AttrInjectionNode) Title() string {
	return fmt.Sprintf(ATTR_INJECTION_NODE_TITLE_TEMPLATE, n.injname)
}
func (n *AttrInjectionNode) WriteTo(btc *BreakthroughContext) {
	btc.WriteFragment(" ")
	btc.MarkInjection(n.injname)
	btc.Report("ok")
}
