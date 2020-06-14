package html

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sectionTestCases = []HTMLTestCase{
	{
		name:    "section",
		output:  "<section></section>",
		element: Section(),
	},
}

func Test_SectionGeneration(t *testing.T) {
	for _, testCase := range sectionTestCases {
		t.Logf("Test %s", testCase.name)

		output := testCase.element.Render()
		assert.Equal(t, testCase.output, output)
	}
}
