package go2html

import "fmt"

type (
	Node interface {
		Title() string
		WriteTo(btc *BreakthroughContext)
	}
	WrongNode struct {
		originalNode string
		errors       []string
	}
)

const WRONG_NODE_TITLE_TEMPLATE = "WRONG(%s)"
const WRONG_NODE_ERROR_REPORT_TEMPLATE = "error: %s"

func Wrong(originalNode string, errors []string) Node {
	return &WrongNode{
		originalNode,
		errors,
	}
}
func (n *WrongNode) Title() string {
	return fmt.Sprintf(WRONG_NODE_TITLE_TEMPLATE, n.originalNode)
}
func (n *WrongNode) WriteTo(btc *BreakthroughContext) {
	for _, rr := range n.errors {
		btc.Report(fmt.Sprintf(WRONG_NODE_ERROR_REPORT_TEMPLATE, rr))
	}
	btc.WriteFragment(fmt.Sprintf("<!-- %s -->", n.Title()))
}
