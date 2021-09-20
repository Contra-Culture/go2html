package go2html

import (
	"fmt"
	"strings"
)

const (
	NO_ELEM_TYPE = elemType(iota)
	VOID_ELEM_TYPE
	TEMPLATE_ELEM_TYPE
	RAW_TEXT_ELEM_TYPE
	ESCAPABLE_RAW_TEXT_ELEM_TYPE
	FOREIGN_ELEM_TYPE
	NORMAL_ELEM_TYPE
)

const (
	injectionType = iota
	elementType
	textType
	commentType
	attrType
	attrInjectionType
	attrValueInjectionType
)

type (
	Node interface {
		title() string
		writeTo(btc *breakthroughContext)
	}
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
	btc := &breakthroughContext{
		template:   t,
		nodeReport: t.report,
	}
	for _, node := range t.nodes {
		node.writeTo(btc.child(node.title()))
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
