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
func (n *CommentNode) Commit(pp *PrecompilingProxy) {
	pp.AppendFragment(fmt.Sprintf("<!-- %s -->", n.text))
	pp.Report("ok")
}
