package go2html

import (
	"strings"

	"github.com/Contra-Culture/go2html/fragments"
	"github.com/Contra-Culture/go2html/node"
)

type (
	Constructor struct {
		nodes []*node.Node
	}
	Template struct {
		key       string
		nodes     []*node.Node
		fragments *fragments.Fragments
	}
	elemType  int
	injection struct {
		key       string
		modifiers []func(string) string
	}
	templateInjection struct {
		key string
	}
	repetition struct {
		key      string
		template *Template
	}
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

var safeTextReplacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "\"", "&quot", "'", "&quot")

func NewTemplate(key string, configure func(*TemplateCfgr)) *Template {
	fs := []interface{}{}
	t := &Template{
		key:       key,
		nodes:     []*node.Node{},
		fragments: fragments.New(fs),
	}
	configure(&TemplateCfgr{
		template: t,
	})
	return t
}
func (t *Template) Nodes() []*node.Node {
	return t.nodes
}
func (t *Template) Fragments() []interface{} {
	return t.fragments.Fragments()
}
func (t *Template) Populate(rawReplacements map[string]interface{}) string {
	var sb strings.Builder
	for _, rawFragment := range t.fragments.Fragments() {
		switch fragment := rawFragment.(type) {
		case string:
			sb.WriteString(fragment)
		case injection:
			rawRepl, _ := rawReplacements[fragment.key]
			repl, _ := rawRepl.(string)
			for _, modify := range fragment.modifiers {
				repl = modify(repl)
			}
			sb.WriteString(repl)
		case templateInjection:
			rawNestedReplacement, _ := rawReplacements[fragment.key]
			nestedReplacement, _ := rawNestedReplacement.(map[string]interface{})
			rawTemplate, _ := nestedReplacement["template"]
			rawValues := nestedReplacement["values"]
			switch values := rawValues.(type) {
			case map[string]interface{}:
				switch template := rawTemplate.(type) {
				case *Template:
					sb.WriteString(template.Populate(values))
				default:
					panic("no template provided")
				}
			default:
				panic("no values provided")
			}
		case *Template:
			rawNestedReplacement, _ := rawReplacements[fragment.key]
			nestedReplacement, _ := rawNestedReplacement.(map[string]interface{})
			sb.WriteString(fragment.Populate(nestedReplacement))
		case repetition:
			rawNestedReplacements, _ := rawReplacements[fragment.key]
			nestedReplacements, _ := rawNestedReplacements.([]map[string]interface{})
			for _, nestedReplacement := range nestedReplacements {
				result := fragment.template.Populate(nestedReplacement)
				sb.WriteString(result)
			}
		}
	}
	return sb.String()
}
func HTMLEscape(raw string) string {
	return safeTextReplacer.Replace(raw)
}
func elemTyp(name string) elemType {
	for _, ve := range voidElements {
		if name == ve {
			return VOID_ELEM_TYPE
		}
	}
	for _, te := range templateElements {
		if name == te {
			return TEMPLATE_ELEM_TYPE
		}
	}
	for _, rte := range rawTextElements {
		if name == rte {
			return RAW_TEXT_ELEM_TYPE
		}
	}
	for _, erte := range escapableRawTextElements {
		if name == erte {
			return ESCAPABLE_RAW_TEXT_ELEM_TYPE
		}
	}
	for _, ne := range normalElements {
		if name == ne {
			return NORMAL_ELEM_TYPE
		}
	}
	return NO_ELEM_TYPE
}
