package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var asideTestCases = []HTMLTestCase{
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
}

func Test_AsideGeneration(t *testing.T) {
	for _, testCase := range asideTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
