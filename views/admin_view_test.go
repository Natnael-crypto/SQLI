package views

import (
	"bytes"
	"html/template"
	"sqli/initializers"
	"sqli/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminRender_Success(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template using html/template
	initializers.Template = template.Must(template.New("admin.html").Parse("Username: {{.Username}}, Products: {{len .Products}}"))

	adminVM := AdminVM{
		Username: "admin_user",
		Products: []models.Product{
			{ID: 1, Name: "Laptop", Category: "Electronics", Price: 1200.00, Description: "Gaming laptop"},
			{ID: 2, Name: "Mouse", Category: "Electronics", Price: 25.00, Description: "Wireless mouse"},
		},
	}

	AdminRender(&buf, adminVM)

	expectedOutput := "Username: admin_user, Products: 2"
	assert.Equal(t, expectedOutput, buf.String())
}

func TestAdminRender_TemplateExecutionFailure(t *testing.T) {
	var buf bytes.Buffer

	// Mock an invalid template (missing required field to trigger an error)
	initializers.Template = template.Must(template.New("admin.html").Parse("{{.InvalidField}}"))

	adminVM := AdminVM{
		Username: "admin_user",
	}

	AdminRender(&buf, adminVM)

	// Expect an empty output because template execution should fail
	assert.Empty(t, buf.String())
}

func TestAdminRender_EmptyProducts(t *testing.T) {
	var buf bytes.Buffer

	// Mock the template for an empty product list scenario
	initializers.Template = template.Must(template.New("admin.html").Parse("Username: {{.Username}}, Products: {{len .Products}}"))

	adminVM := AdminVM{
		Username: "admin_user",
		Products: []models.Product{}, // No products
	}

	AdminRender(&buf, adminVM)

	expectedOutput := "Username: admin_user, Products: 0"
	assert.Equal(t, expectedOutput, buf.String())
}
