package go2html



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
		Title() string
		WriteTo(btc *BreakthroughContext)
	}
)
