package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var brTestCases = []HTMLTestCase{
	{
		name:    "br",
		output:  "<br/>",
		element: Br(),
	},
	{
		name:    "b with content and class",
		output:  "<br class=\"name\"/>",
		element: Br(Class("name")).AddElements(Content("something")),
	},
}

func Test_BrGeneration(t *testing.T) {
	for _, testCase := range brTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
