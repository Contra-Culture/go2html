package go2html

type PrecompilingProxy struct {
	template   *Template
	nodeReport *NodeReport
}

func (pp *PrecompilingProxy) Report(message string) {
	pp.nodeReport.Messages = append(pp.nodeReport.Messages, message)
}
func (pp *PrecompilingProxy) Child(title string) *PrecompilingProxy {
	nr := &NodeReport{
		Title:    title,
		Messages: []string{},
		Children: []*NodeReport{},
	}
	pp.nodeReport.Children = append(pp.nodeReport.Children, nr)
	return &PrecompilingProxy{
		template:   pp.template,
		nodeReport: nr,
	}
}
func (pp *PrecompilingProxy) AppendFragment(fragment interface{}) {
	t := pp.template
	t.fragments = append(t.fragments, fragment)
}
