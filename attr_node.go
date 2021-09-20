package go2html

import "fmt"

type AttrNode struct {
	name  string
	value string
}

const ATTR_NODE_TITLE = "attr="

func Attr(name string, value string) *AttrNode {
	return &AttrNode{
		name,
		value,
	}
}
func (n *AttrNode) title() string {
	return ATTR_NODE_TITLE
}
func (n *AttrNode) writeTo(btc *breakthroughContext) {
	btc.writeFragment(fmt.Sprintf(" %s=\"%s\"", n.name, n.value))
	btc.report("ok")
}
