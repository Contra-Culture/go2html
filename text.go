package go2html

import "strings"

type TextNode struct {
	title string
	text  string
}

var safeTextReplacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "\"", "&quot", "'", "&quot")

const TEXT_NODE_TITLE = "\"text\""
const RAW_TEXT_NODE_TITLE = "!\"text\""

func Text(text string) *TextNode {
	return &TextNode{
		title: TEXT_NODE_TITLE,
		text:  safeTextReplacer.Replace(text),
	}
}
func RawText(text string) *TextNode {
	return &TextNode{
		title: RAW_TEXT_NODE_TITLE,
		text:  text,
	}
}
func (n *TextNode) Title() string {
	return n.title
}
func (n *TextNode) WriteTo(btc *BreakthroughContext) {
	btc.WriteFragment(n.text)
	btc.Report("ok")
}
