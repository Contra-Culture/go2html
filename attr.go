package go2html

import "fmt"

type AttrNode struct {
	name  string
	value string
}

const ATTR_NODE_TITLE = "attr="

func Attr(name string, value string) Node {
	return &AttrNode{
		name,
		value,
	}
}
func (n *AttrNode) Title() string {
	return ATTR_NODE_TITLE
}
func (n *AttrNode) WriteTo(btc *BreakthroughContext) {
	btc.WriteFragment(fmt.Sprintf(" %s=\"%s\"", n.name, n.value))
	btc.Report("ok")
}
