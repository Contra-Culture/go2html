package go2html

import "fmt"

type BreakthroughContext struct {
	template   *Template
	nodeReport *NodeReport
}

func (btc *BreakthroughContext) Report(message string) {
	btc.nodeReport.Messages = append(btc.nodeReport.Messages, message)
}
func (btc *BreakthroughContext) Child(title string) *BreakthroughContext {
	nr := &NodeReport{
		Title:    title,
		Messages: []string{},
		Children: []*NodeReport{},
	}
	btc.nodeReport.Children = append(btc.nodeReport.Children, nr)
	return &BreakthroughContext{
		template:   btc.template,
		nodeReport: nr,
	}
}
func (btc *BreakthroughContext) WriteFragment(f string) {
	t := btc.template
	isLastElementReplacement := len(t.marks) > 0 && t.marks[len(t.marks)-1] == len(t.precompiled)-1
	if isLastElementReplacement || len(t.precompiled) == 0 {
		t.precompiled = append(t.precompiled, f)
		return
	}
	lastIdx := len(t.precompiled) - 1
	t.precompiled[lastIdx] = fmt.Sprintf("%s%s", t.precompiled[lastIdx], f)
}
func (btc *BreakthroughContext) MarkInjection(key string) {
	t := btc.template
	mark := len(t.precompiled)
	if mark < 0 {
		mark = 0
	}
	t.marks = append(t.marks, mark)
	t.marksMapping[key] = mark
	t.precompiled = append(t.precompiled, fmt.Sprintf("{{ %s }}", key))
}
