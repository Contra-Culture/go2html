package go2html

import (
	"fmt"
)

type (
	ElementNode struct {
		typ      elemType
		elem     string
		attrs    []Node
		children []Node
	}
	elemType int
)

const ELEMENT_NODE_TITLE_TEMPLATE = "<%s>"

var (
	elemTypeNames = map[elemType]string{
		VOID_ELEM_TYPE:               "void",
		TEMPLATE_ELEM_TYPE:           "template",
		RAW_TEXT_ELEM_TYPE:           "raw text",
		ESCAPABLE_RAW_TEXT_ELEM_TYPE: "escapable raw text",
		FOREIGN_ELEM_TYPE:            "foreign", // no support yet
		NORMAL_ELEM_TYPE:             "normal",
	}
	voidElements = []string{
		"area",
		"base",
		"br",
		"col",
		"embed",
		"hr",
		"img",
		"input",
		"link",
		"meta",
		"param",
		"source",
		"track",
		"wbr",
	}
	templateElements = []string{
		"template",
	}
	rawTextElements = []string{
		"script",
		"style",
	}
	escapableRawTextElements = []string{
		"textarea",
		"title",
	}
	normalElements = []string{
		"html",
		"base",
		"head",
		"link",
		"meta",
		"style",
		"title",
		"body",
		"address",
		"article",
		"aside",
		"footer",
		"header",
		"h1",
		"h2",
		"h3",
		"h4",
		"h5",
		"h6",
		"main",
		"nav",
		"section",
		"blockquote",
		"dd",
		"div",
		"dl",
		"dt",
		"figcaption",
		"figure",
		"hr",
		"li",
		"ol",
		"p",
		"pre",
		"ul",
		"a",
		"abbr",
		"b",
		"bdi",
		"bdo",
		"br",
		"cite",
		"code",
		"data",
		"dfn",
		"em",
		"i",
		"kbd",
		"mark",
		"q",
		"rp",
		"rt",
		"ruby",
		"s",
		"samp",
		"small",
		"span",
		"strong",
		"sub",
		"sup",
		"time",
		"u",
		"var",
		"wbr",
		"area",
		"audio",
		"img",
		"map",
		"track",
		"video",
		"embed",
		"iframe",
		"object",
		"param",
		"picture",
		"portal",
		"source",
		"svg",
		"math",
		"canvas",
		"noscript",
		"script",
		"del",
		"ins",
		"caption",
		"col",
		"colgroup",
		"table",
		"tbody",
		"td",
		"tfoot",
		"th",
		"thead",
		"tr",
		"button",
		"datalist",
		"fieldset",
		"form",
		"input",
		"label",
		"legand",
		"meter",
		"optgroup",
		"option",
		"output",
		"progress",
		"select",
		"textarea",
		"details",
		"dialog",
		"menu",
		"summary",
		"slot",
		"template",
		// obsolete/deprecated
		"acronym",
		"applet",
		"baseform",
		"bgsound",
		"big",
		"blink",
		"center",
		"content",
		"dir",
		"font",
		"frame",
		"frameset",
		"hgroup",
		"image",
		"keygen",
		"marquee",
		"menuitem",
		"nobr",
		"noembed",
		"noframes",
		"plaintext",
		"rb",
		"rtc",
		"shadow",
		"spacer",
		"strike",
		"tt",
		"xmp",
	}
)

func Elem(name string, attrs []Node, children []Node) *ElementNode {
	return &ElementNode{
		typ:      elemTyp(name),
		elem:     name,
		attrs:    attrs,
		children: children,
	}
}
func elemTyp(name string) elemType {
	for _, ve := range voidElements {
		if name == ve {
			return VOID_ELEM_TYPE
		}
	}
	for _, te := range templateElements {
		if name == te {
			return TEMPLATE_ELEM_TYPE
		}
	}
	for _, rte := range rawTextElements {
		if name == rte {
			return RAW_TEXT_ELEM_TYPE
		}
	}
	for _, erte := range escapableRawTextElements {
		if name == erte {
			return ESCAPABLE_RAW_TEXT_ELEM_TYPE
		}
	}
	for _, ne := range normalElements {
		if name == ne {
			return NORMAL_ELEM_TYPE
		}
	}
	return NO_ELEM_TYPE
}
func (n *ElementNode) title() string {
	return fmt.Sprintf(ELEMENT_NODE_TITLE_TEMPLATE, n.elem)
}
func (n *ElementNode) writeTo(btc *breakthroughContext) {
	btc.writeFragment(fmt.Sprintf("<%s", n.elem))
	if len(n.attrs) > 0 {
		attrbtc := btc.child("attrs")
		for _, attr := range n.attrs {
			attr.writeTo(attrbtc)
		}
	}
	if n.typ == VOID_ELEM_TYPE {
		btc.writeFragment("/>")
		if len(n.children) != 0 {
			btc.report("error: void element can't have children (children ignored)")
			return
		}
		btc.report("ok: self-closing")
		return
	}
	btc.writeFragment(">")
	btc.report("ok: opening")
	for _, child := range n.children {
		child.writeTo(btc.child(child.title()))
	}
	btc.writeFragment(fmt.Sprintf("</%s>", n.elem))
	btc.report("ok: closing")
}
