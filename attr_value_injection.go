package go2html

import "fmt"

type AttrValueInjectionNode struct {
	name string
	key  injectionKey
}

const ATTR_VALUE_INJECTION_NODE_TITLE_TEMPLATE = "attr={{%s}}"

func AttrValueInjection(name string, key string) *AttrValueInjectionNode {
	return &AttrValueInjectionNode{
		name,
		injectionKey(key),
	}
}
func (n *AttrValueInjectionNode) Title() string {
	return fmt.Sprintf(ATTR_VALUE_INJECTION_NODE_TITLE_TEMPLATE, n.key)
}
func (n *AttrValueInjectionNode) Commit(pp *PrecompilingProxy) {
	pp.AppendFragment(fmt.Sprintf(" %s=\"", n.name))
	pp.AppendFragment(n.key)
	pp.AppendFragment("\"")
	pp.Report("ok")
}
