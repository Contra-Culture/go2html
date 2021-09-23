package go2html

import (
	"fmt"
	"strings"
)

type (
	Template struct {
		name         string
		nodes        []Node
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
	TemplateNode struct {
		name     string
		scope    string
		template *Template
	}
)

func Tmplt(name string, nodes ...Node) *Template {
	return &Template{
		name:         name,
		nodes:        nodes,
		marks:        []int{},
		marksMapping: map[string]int{},
		precompiled:  []string{},
		report: &NodeReport{
			Title:    fmt.Sprintf("TEMPLATE(%s)", name),
			Messages: []string{},
			Children: []*NodeReport{},
		},
	}
}
func (t *Template) isPrecompiled() bool {
	return len(t.precompiled) > 0
}
func (t *Template) Precompile() *NodeReport {
	if t.isPrecompiled() {
		return t.report
	}
	btc := &BreakthroughContext{
		template:   t,
		nodeReport: t.report,
	}
	for _, node := range t.nodes {
		node.WriteTo(btc.Child(node.Title()))
	}
	return t.report
}
func (t *Template) Populate(replacements map[string]interface{}) string {
	if !t.isPrecompiled() {
		panic("template should be precompiled first")
	}
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
			fragments[idx] = repl.Populate(replacements)
		default:
			panic("wrong replacement type")
		}
	}
	return strings.Join(fragments, "")
}

const NO_INJECTION_SCOPE = ""
const NO_ALIAS = ""
const INJECTION_SCOPE_TEMPLATE = "%s.%s"
const TEMPLATE_NODE_TITLE_TEMPLATE = "TEMPLATE(%s)"

func (t *Template) Node(alias, scope string) Node {
	name := t.name
	if alias != NO_ALIAS {
		name = alias
	}
	t.Precompile()
	return &TemplateNode{
		name:     name,
		scope:    scope,
		template: t,
	}
}
func (n *TemplateNode) Title() string {
	return fmt.Sprintf(TEMPLATE_NODE_TITLE_TEMPLATE, n.name)
}
func (n *TemplateNode) WriteTo(btc *BreakthroughContext) {
outer:
	for i, p := range n.template.precompiled {
		for _, mi := range n.template.marks {
			if mi == i {
				key := p[3 : len(p)-3]
				if n.scope != NO_INJECTION_SCOPE {
					key = fmt.Sprintf(INJECTION_SCOPE_TEMPLATE, n.scope, key)
				}
				btc.MarkInjection(key)
				btc.Report(fmt.Sprintf("ok: injection (%s)", key))
				continue outer
			}
		}
		btc.WriteFragment(p)
		btc.Report("ok")
	}
}
