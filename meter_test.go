package html

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var meterTestCases = []HTMLTestCase{
	{
		name:    "meter",
		output:  "<meter></meter>",
		element: Meter(),
	},
	{
		name:   "meter with content",
		output: "<meter id=\"disk_c\" value=\"2\" min=\"0\" max=\"10\">2 out of 10</meter>",
		element: Meter(
			Id("disk_c"),
			Value("2"),
			Min("0"),
			Max("10"),
		).AddElements(
			Content("2 out of 10"),
		),
	},
}

func Test_MeterGeneration(t *testing.T) {
	for _, testCase := range meterTestCases {
		t.Logf("Test %s", testCase.name)

		output := testCase.element.Render()
		assert.Equal(t, testCase.output, output)
	}
}
