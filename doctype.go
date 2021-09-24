package go2html

type (
	DoctypeNode struct{}
)

const DOCTYPE_TITLE = "doctype"
const DOCTYPE = "<!DOCTYPE html>"

func Doctype() Node {
	return &DoctypeNode{}
}
func (n *DoctypeNode) Title() string {
	return DOCTYPE_TITLE
}
func (n *DoctypeNode) WriteTo(btc *BreakthroughContext) {
	btc.WriteFragment(DOCTYPE)
	btc.Report("ok")
}
