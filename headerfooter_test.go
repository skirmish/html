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
