package go2html

import (
	"fmt"
	"strings"
)

type (
	TemplateSpec struct {
		name  string
		nodes []Node
	}
	Template struct {
		name      string
		fragments []interface{}
	}
	NodeReport struct {
		Title    string
		Messages []string
		Children []*NodeReport
	}
	TemplateNode struct {
		alias     string
		scope     string
		fragments []interface{}
	}
)

func Spec(name string, nodes ...Node) *TemplateSpec {
	return &TemplateSpec{
		name:  name,
		nodes: nodes,
	}
}
func (t *TemplateSpec) Precompile() (*Template, *NodeReport) {
	template := &Template{
		name:      t.name,
		fragments: []interface{}{},
	}
	pp := &PrecompilingProxy{
		template: template,
		nodeReport: &NodeReport{
			Title:    fmt.Sprintf("TEMPLATE(%s)", t.name),
			Messages: []string{},
			Children: []*NodeReport{},
		},
	}
	for _, n := range t.nodes {
		n.Commit(pp.Child(n.Title()))
	}
	return template, pp.nodeReport
}
func (t *Template) Populate(replacements map[string]interface{}) string {
	var result strings.Builder
	for _, rawFragment := range t.fragments {
		switch fragment := rawFragment.(type) {
		case string:
			result.WriteString(fragment)
		case PopulatingNode:
			reps := replacements
			key, ok := fragment.Scope()
			if ok {
				reps, ok := replacements[key]
				if !ok {
					panic(fmt.Sprintf("replacement for \"%s\" key is not provied", key))
				}
				result.WriteString(fragment.Populate(reps))
				continue
			}
			result.WriteString(fragment.Populate(reps))
		case injectionKey:
			key := string(fragment)
			rawReplacement, ok := replacements[key]
			if !ok {
				panic(fmt.Sprintf("replacement for \"%s\" key is not provied", key))
			}
			str, ok := rawReplacement.(string)
			if !ok {
				panic("wrong replacement type")
			}
			result.WriteString(str)
		default:
			panic(fmt.Sprintf("wrong fragment type %#v", rawFragment))

		}
	}
	return result.String()
}

const NO_INJECTION_SCOPE = ""
const NO_ALIAS = ""
const INJECTION_SCOPE_TEMPLATE = "%s.%s"
const TEMPLATE_NODE_TITLE_TEMPLATE = "TEMPLATE(%s)"

func (t *Template) Node(alias, scope string) *TemplateNode {
	if alias == NO_ALIAS {
		alias = t.name
	}
	return &TemplateNode{
		alias:     alias,
		scope:     scope,
		fragments: t.fragments,
	}
}
func (n *TemplateNode) Title() string {
	return fmt.Sprintf(TEMPLATE_NODE_TITLE_TEMPLATE, n.alias)
}
func (n *TemplateNode) Commit(pp *PrecompilingProxy) {
	pp.AppendFragment(n)
	pp.Report("ok")
}
func (n *TemplateNode) Populate(rawReplacements interface{}) string {
	replacements, ok := rawReplacements.(map[string]interface{})
	if !ok {
		panic("replacements should be of type map[string]interface{}")
	}
	var (
		result  strings.Builder
		rawReps interface{}
	)
	for _, rawFragment := range n.fragments {
		switch fragment := rawFragment.(type) {
		case string:
			result.WriteString(fragment)
		case PopulatingNode:
			reps := replacements
			key, ok := fragment.Scope()
			if ok {
				rawReps, ok = replacements[key]
				if !ok {
					panic(fmt.Sprintf("replacement for \"%s\" key is not provied", key))
				}
				reps, ok = rawReps.(map[string]interface{})
				if !ok {
					panic("wrong replacement type")
				}
			}
			result.WriteString(fragment.Populate(reps))
		case injectionKey:
			key := string(fragment)
			rawReplacement, ok := replacements[key]
			if !ok {
				panic(fmt.Sprintf("replacement for \"%s\" key is not provied", key))
			}
			str, ok := rawReplacement.(string)
			if !ok {
				panic("wrong replacement type")
			}
			result.WriteString(str)
		default:
			panic(fmt.Sprintf("wrong fragment type %#v", rawFragment))
		}
	}
	return result.String()
}
func (n *TemplateNode) Scope() (string, bool) {
	return n.scope, n.scope != NO_INJECTION_SCOPE
}
