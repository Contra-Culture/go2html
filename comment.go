package go2html

import "fmt"

type CommentNode struct {
	text string
}

const COMMENT_NODE_TITLE = "<!---->"

func Comment(text string) *CommentNode {
	return &CommentNode{text}
}
func (n *CommentNode) Title() string {
	return COMMENT_NODE_TITLE
}
func (n *CommentNode) WriteTo(btc *BreakthroughContext) {
	btc.WriteFragment(fmt.Sprintf("<!-- %s -->", n.text))
	btc.Report("ok")
}
