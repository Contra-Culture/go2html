package go2html

import "strings"

type (
	Template struct {
		key       string
		fragments []interface{}
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

var (
	safeTextReplacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "\"", "&quot", "'", "&quot")
	voidElements     = []string{
		"area",
		"base",
		"br",
		"col",
		"embed",
		"hr",
		"img",
		"input",
		"link",
		"meta",
		"param",
		"source",
		"track",
		"wbr",
	}
	templateElements = []string{
		"template",
	}
	rawTextElements = []string{
		"script",
		"style",
	}
	escapableRawTextElements = []string{
		"textarea",
		"title",
	}
	normalElements = []string{
		"html",
		"base",
		"head",
		"link",
		"meta",
		"style",
		"title",
		"body",
		"address",
		"article",
		"aside",
		"footer",
		"header",
		"h1",
		"h2",
		"h3",
		"h4",
		"h5",
		"h6",
		"main",
		"nav",
		"section",
		"blockquote",
		"dd",
		"div",
		"dl",
		"dt",
		"figcaption",
		"figure",
		"hr",
		"li",
		"ol",
		"p",
		"pre",
		"ul",
		"a",
		"abbr",
		"b",
		"bdi",
		"bdo",
		"br",
		"cite",
		"code",
		"data",
		"dfn",
		"em",
		"i",
		"kbd",
		"mark",
		"q",
		"rp",
		"rt",
		"ruby",
		"s",
		"samp",
		"small",
		"span",
		"strong",
		"sub",
		"sup",
		"time",
		"u",
		"var",
		"wbr",
		"area",
		"audio",
		"img",
		"map",
		"track",
		"video",
		"embed",
		"iframe",
		"object",
		"param",
		"picture",
		"portal",
		"source",
		"svg",
		"math",
		"canvas",
		"noscript",
		"script",
		"del",
		"ins",
		"caption",
		"col",
		"colgroup",
		"table",
		"tbody",
		"td",
		"tfoot",
		"th",
		"thead",
		"tr",
		"button",
		"datalist",
		"fieldset",
		"form",
		"input",
		"label",
		"legand",
		"meter",
		"optgroup",
		"option",
		"output",
		"progress",
		"select",
		"textarea",
		"details",
		"dialog",
		"menu",
		"summary",
		"slot",
		"template",
		// obsolete/deprecated
		"acronym",
		"applet",
		"baseform",
		"bgsound",
		"big",
		"blink",
		"center",
		"content",
		"dir",
		"font",
		"frame",
		"frameset",
		"hgroup",
		"image",
		"keygen",
		"marquee",
		"menuitem",
		"nobr",
		"noembed",
		"noframes",
		"plaintext",
		"rb",
		"rtc",
		"shadow",
		"spacer",
		"strike",
		"tt",
		"xmp",
	}
)

func Tmplt(key string, config func(*TemplateConfiguringProxy)) *Template {
	t := &Template{
		key:       key,
		fragments: []interface{}{},
	}
	config(&TemplateConfiguringProxy{
		template: t,
	})
	return t
}
func (t *Template) Populate(rawReplacements map[string]interface{}) string {
	var sb strings.Builder
	for _, rawFragment := range t.fragments {
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
