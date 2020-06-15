package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var definitionTestCases = []HTMLTestCase{
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
}

func Test_DefinitionListGeneration(t *testing.T) {
	for _, testCase := range definitionTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
