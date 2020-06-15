package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sectionTestCases = []HTMLTestCase{
	{
		name:    "section",
		output:  "<section/>",
		element: Section(),
	},
}

func Test_SectionGeneration(t *testing.T) {
	for _, testCase := range sectionTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}
