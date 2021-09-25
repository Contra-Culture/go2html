package go2html

import (
	"fmt"
	"strings"
)

type (
	RepNode struct {
		name       string
		key        string
		repeatable *TemplateNode
	}
)

const REP_NODE_TITLE_PATTERN = "{%s}*"
const NO_REP_SCOPE = ""

func Rep(name string, key string, repeatable Node) *RepNode {
	rep, ok := repeatable.(*TemplateNode)
	if !ok {
		panic("repeatable should be of type *TemplateNode")
	}
	return &RepNode{
		name:       name,
		key:        key,
		repeatable: rep,
	}
}
func (n *RepNode) Title() string {
	return fmt.Sprintf(REP_NODE_TITLE_PATTERN, n.name)
}
func (n *RepNode) Commit(pp *PrecompilingProxy) {
	pp.AppendFragment(n)
	pp.Report("ok")
}
func (n *RepNode) Populate(rawReplacements interface{}) string {
	var (
		replacements []interface{}
		ok           bool
	)
	if rawReplacements != nil {
		replacements, ok = rawReplacements.([]interface{})
		if !ok {
			panic("replacements should be of type []interface{}")
		}
	}
	var result strings.Builder
	for i, rawRepl := range replacements {
		repl, ok := rawRepl.(map[string]interface{})
		if !ok {
			panic(fmt.Sprintf("wrong replacement type \"%s[%d]\"", n.key, i))
		}
		result.WriteString(n.repeatable.Populate(repl))
	}
	return result.String()
}
func (n *RepNode) Scope() (string, bool) {
	return n.key, true
}
