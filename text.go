package go2html

import "strings"

type TextNode struct {
	text string
}

var safeTextReplacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "\"", "&quot", "'", "&quot")

const TEXT_NODE_TITLE = "\"text\""

func Text(text string) *TextNode {
	return &TextNode{safeTextReplacer.Replace(text)}
}
func RawText(text string) *TextNode {
	return &TextNode{text}
}
func (n *TextNode) title() string {
	return TEXT_NODE_TITLE
}
func (n *TextNode) writeTo(btc *breakthroughContext) {
	btc.writeFragment(n.text)
	btc.report("ok")
}
