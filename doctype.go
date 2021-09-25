package go2html

type (
	DoctypeNode struct{}
)

const DOCTYPE_TITLE = "doctype"
const DOCTYPE = "<!DOCTYPE html>"

func Doctype() *DoctypeNode {
	return &DoctypeNode{}
}
func (n *DoctypeNode) Title() string {
	return DOCTYPE_TITLE
}
func (n *DoctypeNode) Commit(pp *PrecompilingProxy) {
	pp.AppendFragment(DOCTYPE)
	pp.Report("ok")
}
