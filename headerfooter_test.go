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
