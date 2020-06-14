package html

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

func (m *meta) Render() string {
	output := "<meta"
	m.InitAttributes()
	for _, kv := range m.KeyVals {
		m.Element.AddAttribute(kv.Key, kv.Value)
	}
	//if m.Charset!="" {
	//	m.Element.AddAttribute("charset",m.Charset)
	//}
	output += m.Element.RenderAttr()
	output += "/>"
	return output
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
