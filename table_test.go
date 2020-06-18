package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tableTestCases = []HTMLTestCase{
	{
		name:    "table",
		output:  "<table/>",
		element: Table(),
	},
	{
		name: "table with content",
		output: "<table class=\"atab\">" +
			"<caption class=\"myc\">caption</caption>" +
			"<colgroup><col class=\"col1\"/><col class=\"col2\"/></colgroup>" +
			"<thead class=\"theadc\"><tr class=\"head\"><th class=\"the\">head1</th><th class=\"the2\">head2</th></tr></thead>" +
			"<tbody class=\"tb\"><tr class=\"r1\"><td class=\"td1\">data11</td><td class=\"td2\">data21</td></tr><tr class=\"r2\"><td class=\"td1\">data12</td><td class=\"td2\">data22</td></tr></tbody>" +
			"<tfoot class=\"tf\"><td class=\"td1\">foot1</td><td class=\"td2\">foot2</td></tfoot></table>",

		element: Table(Class("atab")).AddElements(
			Caption(Class("myc")).AddElements(Content("caption")),
			Colgroup().AddElements(
				Col(Class("col1")),
				Col(Class("col2")),
			),
			Thead(Class("theadc")).AddElements(
				Tr(Class("head")).AddElements(
					Th(Class("the")).AddElements(Content("head1")),
					Th(Class("the2")).AddElements(Content("head2")),
				),
			),
			Tbody(Class("tb")).AddElements(
				Tr(Class("r1")).AddElements(
					Td(Class("td1")).AddElements(Content("data11")),
					Td(Class("td2")).AddElements(Content("data21")),
				),
				Tr(Class("r2")).AddElements(
					Td(Class("td1")).AddElements(Content("data12")),
					Td(Class("td2")).AddElements(Content("data22")),
				),
			),
			Tfoot(Class("tf")).AddElements(
				Td(Class("td1")).AddElements(Content("foot1")),
				Td(Class("td2")).AddElements(Content("foot2")),
			),
		),
	},
}

func Test_TableGeneration(t *testing.T) {
	for _, testCase := range tableTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
