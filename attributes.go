package html

type attribute struct {
	Name string
	Type string
}

var attributes []attribute = []attribute{
	{Name: "accesskey", Type: "global"},
	{Name: "class", Type: "global"},
	{Name: "contenteditable", Type: "global"},
	{Name: "data-*", Type: "global"},
	{Name: "dir", Type: "global"},
	{Name: "draggable", Type: "global"},
	{Name: "hidden", Type: "global"},
	{Name: "id", Type: "global"},
	{Name: "lang", Type: "global"},
	{Name: "spellcheck", Type: "global"},
	{Name: "style", Type: "global"},
	{Name: "tabindex", Type: "global"},
	{Name: "title", Type: "global"},
	{Name: "translate", Type: "global"},
}

// Name - Attribute for form
func Name(name string) func(Element) {
	return func(f Element) {
		f.AddAttribute("name", name)
	}
}

// Src - for image
func Src(source string) func(Element) {
	return func(i Element) {
		i.AddAttribute("src", source)
	}
}

// Href - for image
func Href(href string) func(Element) {
	return func(i Element) {
		i.AddAttribute("href", href)
	}
}

// UseMap - for image
func UseMap(usemap string) func(Element) {
	return func(i Element) {
		i.AddAttribute("usemap", usemap)
	}
}

// Shape - for area
func Shape(shape string) func(Element) {
	return func(i Element) {
		i.AddAttribute("shape", shape)
	}
}

// Coords - for area
func Coords(coords string) func(Element) {
	return func(i Element) {
		i.AddAttribute("coords", coords)
	}
}

// Charset - for meta
func Charset(charset string) func(Element) {
	return func(m Element) {
		m.AddAttribute("charset", charset)
	}
}

// SrcSet - for Source
func SrcSet(srcset string) func(Element) {
	return func(m Element) {
		m.AddAttribute("srcset", srcset)
	}
}

// Media - for source
func Media(media string) func(Element) {
	return func(m Element) {
		m.AddAttribute("media", media)
	}
}

// Type - script/style
func Type(t string) func(Element) {
	return func(m Element) {
		m.AddAttribute("type", t)
	}
}

// ContentAttr - meta
func ContentAttr(t string) func(Element) {
	return func(m Element) {
		m.AddAttribute("content", t)
	}
}

//Global Attributes
////accesskey	Specifies a shortcut key to activate/focus an element
////class	Specifies one or more classnames for an element (refers to a class in a style sheet)
////contenteditable	Specifies whether the content of an element is editable or not
////data-*	Used to store custom data private to the page or application
////dir	Specifies the text direction for the content in an element
////draggable	Specifies whether an element is draggable or not
////hidden	Specifies that an element is not yet, or is no longer, relevant
////id	Specifies a unique id for an element
////lang	Specifies the language of the element's content
////spellcheck	Specifies whether the element is to have its spelling and grammar checked or not
////style	Specifies an inline CSS style for an element
////tabindex	Specifies the tabbing order of an element
////title	Specifies extra information about an element
////translate	Specifies whether the content of an element should be translated or not

//Form attributes
//Attribute	Value	Description
//accept-charset	character_set	Specifies the character encodings that are to be used for the form submission
//action	URL	Specifies where to send the form-data when a form is submitted
//autocomplete	on
//off	Specifies whether a form should have autocomplete on or off
//enctype	application/x-www-form-urlencoded
//multipart/form-data
//text/plain	Specifies how the form-data should be encoded when submitting it to the server (only for method="post")
//method	get
//post	Specifies the HTTP method to use when sending form-data
//name	text	Specifies the name of a form
//novalidate	novalidate	Specifies that the form should not be validated when submitted
//rel	external
//help
//license
//next
//nofollow
//noopener
//noreferrer
//opener
//prev
//search	Specifies the relationship between a linked resource and the current document
//target	_blank
//_self
//_parent
//_top	Specifies where to display the response that is received after submitting the form

/*
// Form event attributes
Attribute	Value	Description
onblur	script	Fires the moment that the element loses focus
onchange	script	Fires the moment when the value of the element is changed
oncontextmenu	script	Script to be run when a context menu is triggered
onfocus	script	Fires the moment when the element gets focus
oninput	script	Script to be run when an element gets user input
oninvalid	script	Script to be run when an element is invalid
onreset	script	Fires when the Reset button in a form is clicked
onsearch	script	Fires when the user writes something in a search field (for <input="search">)
onselect	script	Fires after some text has been selected in an element
onsubmit	script	Fires when a form is submitted*/

// FormId - meter
func FormId(formid string) func(Element) {
	return func(i Element) {
		i.AddAttribute("form", formid)
	}
}

// High - meter
func High(high string) func(Element) {
	return func(i Element) {
		i.AddAttribute("high", high)
	}
}

// Low - meter
func Low(low string) func(Element) {
	return func(i Element) {
		i.AddAttribute("low", low)
	}
}

// Max - meter
func Max(max string) func(Element) {
	return func(i Element) {
		i.AddAttribute("max", max)
	}
}

// Min - meter
func Min(min string) func(Element) {
	return func(i Element) {
		i.AddAttribute("min", min)
	}
}

// Optimum - meter
func Optimum(optimum string) func(Element) {
	return func(i Element) {
		i.AddAttribute("optimum", optimum)
	}
}

// Value - meter
func Value(value string) func(Element) {
	return func(i Element) {
		i.AddAttribute("value", value)
	}
}

// Dir - bdo
func Dir(dir string) func(Element) {
	return func(i Element) {
		i.AddAttribute("dir", dir)
	}
}

/* Meter Attributes
Attribute	Value	Description
form	form_id	Specifies which form the <meter> element belongs to
high	number	Specifies the range that is considered to be a high value
low	number	Specifies the range that is considered to be a low value
max	number	Specifies the maximum value of the range
min	number	Specifies the minimum value of the range. Default value is 0
optimum	number	Specifies what value is the optimal value for the gauge
value	number	Required. Specifies the current value of the gauge
*/
