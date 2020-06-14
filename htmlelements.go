package html

type HtmlElement interface {
	Render() string
	AddElements(elements ...HtmlElement) HtmlElement
	addAttribute(key string, val string)
	setId(id string)
}

type KeyVal struct {
	Key   string
	Value string
}

type Element struct {
	Id         string
	Name       string
	attributes []KeyVal
}

func (e *Element) InitAttributes() {
	//e.attributes = make([]KeyVal,0)
}

func (e *Element) AddAttribute(key string, val string) {
	e.attributes = append(e.attributes, KeyVal{
		Key:   key,
		Value: val,
	})
}

func (e *Element) setId(id string) {
	e.AddAttribute("ID", id)
}

func (e Element) RenderAttr() string {
	if len(e.attributes) == 0 {
		return ""
	}
	output := " "
	for index, attr := range e.attributes {
		output += attr.Key
		output += "=\""
		output += attr.Value
		output += "\""
		if index+1 < len(e.attributes) {
			output += " "
		}
	}
	return output
}
