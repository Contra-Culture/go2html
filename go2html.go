package go2html

import (
	"fmt"
	"strings"
)

type elemType int

const (
	NO_ELEM_TYPE = elemType(iota)
	VOID_ELEM_TYPE
	TEMPLATE_ELEM_TYPE
	RAW_TEXT_ELEM_TYPE
	ESCAPABLE_RAW_TEXT_ELEM_TYPE
	FOREIGN_ELEM_TYPE
	NORMAL_ELEM_TYPE
)

var elemTypeNames = map[elemType]string{
	VOID_ELEM_TYPE:               "void",
	TEMPLATE_ELEM_TYPE:           "template",
	RAW_TEXT_ELEM_TYPE:           "raw text",
	ESCAPABLE_RAW_TEXT_ELEM_TYPE: "escapable raw text",
	FOREIGN_ELEM_TYPE:            "foreign", // no support yet
	NORMAL_ELEM_TYPE:             "normal",
}
var safeTextReplacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "\"", "&quot", "'", "&quot")
var voidElements = []string{
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
var templateElements = []string{
	"template",
}
var rawTextElements = []string{
	"script",
	"style",
}
var escapableRawTextElements = []string{
	"textarea",
	"title",
}
var normalElements = []string{
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

const (
	injectionType = iota
	elementType
	textType
	commentType
)

type (
	InjectionNode struct {
		key string
	}
	ElementNode struct {
		typ      elemType
		elem     string
		props    [][2]string
		children []*Node
	}
	TextNode struct {
		text string
	}
	CommentNode struct {
		text string
	}
	Node struct {
		typ  int
		impl interface{}
	}
	Template struct {
		root         *Node
		marks        []int
		marksMapping map[string]int
		precompiled  []string
		report       *NodeReport
	}
	NodeReport struct {
		Title    string
		Messages []string
		Children []*NodeReport
	}
	breakthroughContext struct {
		template   *Template
		nodeReport *NodeReport
	}
)

func (btc *breakthroughContext) report(message string) {
	btc.nodeReport.Messages = append(btc.nodeReport.Messages, message)
}
func (btc *breakthroughContext) child(title string) *breakthroughContext {
	nr := &NodeReport{
		Title:    title,
		Messages: []string{},
		Children: []*NodeReport{},
	}
	btc.nodeReport.Children = append(btc.nodeReport.Children, nr)
	return &breakthroughContext{
		template:   btc.template,
		nodeReport: nr,
	}
}
func (btc *breakthroughContext) writeFragment(f string) {
	t := btc.template
	isLastElementReplacement := len(t.marks) > 0 && t.marks[len(t.marks)-1] == len(t.precompiled)-1
	if isLastElementReplacement || len(t.precompiled) == 0 {
		t.precompiled = append(t.precompiled, f)
		return
	}
	lastIdx := len(t.precompiled) - 1
	t.precompiled[lastIdx] = fmt.Sprintf("%s%s", t.precompiled[lastIdx], f)
}
func (btc *breakthroughContext) markInjection(key string) {
	t := btc.template
	mark := len(t.precompiled)
	if mark < 0 {
		mark = 0
	}
	t.marks = append(t.marks, mark)
	t.marksMapping[key] = mark
	t.precompiled = append(t.precompiled, fmt.Sprintf("{{ %s }}", key))
	t.marks = append(t.marks, mark)
	t.marksMapping[key] = mark
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
func Elem(name string, props [][2]string, children ...*Node) *Node { // repeat, or, optional, variant nodes
	var et = elemTyp(name)
	return &Node{
		typ: elementType,
		impl: &ElementNode{
			typ:      et,
			elem:     name,
			props:    props,
			children: children,
		},
	}
}
func Comment(text string) *Node {
	return &Node{
		typ:  commentType,
		impl: &CommentNode{text},
	}
}
func Injection(key string) *Node {
	return &Node{
		typ:  injectionType,
		impl: &InjectionNode{key},
	}
}
func Text(text string) *Node {
	return &Node{
		impl: &TextNode{safeTextReplacer.Replace(text)},
		typ:  textType,
	}
}
func RawText(text string) *Node {
	return &Node{
		impl: &TextNode{text},
		typ:  textType,
	}
}
func (node *Node) title() string {
	switch unpacked := node.impl.(type) {
	case *TextNode:
		return "\"text\""
	case *CommentNode:
		return "<!---->"
	case *InjectionNode:
		return fmt.Sprintf("{{%s}}", unpacked.key)
	case *ElementNode:
		return fmt.Sprintf("<%s>", unpacked.elem)
	}
	return "!wrong node type!"
}
func (node *Node) writeTo(btc *breakthroughContext) {
	switch unpacked := node.impl.(type) {
	case *TextNode:
		btc.writeFragment(unpacked.text)
		btc.report("ok")
	case *CommentNode:
		btc.writeFragment(fmt.Sprintf("<!-- %s -->", unpacked.text))
		btc.report("ok")
	case *InjectionNode:
		btc.markInjection(unpacked.key)
		btc.report("ok")
	case *ElementNode:
		var sb strings.Builder
		sb.WriteRune('<')
		sb.WriteString(unpacked.elem)
		for _, pair := range unpacked.props {
			sb.WriteRune(' ')
			sb.WriteString(pair[0])
			sb.WriteString("=\"")
			sb.WriteString(pair[1])
			sb.WriteRune('"')
		}
		if unpacked.typ == VOID_ELEM_TYPE {
			sb.WriteString("/>")
			btc.writeFragment(sb.String())
			if len(unpacked.children) != 0 {
				btc.report("error: void element can't have children (children ignored)")
				return
			}
			btc.report("ok: self-closing")
			return
		}
		sb.WriteRune('>')
		btc.writeFragment(sb.String())
		btc.report("ok: opening")
		for _, elem := range unpacked.children {
			elem.writeTo(btc.child(elem.title()))
		}
		sb.Reset()
		sb.WriteString("</")
		sb.WriteString(unpacked.elem)
		sb.WriteRune('>')
		btc.writeFragment(sb.String())
		btc.report("ok: closing")
	default:
		btc.report("error: wrong node type")
	}
}
func (node *Node) Template() *Template {
	nodeReport := &NodeReport{
		Title:    node.title(),
		Messages: []string{},
		Children: []*NodeReport{},
	}
	template := &Template{
		root:         node,
		marks:        []int{},
		marksMapping: map[string]int{},
		precompiled:  []string{},
		report:       nodeReport,
	}
	btc := &breakthroughContext{
		template,
		nodeReport,
	}
	node.writeTo(btc)
	return template
}
func (t *Template) CompileWith(replacements map[string]interface{}) string {
	fragments := append([]string{}, t.precompiled...)
	for key, idx := range t.marksMapping {
		unknownRepl, ok := replacements[key]
		if !ok {
			panic(fmt.Sprintf("replacement for \"%s\" key is not provied", key))
		}
		switch repl := unknownRepl.(type) {
		case string:
			fragments[idx] = repl
		case *Template:
			fragments[idx] = repl.CompileWith(replacements)
		default:
			panic("wrong replacement type")
		}
	}
	return strings.Join(fragments, "")
}
func (t *Template) Report() *NodeReport {
	return t.report
}
