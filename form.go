package html

type form struct {
	Element
	Children []HtmlElement
}

func Form(attrs ...func(HtmlElement)) HtmlElement {
	f := &form{}
	for _, attr := range attrs {
		attr(f)
	}
	return f
}

func (f *form) Render() string {
	f.Element.InitAttributes()
	output := "<form"
	if f.Id != "" {
		f.AddAttribute("ID", f.Id)
	}
	if f.Name != "" {
		f.AddAttribute("name", f.Name)
	}
	output += f.Element.RenderAttr() // Attributes
	output += ">"

	for _, child := range f.Children {
		output += child.Render()
	}

	// Render end
	output += "</form>"
	return output
}

func (f *form) AddElements(elements ...HtmlElement) HtmlElement {
	for _, element := range elements {
		f.Children = append(f.Children, element)
	}
	return f
}

func (f *form) addAttribute(key string, val string) {
	f.Element.AddAttribute(key, val)
}
