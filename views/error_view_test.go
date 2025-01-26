package views

import (
	"bytes"
	"html/template"
	"sqli/initializers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorRender_Success(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for error rendering
	initializers.Template = template.Must(template.New("error.html").Parse("<div class='error'>{{.}}</div>"))

	errorMsg := "An unexpected error occurred."

	ErrorRender(&buf, errorMsg)

	expectedOutput := "<div class='error'>An unexpected error occurred.</div>"
	assert.Equal(t, expectedOutput, buf.String())
}

func TestErrorRender_TemplateExecutionFailure(t *testing.T) {
	var buf bytes.Buffer

	// Mock a faulty template that will cause execution to fail
	initializers.Template = template.Must(template.New("error.html").Parse("{{.InvalidField}}"))

	errorMsg := "An unexpected error occurred."

	ErrorRender(&buf, errorMsg)

	// Expect an empty output since the template execution should fail
	assert.Empty(t, buf.String())
}

func TestErrorRender_EmptyMessage(t *testing.T) {
	var buf bytes.Buffer

	// Mock template handling empty error messages
	initializers.Template = template.Must(template.New("error.html").Parse("<div class='error'>{{.}}</div>"))

	errorMsg := ""

	ErrorRender(&buf, errorMsg)

	expectedOutput := "<div class='error'></div>"
	assert.Equal(t, expectedOutput, buf.String())
}