package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bTestCases = []HTMLTestCase{
	{
		name:    "b",
		output:  "<b/>",
		element: B(),
	},
	{
		name:    "b with content and class",
		output:  "<b class=\"name\">something</b>",
		element: B(Class("name")).AddElements(Content("something")),
	},
}

func Test_BGeneration(t *testing.T) {
	for _, testCase := range bTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
