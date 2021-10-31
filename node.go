package go2html

type (
	Node struct {
		Kind     string
		Title    string
		Children []*Node
	}
)

const (
	ELEM_NODE_KIND                      = "element"
	ATTRIBUTE_NODE_KIND                 = "attribute"
	ATTRIBUTE_INJECTION_NODE_KIND       = "attribute-injection"
	ATTRIBUTE_VALUE_INJECTION_NODE_KIND = "attribute-value-injection"
	TEMPLATE_NODE_KIND                  = "template"
	COMMENT_NODE_KIND                   = "comment"
	DOCTYPE_NODE_KIND                   = "doctype"
	TEXT_INJECTION_NODE_KIND            = "text-injection"
	UNSAFE_TEXT_INJECTION_NODE_KIND     = "unsafe-text-injection"
	TEXT_NODE_KIND                      = "text"
	UNSAFE_TEXT_NODE_KIND               = "unsafe-text"
	REPEAT_NODE_KIND                    = "repeat"
)
const (
	COMMENT_NODE_TITLE = "COMMENT"
	DOCTYPE_NODE_TITLE = "!DOCTYPE"
	TEXT_NODE_TITLE    = "TEXT"
)

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}
