package go2html

import "fmt"

type InjectionNode struct {
	key string
}

const INJECTION_NODE_TITLE_TEMPLATE = "{{%s}}"

func Injection(key string) *InjectionNode {
	return &InjectionNode{key}
}
func (n *InjectionNode) Title() string {
	return fmt.Sprintf(INJECTION_NODE_TITLE_TEMPLATE, n.key)
}
func (n *InjectionNode) WriteTo(btc *BreakthroughContext) {
	btc.MarkInjection(n.key)
	btc.Report("ok")
}
