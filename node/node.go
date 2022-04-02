package node

import (
	"fmt"
	"strings"
)

type (
	Kind int
	Node struct {
		kind             Kind
		title            string
		isClassNamespece bool
		classNamespace   string
		children         []*Node
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
const (
	ELEMENT_KIND_STRING                   = "element"
	ATTRIBUTE_KIND_STRING                 = "attribute"
	ATTRIBUTE_INJECTION_KIND_STRING       = "attribute-injection"
	ATTRIBUTE_VALUE_INJECTION_KIND_STRING = "attribute-value-injection"
	TEMPLATE_KIND_STRING                  = "template"
	TEMPLATE_INJECTION_KIND_STRING        = "template-injection"
	COMMENT_KIND_STRING                   = "comment"
	DOCTYPE_KIND_STRING                   = "doctype"
	TEXT_INJECTION_KIND_STRING            = "text-injection"
	UNSAFE_TEXT_INJECTION_KIND_STRING     = "unsafe-text-injection"
	TEXT_KIND_STRING                      = "text"
	UNSAFE_TEXT_KIND_STRING               = "unsafe-text"
	REPEAT_KIND_STRING                    = "repeat"
	VARIANTS_KIND_STRING                  = "variants"
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
		return ELEMENT_KIND_STRING
	case ATTRIBUTE_NODE_KIND:
		return ATTRIBUTE_KIND_STRING
	case ATTRIBUTES_INJECTION_NODE_KIND:
		return ATTRIBUTE_INJECTION_KIND_STRING
	case ATTRIBUTE_VALUE_INJECTION_NODE_KIND:
		return ATTRIBUTE_VALUE_INJECTION_KIND_STRING
	case TEMPLATE_NODE_KIND:
		return TEMPLATE_KIND_STRING
	case TEMPLATE_INJECTION_NODE_KIND:
		return TEMPLATE_INJECTION_KIND_STRING
	case COMMENT_NODE_KIND:
		return COMMENT_KIND_STRING
	case DOCTYPE_NODE_KIND:
		return DOCTYPE_KIND_STRING
	case TEXT_INJECTION_NODE_KIND:
		return TEXT_INJECTION_KIND_STRING
	case UNSAFE_TEXT_INJECTION_NODE_KIND:
		return UNSAFE_TEXT_INJECTION_KIND_STRING
	case TEXT_NODE_KIND:
		return TEXT_KIND_STRING
	case UNSAFE_TEXT_NODE_KIND:
		return UNSAFE_TEXT_KIND_STRING
	case REPEAT_NODE_KIND:
		return REPEAT_KIND_STRING
	case VARIANTS_NODE_KIND:
		return VARIANTS_KIND_STRING
	default:
		panic(fmt.Sprintf("wrong element kind `%d`", k))
	}
}
