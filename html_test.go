package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type HTMLTestCase struct {
	name    string
	output  string
	element HtmlElement
}

var htmlTestCases = []HTMLTestCase{
	{
		name:    "doctype",
		output:  "<!DOCTYPE html>\n",
		element: DocType(),
	},
	{
		name:    "html",
		output:  "<html/>",
		element: Html(),
	},
	{
		name:    "html lang=en",
		output:  "<html lang=\"en\"/>",
		element: Html(Lang("en")),
	},
	{
		name:    "html lang=en",
		output:  "<html lang=\"en-us\"/>",
		element: Html(Lang("en-us")),
	},
	{
		name:    "body",
		output:  "<body/>",
		element: Body(),
	},
	{
		name:    "body with class",
		output:  "<body class=\"name\"/>",
		element: Body(Class("name")),
	},
	{
		name:    "body with class and content",
		output:  "<body class=\"name\">content</body>",
		element: Body(Class("name")).AddElements(Content("content")),
	},
	{
		name:    "form",
		output:  "<form/>",
		element: Form(),
	},
	{
		name:    "form with id",
		output:  "<form id=\"formId1\"/>",
		element: Form(Id("formId1")),
	},
	{
		name:    "form with id,name",
		output:  "<form id=\"formId1\" name=\"somename\"/>",
		element: Form(Id("formId1"), Name("somename")),
	},
	{
		name:    "fieldset",
		output:  "<fieldset/>",
		element: Fieldset(),
	},
	{
		name:    "fieldset with formid",
		output:  "<fieldset form=\"someform\"/>",
		element: Fieldset(FormId("someform")),
	},
	{
		name:   "form with fieldset with formid",
		output: "<form><fieldset form=\"someform\"><legend class=\"l\">Section</legend><label class=\"som\"><input id=\"in1\" name=\"nom\"/></label><label class=\"som\"><textarea id=\"ta1\">content of the ares</textarea></label><label class=\"sel\"><select class=\"sel\"><optgroup class=\"som\"><option value=\"1\">One</option><option value=\"2\">Two</option></optgroup></select></label><button class=\"cl1\" id=\"sub\" value=\"button\">Click</button><output class=\"cl1\">content</output></fieldset></form>",
		element: Form().AddElements(
			Fieldset(FormId("someform")).AddElements(
				Legend(Class("l")).AddElements(Content("Section")),
				Label(Class("som")).AddElements(
					Input(Id("in1"), Name("nom")),
				),
				Label(Class("som")).AddElements(
					TextArea(Id("ta1")).AddElements(Content("content of the ares")),
				),
				Label(Class("sel")).AddElements(
					Select(Class("sel")).AddElements(
						OptGroup(Class("som")).AddElements(
							Option(Value("1")).AddElements(Content("One")),
							Option(Value("2")).AddElements(Content("Two")),
						),
					),
				),
				Button(Class("cl1"), Id("sub"), Value("button")).AddElements(Content("Click")),
				Output(Class("cl1")).AddElements(Content("content")),
			),
		),
	},
	{
		name:    "Page with head and body",
		output:  "<html><head/><body/></html>",
		element: Html().AddElements(Head(), Body()),
	},
	{
		name:    "script",
		output:  "<script/>",
		element: Script(),
	},
	{
		name:    "script with type and content",
		output:  "<script type=\"text/javascript\">alert(1);\n</script>",
		element: Script(Type("text/javascript")).AddElements(Content("alert(1);\n")),
	},
	{
		name:    "p",
		output:  "<p/>",
		element: P(),
	},
	{
		name:    "p with content",
		output:  "<p>content</p>",
		element: P().AddElements(Content("content")),
	},
	{
		name:    "p with class and content",
		output:  "<p class=\"name\">content</p>",
		element: P(Class("name")).AddElements(Content("content")),
	},
	{
		name:    "style with content",
		output:  "<style>content</style>",
		element: Style().AddElements(Content("content")),
	},
	{
		name:    "style with type content",
		output:  "<style type=\"text/css\">content</style>",
		element: Style(Type("text/css")).AddElements(Content("content")),
	},
	{ //<meta name="viewport" content="width=device-width, initial-scale=1.0">
		name:    "meta viewport",
		output:  "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"/>",
		element: Meta(Name("viewport"), ContentAttr("width=device-width, initial-scale=1.0")),
	},
}

func Test_HtmlGeneration(t *testing.T) {
	for _, testCase := range htmlTestCases {
		t.Logf("Test %s", testCase.name)

		buf := new(bytes.Buffer)
		_, err := testCase.element.Render(buf)
		assert.NoError(t, err, "Rendering")
		assert.Equal(t, testCase.output, buf.String())
	}
}

func runBenchmarkCases(b *testing.B, cases []HTMLTestCase) {
	for _, benchCase := range cases {
		b.Run(benchCase.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf := new(bytes.Buffer)
				n, _ := benchCase.element.Render(buf)
				b.ReportMetric(float64(n), "bytes")
				b.SetBytes(int64(n))
			}
		})
	}
}

/*
func runBenchmarkCasesFast(b *testing.B, cases []HTMLTestCase) {
	for _, benchCase := range cases {
		b.Run(benchCase.name + "-fast", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf := new(bytes.Buffer)
				n, _ := benchCase.element.FastRender(buf)
				b.ReportMetric(float64(n), "bytes")
				b.SetBytes(int64(n))
			}
		})
	}
}
*/
func BenchmarkHtmlGeneration(b *testing.B) {
	runBenchmarkCases(b, htmlTestCases)
	runBenchmarkCases(b, headingTestCases)
	runBenchmarkCases(b, tagsTestCases)
	runBenchmarkCases(b, headerfooterTestCases)
	runBenchmarkCases(b, divspanTestCases)
	runBenchmarkCases(b, mediaTestCases)
	runBenchmarkCases(b, listTestCases)
	runBenchmarkCases(b, tableTestCases)
	runBenchmarkCases(b, kitchenSinkTestCases)

	/*	runBenchmarkCasesFast(b,htmlTestCases)
		runBenchmarkCasesFast(b, headingTestCases)
		runBenchmarkCasesFast(b, tagsTestCases)
		runBenchmarkCasesFast(b, headerfooterTestCases)
		runBenchmarkCasesFast(b, divspanTestCases)
		runBenchmarkCasesFast(b, mediaTestCases)
		runBenchmarkCasesFast(b, listTestCases)
		runBenchmarkCasesFast(b, tableTestCases)
		runBenchmarkCasesFast(b, kitchenSinkTestCases)
	*/
}
