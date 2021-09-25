package go2html

import "fmt"

type AttrInjectionNode struct {
	key injectionKey
}

const ATTR_INJECTION_NODE_TITLE_TEMPLATE = "{{%s}}"

func AttrInjection(key string) *AttrInjectionNode {
	return &AttrInjectionNode{
		injectionKey(key),
	}
}
func (n *AttrInjectionNode) Title() string {
	return fmt.Sprintf(ATTR_INJECTION_NODE_TITLE_TEMPLATE, n.key)
}
func (n *AttrInjectionNode) Commit(pp *PrecompilingProxy) {
	pp.AppendFragment(" ")
	pp.AppendFragment(n.key)
	pp.Report("ok")
}
