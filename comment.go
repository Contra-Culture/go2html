package go2html

import "fmt"

type CommentNode struct {
	text string
}

const COMMENT_NODE_TITLE = "<!---->"

func Comment(text string) *CommentNode {
	return &CommentNode{text}
}
func (n *CommentNode) title() string {
	return COMMENT_NODE_TITLE
}
func (n *CommentNode) writeTo(btc *breakthroughContext) {
	btc.writeFragment(fmt.Sprintf("<!-- %s -->", n.text))
	btc.report("ok")
}
