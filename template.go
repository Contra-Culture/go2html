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
)

func Tmplt(name string, nodes ...Node) *Template {
	return &Template{
		name:         name,
		nodes:        nodes,
		marks:        []int{},
		marksMapping: map[string]int{},
		precompiled:  []string{},
		report: &NodeReport{
			Title:    fmt.Sprintf("TEMPLATE(%s) ROOT", name),
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
