package go2html

import "fmt"

type InjectionNode struct {
	key injectionKey
}

const INJECTION_NODE_TITLE_TEMPLATE = "{{%s}}"

func Injection(key string) *InjectionNode {
	return &InjectionNode{
		key: injectionKey(key),
	}
}
func (n *InjectionNode) Title() string {
	return fmt.Sprintf(INJECTION_NODE_TITLE_TEMPLATE, n.key)
}
func (n *InjectionNode) Commit(pp *PrecompilingProxy) {
	pp.AppendFragment(n.key)
	pp.Report("ok")
}
