package go2html

import "fmt"

type InjectionNode struct {
	key string
}

const INJECTION_NODE_TITLE_TEMPLATE = "{{%s}}"

func Injection(key string) *InjectionNode {
	return &InjectionNode{key}
}
func (n *InjectionNode) title() string {
	return fmt.Sprintf(INJECTION_NODE_TITLE_TEMPLATE, n.key)
}
func (n *InjectionNode) writeTo(btc *breakthroughContext) {
	btc.markInjection(n.key)
	btc.report("ok")
}
