package node

import (
	"fmt"
	"strings"
)

type (
	Kind int
	Node struct {
		kind     Kind
		title    string
		children []*Node
	}
)

const (
	_ Kind = iota
	ELEM_NODE_KIND
	ATTRIBUTE_NODE_KIND
	ATTRIBUTES_INJECTION_NODE_KIND
	ATTRIBUTE_VALUE_INJECTION_NODE_KIND
	TEMPLATE_NODE_KIND
	TEMPLATE_INJECTION_NODE_KIND
	COMMENT_NODE_KIND
	DOCTYPE_NODE_KIND
	TEXT_INJECTION_NODE_KIND
	UNSAFE_TEXT_INJECTION_NODE_KIND
	TEXT_NODE_KIND
	UNSAFE_TEXT_NODE_KIND
	REPEAT_NODE_KIND
	VARIANTS_NODE_KIND
)

func New(k Kind, tinj []string) *Node {
	var t string
	switch k {
	case ELEM_NODE_KIND:
		t = fmt.Sprintf("<%s>", tinj[0])
	case ATTRIBUTE_NODE_KIND:
		t = fmt.Sprintf("%s=", tinj[0])
	case ATTRIBUTES_INJECTION_NODE_KIND:
		t = fmt.Sprintf("?%s=", tinj[0])
	case ATTRIBUTE_VALUE_INJECTION_NODE_KIND:
		t = fmt.Sprintf("%s=?%s", tinj[0], tinj[1])
	case TEMPLATE_NODE_KIND:
		t = fmt.Sprintf("template:%s", tinj[0])
	case TEMPLATE_INJECTION_NODE_KIND:
		t = fmt.Sprintf("template:?%s", tinj[0])
	case COMMENT_NODE_KIND:
		t = "<!-->"
	case DOCTYPE_NODE_KIND:
		t = "!DOCTYPE"
	case TEXT_INJECTION_NODE_KIND:
		t = fmt.Sprintf("?\"%s\"", tinj[0])
	case UNSAFE_TEXT_INJECTION_NODE_KIND:
		t = fmt.Sprintf("?\"%s\"!", tinj[0])
	case TEXT_NODE_KIND:
		t = "\"...\""
	case UNSAFE_TEXT_NODE_KIND:
		t = "\"...\"!"
	case REPEAT_NODE_KIND:
		t = fmt.Sprintf("repeat(template:%s)", tinj[0])
	case VARIANTS_NODE_KIND:
		t = fmt.Sprintf("variants(%s)", strings.Join(tinj, ", "))
	default:
		panic(fmt.Sprintf("wrong element kind `%d`", k)) // can't occur
	}
	return &Node{
		kind:  k,
		title: t,
	}
}

func (n *Node) AddChild(k Kind, tinj []string) *Node {
	node := New(k, tinj)
	n.children = append(
		n.children,
		node,
	)
	return node
}

func String(k Kind) string {
	switch k {
	case ELEM_NODE_KIND:
		return "element"
	case ATTRIBUTE_NODE_KIND:
		return "attribute"
	case ATTRIBUTES_INJECTION_NODE_KIND:
		return "attribute-injection"
	case ATTRIBUTE_VALUE_INJECTION_NODE_KIND:
		return "attribute-value-injection"
	case TEMPLATE_NODE_KIND:
		return "template"
	case TEMPLATE_INJECTION_NODE_KIND:
		return "template-injection"
	case COMMENT_NODE_KIND:
		return "comment"
	case DOCTYPE_NODE_KIND:
		return "doctype"
	case TEXT_INJECTION_NODE_KIND:
		return "text-injection"
	case UNSAFE_TEXT_INJECTION_NODE_KIND:
		return "unsafe-text-injection"
	case TEXT_NODE_KIND:
		return "text"
	case UNSAFE_TEXT_NODE_KIND:
		return "unsafe-text"
	case REPEAT_NODE_KIND:
		return "repeat"
	case VARIANTS_NODE_KIND:
		return "variants"
	default:
		panic(fmt.Sprintf("wrong element kind `%d`", k))
	}
}
