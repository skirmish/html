package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var listTestCases = []HTMLTestCase{
	{
		name:    "ul",
		output:  "<ul/>",
		element: Ul(),
	},
	{
		name:   "ul with children",
		output: "<ul class=\"name\"><li>primary</li><li>secondary</li></ul>",
		element: Ul(Class("name")).AddElements(
			Li().AddElements(Content("primary")),
			Li().AddElements(Content("secondary")),
		),
	},
	{
		name:    "ol",
		output:  "<ol/>",
		element: Ol(),
	},
	{
		name:   "ol with children",
		output: "<ol class=\"name\"><li>first</li><li>second</li></ol>",
		element: Ol(Class("name")).AddElements(
			Li().AddElements(Content("first")),
			Li().AddElements(Content("second")),
		),
	},
	{
		name:    "dl",
		output:  "<dl/>",
		element: Dl(),
	},
	{
		name:   "dl with children",
		output: "<dl class=\"name\"><dt>term</dt><dd>term definition</dd></dl>",
		element: Dl(Class("name")).AddElements(
			Dt().AddElements(Content("term")),
			Dd().AddElements(Content("term definition")),
		),
	},
	{
		name:   "dl with children and classes",
		output: "<dl class=\"name\"><dt class=\"one\">term</dt><dd class=\"two\">term definition</dd></dl>",
		element: Dl(Class("name")).AddElements(
			Dt(Class("one")).AddElements(Content("term")),
			Dd(Class("two")).AddElements(Content("term definition")),
		),
	},
}

func Test_ListGeneration(t *testing.T) {
	for _, testCase := range listTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
