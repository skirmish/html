# Html

A highly opinionated code driven HTML generator.

Usage:
```golang
    page := []HtmlElement{
		DocType(),
		Html(Lang("en")).AddElements(
			Head().AddElements(
				Meta(Key("charset", "utf-8")),
				Title().AddElements(Content("Page Title")),
				Style().AddElements(Content(
					"ul.samples {list-style: none; margin: 0; padding: 0;" +
					" margin-block-start: 1px; margin-block-end: 1px;}\n" +
					"ul.samples li {display: inline-block; width: 300px;}\n" +
					"ul.samples li.good {background-color: #008000;}\n")),
			),
			Body().AddElements(
				Heading(Level(1)).AddElements(Content("The Title of the Page")),
				P().AddElements(Content("This is the first sentence."))),
			Ul().AddElements(
				Li().AddElements(
					Img(Src("data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAA7UlEQVR4"+
						"nGJiwAMGXJJZUFCwrmO95LL/3&#43;tBfBaYhByblY1AMIj1ZHLg54sHQSxGqJzBPn4o61/SV4Zn728imyh4&#43;"+
						"y8IHNv2/SMWywPmZP/9e5abQXsWNqfxMcz6G4XLtZ8YPjKkMOL0Fve&#43;v244JRmUPz5ckINTNvDD37/lErhkdX"+
						"f9/TtNGpesQOyfv7txW/zz708HKJMFVUYvxJSF4dpBbJrUpzz9&#43;/fvr21YpCSK7oKC96QfppS401VwyAdiBpL"+
						"ganCsHA7gwJAyX/MIJPWllRtVHOzawEAGhuub//Z8wO1BTEB&#43;0gQEAAD//3QpXYUSeN4MAAAAAElFTkSuQmCC")),
				),
			),
		),
	}

	output := ""
	for _, el := range page {
		output += el.Render()
	}
```

Elements currently supported (absolute bare minimum):
 * doctype
 * html
 * head
 * title
 * style
 * body
 * h1-6
 * p
 * img
 * a
 * ul
 * li
 

Notes:
This is not optimized yet, first we'll work on correctness, then once it's fairly stable then working on making it 
faster will happen.

All valid html5 tags *want* to be implemented, but for now it's driven by the ones currently in use. 
