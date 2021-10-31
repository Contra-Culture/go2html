package go2html

import (
	"strings"

	"github.com/Contra-Culture/go2html/fragments"
)

type (
	Template struct {
		key       string
		nodes     []*Node
		fragments *fragments.Fragments
	}
	elemType  int
	injection struct {
		key       string
		modifiers []func(string) string
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

func NewTemplate(key string, configure func(*TemplateConfiguringProxy)) *Template {
	fs := []interface{}{}
	t := &Template{
		key:       key,
		nodes:     []*Node{},
		fragments: fragments.New(fs),
	}
	configure(&TemplateConfiguringProxy{
		template: t,
	})
	return t
}
func (t *Template) Nodes() []*Node {
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
