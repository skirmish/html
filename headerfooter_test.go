package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var headerfooterTestCases = []HTMLTestCase{
	{
		name:    "header",
		output:  "<header/>",
		element: Header(),
	},
	{
		name:   "header with children",
		output: "<header class=\"name\">some content</header>",
		element: Header(Class("name")).AddElements(
			Content("some content"),
		),
	},
	{
		name:    "footer",
		output:  "<footer/>",
		element: Footer(),
	},
	{
		name:   "footer with children",
		output: "<footer class=\"name\">some content</footer>",
		element: Footer(Class("name")).AddElements(
			Content("some content"),
		),
	},
	{
		name:    "section",
		output:  "<section/>",
		element: Section(),
	},
	{
		name:    "section with content",
		output:  "<section class=\"name\">content</section>",
		element: Section(Class("name")).AddElements(Content("content")),
	},
	{
		name:    "nav",
		output:  "<nav/>",
		element: Nav(),
	},
	{
		name:    "nav with content",
		output:  "<nav class=\"name\">content</nav>",
		element: Nav(Class("name")).AddElements(Content("content")),
	},
	{
		name:    "aside",
		output:  "<aside/>",
		element: Aside(),
	},
	{
		name:    "aside with content and class",
		output:  "<aside class=\"name\">something</aside>",
		element: Aside(Class("name")).AddElements(Content("something")),
	},
	{
		name:    "article",
		output:  "<article/>",
		element: Article(),
	},
	{
		name:    "article with content and class",
		output:  "<article class=\"name\">something</article>",
		element: Article(Class("name")).AddElements(Content("something")),
	},
	{
		name:   "details with content, summary and class",
		output: "<details class=\"name\"><summary class=\"name\">something</summary><p>something</p></details>",
		element: Details(Class("name")).AddElements(
			Summary(Class("name")).AddElements(Content("something")),
			P().AddElements(Content("something")),
		),
	},
	{
		name:    "main with content and class",
		output:  "<main class=\"name\">something</main>",
		element: Main(Class("name")).AddElements(Content("something")),
	},
	{
		name:   "mark with content and class",
		output: "<p class=\"name\"><mark class=\"id\">something</mark></p>",
		element: P(Class("name")).AddElements(
			Mark(Class("id")).AddElements(Content("something")),
		),
	},
	{
		name:   "time with content and class",
		output: "<p class=\"name\">some part should be done by<time id=\"id1\">10:00pm</time>, probably.</p>",
		element: P(Class("name")).AddElements(
			Content("some part should be done by"),
			Time(Id("id1")).AddElements(Content("10:00pm")),
			Content(", probably."),
		),
	},
}

func Test_HeaderFooterGeneration(t *testing.T) {
	for _, testCase := range headerfooterTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
