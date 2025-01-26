package views

import (
	"bytes"
	"html/template"
	"sqli/initializers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginRender_Success(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for login rendering
	initializers.Template = template.Must(template.New("login.html").Parse("Error: {{.}}"))

	// Test case with no error message
	LoginRender(&buf)

	expectedOutput := "Error: "
	assert.Equal(t, expectedOutput, buf.String())
}

func TestLoginRender_WithError(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for login rendering
	initializers.Template = template.Must(template.New("login.html").Parse("Error: {{.}}"))

	// Test case with an error message
	LoginRender(&buf, assert.AnError)

	expectedOutput := "Error: " + assert.AnError.Error()
	assert.Equal(t, expectedOutput, buf.String())
}

func TestChangePasswordRender_Success(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for change password rendering
	initializers.Template = template.Must(template.New("change_password.html").Parse("Has Error: {{.}}"))

	// Test case with no error
	ChangePasswordRender(&buf, false)

	expectedOutput := "Has Error: false"
	assert.Equal(t, expectedOutput, buf.String())
}

func TestChangePasswordRender_WithError(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for change password rendering
	initializers.Template = template.Must(template.New("change_password.html").Parse("Has Error: {{.}}"))

	// Test case with error
	ChangePasswordRender(&buf, true)

	expectedOutput := "Has Error: true"
	assert.Equal(t, expectedOutput, buf.String())
}

func TestForgotPasswordRender_Success(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for forgot password rendering
	initializers.Template = template.Must(template.New("forgot_password.html").Parse("Data: {{.}}"))

	// Test case with no additional data
	ForgotPasswordRender(&buf)

	expectedOutput := "Data: "
	assert.Equal(t, expectedOutput, buf.String())
}

func TestForgotPasswordRender_WithData(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for forgot password rendering
	initializers.Template = template.Must(template.New("forgot_password.html").Parse("Data: {{.}}"))

	// Test case with some additional data
	ForgotPasswordRender(&buf, "Some extra data")

	expectedOutput := "Data: Some extra data"
	assert.Equal(t, expectedOutput, buf.String())
}
