package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var divspanTestCases = []HTMLTestCase{
	{
		name:    "div",
		output:  "<div/>",
		element: Div(),
	},
	{
		name:   "div with children",
		output: "<div class=\"name\">some content</div>",
		element: Div(Class("name")).AddElements(
			Content("some content"),
		),
	},
	{
		name:    "span",
		output:  "<span/>",
		element: Span(),
	},
	{
		name:   "span with children",
		output: "<span class=\"name\">some content</span>",
		element: Span(Class("name")).AddElements(
			Content("some content"),
		),
	},
}

func Test_DivSpanGeneration(t *testing.T) {
	for _, testCase := range divspanTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
