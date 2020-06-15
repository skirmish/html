package html

import "io"

type meta struct {
	Element
	KeyVals []KeyVal
}

func Meta(attrs ...func(*meta)) HtmlElement {
	m := &meta{}
	for _, attr := range attrs {
		attr(m)
	}
	return m
}

func Key(key string, value string) func(*meta) {
	return func(m *meta) {
		m.KeyVals = append(m.KeyVals, KeyVal{
			Key:   key,
			Value: value,
		})
	}
}

func (a *meta) Render(w io.Writer) (int, error) {
	return a.Element.Render(w, "meta", nil)
}

func (m *meta) AddElements(elements ...HtmlElement) HtmlElement {
	//for _,element := range elements {
	//	p.Children = append(p.Children,element)
	//}
	return m
}

func (m *meta) addAttribute(key string, val string) {
	m.Element.AddAttribute(key, val)
}

//<meta charset=\"utf-8\"/>"
