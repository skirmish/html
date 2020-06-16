package html

import "io"

type content struct {
	Element
	Content string
}

// Content adds raw content to the stream.
func Content(data string) HtmlElement {
	c := &content{Content: data}
	return c
}

func (c *content) Render(w io.Writer) (int, error) {
	return w.Write([]byte(c.Content))
}

func (c *content) AddElements(elements ...HtmlElement) HtmlElement {
	return c
}

func (c *content) addAttribute(key string, val string) {
	//c.Element.AddAttribute(key, val)
}
