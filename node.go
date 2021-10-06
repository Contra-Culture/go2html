package go2html

type (
	Node struct {
		PosBegin                 FragmentPosition
		PosEnd                   FragmentPosition
		Kind                     string
		Title                    string
		Attributes               map[string]string
		AttributeInjections      map[string]injection
		AttributeValueInjections map[string]injection
		Children                 []*Node
	}
)

const (
	ELEM_NODE_KIND                  = "element"
	TEMPLATE_NODE_KIND              = "template"
	COMMENT_NODE_KIND               = "comment"
	DOCTYPE_NODE_KIND               = "doctype"
	TEXT_INJECTION_NODE_KIND        = "text-injection"
	UNSAFE_TEXT_INJECTION_NODE_KIND = "unsafe-text-injection"
	TEXT_NODE_KIND                  = "text"
	UNSAFE_TEXT_NODE_KIND           = "unsafe-text"
	REPEAT_NODE_KIND                = "repeat"
)
const (
	COMMENT_NODE_TITLE = "COMMENT"
	DOCTYPE_NODE_TITLE = "!DOCTYPE"
	TEXT_NODE_TITLE    = "TEXT"
)

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}
func (n *Node) Position() [2][2]int {
	return [2][2]int{
		{
			n.PosBegin.FragmentIndex,
			n.PosBegin.RangeBegin,
		},
		{
			n.PosEnd.FragmentIndex,
			n.PosEnd.RangeEnd,
		},
	}
}
