package go2html

import "fmt"

type (
	ElemConfiguringProxy struct {
		tcp *TemplateConfiguringProxy
	}
)

func (ecp *ElemConfiguringProxy) AttrInjection(key string) {
	ecp.tcp.appendFragment(" ")
	ecp.tcp.appendFragment(injection{
		key: key,
	})
}
func (ecp *ElemConfiguringProxy) AttrValueInjection(name string, key string) {
	ecp.tcp.appendFragment(fmt.Sprintf(" %s=\"", name))
	ecp.tcp.appendFragment(injection{
		key: key,
	})
	ecp.tcp.appendFragment("\"")
}
func (ecp *ElemConfiguringProxy) Attr(name string, value string) {
	ecp.tcp.appendFragment(fmt.Sprintf(" %s=\"%s\"", name, value))
}
