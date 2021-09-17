package go2html

import (
	"fmt"
	"strings"
)

var safeTextReplacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "\"", "&quot", "'", "&quot")

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
		root        *Node
		markMapping map[string]int
		precompiled []string
	}
	fallthroughContext struct {
		precompiled    []string
		markMapping    map[string]int
		lastMappingIdx int
	}
)

func (fc *fallthroughContext) writeFragment(f string) {
	isLastElementReplacement := fc.lastMappingIdx == len(fc.precompiled)-1
	if isLastElementReplacement || len(fc.precompiled) == 0 {
		fc.precompiled = append(fc.precompiled, f)
		return
	}
	lastIdx := len(fc.precompiled) - 1
	fc.precompiled[lastIdx] = fmt.Sprintf("%s\n%s", fc.precompiled[lastIdx], f)
}
func (fc *fallthroughContext) markInjection(key string) {
	idx := len(fc.precompiled)
	if idx < 0 {
		idx = 0
	}
	fc.markMapping[key] = idx
	fc.precompiled = append(fc.precompiled, fmt.Sprintf("{{ %s }}", key))
	fc.lastMappingIdx = idx
}

func Elem(name string, props [][2]string, children ...*Node) *Node { // repeat, or, optional, variant nodes
	return &Node{
		typ: elementType,
		impl: &ElementNode{
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
func (node *Node) writeTo(fc *fallthroughContext) {
	switch unpacked := node.impl.(type) {
	case *TextNode:
		fc.writeFragment(unpacked.text)
	case *CommentNode:
		fc.writeFragment(fmt.Sprintf("<!-- %s  -->", unpacked.text))
	case *InjectionNode:
		fc.markInjection(unpacked.key)
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
		sb.WriteRune('>')
		fc.writeFragment(sb.String())
		for _, elem := range unpacked.children {
			elem.writeTo(fc)
		}
		sb.Reset()
		sb.WriteString("</")
		sb.WriteString(unpacked.elem)
		sb.WriteRune('>')
		fc.writeFragment(sb.String())
	default:
		panic("wrong node type")
	}
}
func (node *Node) Template() (template *Template) {
	fc := &fallthroughContext{
		markMapping: map[string]int{},
		precompiled: []string{},
	}
	node.writeTo(fc)
	return &Template{
		root:        node,
		markMapping: fc.markMapping,
		precompiled: fc.precompiled,
	}
}
func (t *Template) CompileWith(replacements map[string]interface{}) string {
	fragments := []string{}
	for _, f := range t.precompiled {
		fragments = append(fragments, f)
	}
	for key, idx := range t.markMapping {
		unknownRepl, ok := replacements[key]
		if !ok {
			panic(fmt.Sprintf("replacement for `%s` key is not provied", key))
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
