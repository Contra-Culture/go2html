package go2html

import "fmt"

type (
	Node interface {
		Title() string
		Commit(pp *PrecompilingProxy)
	}
	PopulatingNode interface {
		Populate(replacements interface{}) string
		Scope() (string, bool)
	}
	WrongNode struct {
		originalNode string
		errors       []string
	}
	injectionKey string
)

const WRONG_NODE_TITLE_TEMPLATE = "WRONG(%s)"
const WRONG_NODE_ERROR_REPORT_TEMPLATE = "error: %s"

func Wrong(originalNode string, errors []string) *WrongNode {
	return &WrongNode{
		originalNode,
		errors,
	}
}
func (n *WrongNode) Title() string {
	return fmt.Sprintf(WRONG_NODE_TITLE_TEMPLATE, n.originalNode)
}
func (n *WrongNode) Commit(pp *PrecompilingProxy) {
	for _, rr := range n.errors {
		pp.Report(fmt.Sprintf(WRONG_NODE_ERROR_REPORT_TEMPLATE, rr))
	}
	pp.AppendFragment(fmt.Sprintf("<!-- %s -->", n.Title()))
}
